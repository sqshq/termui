// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

// Sparkline is like: ▅▆▂▂▅▇▂▂▃▆▆▆▅▃. The data points should be non-negative integers.
type Sparkline struct {
	Data       []int
	Title      string
	TitleAttrs AttrPair
	LineColor  Attribute
	MaxVal     int
}

// SparklineGroup is a renderable widget which groups together the given sparklines.
type SparklineGroup struct {
	Block
	Sparklines []*Sparkline
}

var SparkChars = []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

// Add appends a given Sparkline to a SparklineGroup
func (sg *SparklineGroup) Add(sl *Sparkline) {
	sg.Sparklines = append(sg.Sparklines, sl)
}

// NewSparkline returns a unrenderable single sparkline that needs to be added to a SparklineGroup
func NewSparkline() *Sparkline {
	return &Sparkline{
		TitleAttrs: Theme.Sparkline.Title,
		LineColor:  Theme.Sparkline.Line,
	}
}

func NewSparklineGroup(sls ...*Sparkline) *SparklineGroup {
	return &SparklineGroup{
		Block:      *NewBlock(),
		Sparklines: sls,
	}
}

func (slg *SparklineGroup) Draw(buf *Buffer) {
	slg.Block.Draw(buf)

	sparklineHeight := (slg.Dy() - 2) / len(slg.Sparklines)

	for i, sl := range slg.Sparklines {
		heightOffset := (sparklineHeight * (i + 1))
		barHeight := sparklineHeight
		if i == len(slg.Sparklines)-1 {
			heightOffset = slg.Dy() - 2
			barHeight = (slg.Dy() - 2) - (sparklineHeight * i)
		}
		if sl.Title != "" {
			barHeight--
		}

		maxVal := sl.MaxVal
		if maxVal == 0 {
			maxVal, _ = GetMaxIntFromSlice(sl.Data)
		}

		// draw line
		for j := 0; j < len(sl.Data) && j < (slg.Dx()-2); j++ {
			data := sl.Data[j]
			height := int((float64(data) / float64(maxVal)) * float64(barHeight))
			sparkChar := SparkChars[len(SparkChars)-1]
			for k := 0; k < height; k++ {
				buf.SetCell(
					Cell{sparkChar, AttrPair{sl.LineColor, ColorDefault}},
					image.Pt(j+slg.Min.X+1, slg.Min.Y+heightOffset-k),
				)
			}
			if height == 0 {
				sparkChar = SparkChars[0]
				buf.SetCell(
					Cell{sparkChar, AttrPair{sl.LineColor, ColorDefault}},
					image.Pt(j+slg.Min.X+1, slg.Min.Y+heightOffset),
				)
			}
		}

		if sl.Title != "" {
			// draw title
			buf.SetString(
				TrimString(sl.Title, slg.Dx()-2),
				image.Pt(slg.Min.X+1, slg.Min.Y+heightOffset-barHeight),
				sl.TitleAttrs,
			)
		}
	}
}

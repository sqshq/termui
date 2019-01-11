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
	LineAttr   Attribute
	MaxVal     int
}

// SparklineGroup is a renderable widget which groups together the given sparklines.
type SparklineGroup struct {
	Block
	Sparklines []*Sparkline
}

// NewSparkline returns a unrenderable single sparkline that needs to be added to a SparklineGroup
func NewSparkline() *Sparkline {
	return &Sparkline{
		TitleAttrs: Theme.Sparkline.Title,
		LineAttr:   Theme.Sparkline.Line,
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

	sparklineHeight := slg.Inner.Dy() / len(slg.Sparklines)

	for i, sl := range slg.Sparklines {
		heightOffset := (sparklineHeight * (i + 1))
		barHeight := sparklineHeight
		if i == len(slg.Sparklines)-1 {
			heightOffset = slg.Inner.Dy()
			barHeight = slg.Inner.Dy() - (sparklineHeight * i)
		}
		if sl.Title != "" {
			barHeight--
		}

		maxVal := sl.MaxVal
		if maxVal == 0 {
			maxVal, _ = GetMaxIntFromSlice(sl.Data)
		}

		// draw line
		for j := 0; j < len(sl.Data) && j < slg.Inner.Dx(); j++ {
			data := sl.Data[j]
			height := int((float64(data) / float64(maxVal)) * float64(barHeight))
			sparkChar := SPARK_CHARS[len(SPARK_CHARS)-1]
			for k := 0; k < height; k++ {
				buf.SetCell(
					Cell{sparkChar, AttrPair{sl.LineAttr, ColorDefault}},
					image.Pt(j+slg.Inner.Min.X, slg.Inner.Min.Y-1+heightOffset-k),
				)
			}
			if height == 0 {
				sparkChar = SPARK_CHARS[0]
				buf.SetCell(
					Cell{sparkChar, AttrPair{sl.LineAttr, ColorDefault}},
					image.Pt(j+slg.Inner.Min.X, slg.Inner.Min.Y-1+heightOffset),
				)
			}
		}

		if sl.Title != "" {
			// draw title
			buf.SetString(
				TrimString(sl.Title, slg.Inner.Dx()),
				sl.TitleAttrs,
				image.Pt(slg.Inner.Min.X, slg.Inner.Min.Y-1+heightOffset-barHeight),
			)
		}
	}
}

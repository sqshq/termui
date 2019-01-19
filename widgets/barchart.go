// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"fmt"
	"image"

	rw "github.com/mattn/go-runewidth"

	. "github.com/gizak/termui"
)

type BarChart struct {
	Block
	BarAttrs   []Attribute
	LabelAttrs []Attribute
	NumAttrs   []Attribute
	NumFmt     func(float64) string
	Data       []float64
	Labels     []string
	BarWidth   int
	BarGap     int
	MaxVal     float64
}

func NewBarChart() *BarChart {
	return &BarChart{
		Block:      *NewBlock(),
		BarAttrs:   Theme.BarChart.Bars,
		NumAttrs:   Theme.BarChart.Nums,
		LabelAttrs: Theme.BarChart.Labels,
		NumFmt:     func(n float64) string { return fmt.Sprint(n) },
		BarGap:     1,
		BarWidth:   3,
	}
}

func (self *BarChart) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	maxVal := self.MaxVal
	if maxVal == 0 {
		maxVal, _ = GetMaxFloat64FromSlice(self.Data)
	}

	barXCoordinate := self.Inner.Min.X

	for i, data := range self.Data {
		// draw bar
		height := int((data / maxVal) * float64(self.Inner.Dy()-1))
		for x := barXCoordinate; x < MinInt(barXCoordinate+self.BarWidth, self.Inner.Max.X); x++ {
			for y := self.Inner.Max.Y - 2; y > (self.Inner.Max.Y-2)-height; y-- {
				c := Cell{' ', AttrPair{ColorDefault, SelectAttr(self.BarAttrs, i)}}
				buf.SetCell(c, image.Pt(x, y))
			}
		}

		// draw label
		if i < len(self.Labels) {
			labelXCoordinate := barXCoordinate +
				int((float64(self.BarWidth) / 2)) -
				int((float64(rw.StringWidth(self.Labels[i])) / 2))
			buf.SetString(
				self.Labels[i],
				AttrPair{SelectAttr(self.LabelAttrs, i), ColorDefault},
				image.Pt(labelXCoordinate, self.Inner.Max.Y-1),
			)
		}

		// draw number
		numberXCoordinate := barXCoordinate + int((float64(self.BarWidth) / 2))
		if numberXCoordinate <= self.Inner.Max.X {
			buf.SetString(
				self.NumFmt(data),
				AttrPair{
					SelectAttr(self.NumAttrs, i+1),
					SelectAttr(self.BarAttrs, i),
				},
				image.Pt(numberXCoordinate, self.Inner.Max.Y-2),
			)
		}

		barXCoordinate += (self.BarWidth + self.BarGap)
	}
}

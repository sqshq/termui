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
	NumFmt     func(int) string
	Data       []int
	Labels     []string
	BarWidth   int
	BarGap     int
	MaxVal     int
}

func NewBarChart() *BarChart {
	return &BarChart{
		Block:      *NewBlock(),
		BarAttrs:   Theme.BarChart.Bars,
		NumAttrs:   Theme.BarChart.Nums,
		LabelAttrs: Theme.BarChart.Labels,
		NumFmt:     func(n int) string { return fmt.Sprint(n) },
		BarGap:     1,
		BarWidth:   3,
	}
}

func (bc *BarChart) Draw(buf *Buffer) {
	bc.Block.Draw(buf)

	maxVal := bc.MaxVal
	if maxVal == 0 {
		maxVal, _ = GetMaxIntFromSlice(bc.Data)
	}

	barXCoordinate := bc.Inner.Min.X

	for i, data := range bc.Data {
		// draw bar
		height := int((float64(data) / float64(maxVal)) * float64(bc.Inner.Dy()-1))
		for x := barXCoordinate; x < MinInt(barXCoordinate+bc.BarWidth, bc.Inner.Max.X); x++ {
			for y := bc.Inner.Max.Y - 2; y > (bc.Inner.Max.Y-2)-height; y-- {
				c := Cell{' ', AttrPair{ColorDefault, SelectAttr(bc.BarAttrs, i)}}
				buf.SetCell(c, image.Pt(x, y))
			}
		}

		// draw label
		if i < len(bc.Labels) {
			labelXCoordinate := barXCoordinate +
				int((float64(bc.BarWidth) / 2)) -
				int((float64(rw.StringWidth(bc.Labels[i])) / 2))
			buf.SetString(
				bc.Labels[i],
				AttrPair{SelectAttr(bc.LabelAttrs, i), ColorDefault},
				image.Pt(labelXCoordinate, bc.Inner.Max.Y-1),
			)
		}

		// draw number
		numberXCoordinate := barXCoordinate + int((float64(bc.BarWidth) / 2))
		if numberXCoordinate <= bc.Inner.Max.X {
			buf.SetString(
				fmt.Sprintf("%d", data),
				AttrPair{
					SelectAttr(bc.NumAttrs, i+1),
					SelectAttr(bc.BarAttrs, i),
				},
				image.Pt(numberXCoordinate, bc.Inner.Max.Y-2),
			)
		}

		barXCoordinate += (bc.BarWidth + bc.BarGap)
	}
}

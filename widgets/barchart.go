// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"fmt"
	"image"

	. "github.com/gizak/termui"
	rw "github.com/mattn/go-runewidth"
)

type BarChart struct {
	Block
	BarColors   []Attribute
	LabelColors []Attribute
	NumColors   []Attribute
	NumFmt      func(int) string
	Data        []int
	Labels      []string
	BarWidth    int
	BarGap      int
	MaxVal      int
}

func NewBarChart() *BarChart {
	return &BarChart{
		Block:       *NewBlock(),
		BarColors:   Theme.BarChart.Bars,
		NumColors:   Theme.BarChart.Nums,
		LabelColors: Theme.BarChart.Labels,
		NumFmt:      func(n int) string { return fmt.Sprint(n) },
		BarGap:      1,
		BarWidth:    3,
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
				c := Cell{' ', AttrPair{ColorDefault, SelectAttr(bc.BarColors, i)}}
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
				image.Pt(labelXCoordinate, bc.Inner.Max.Y-1),
				AttrPair{SelectAttr(bc.LabelColors, i), ColorDefault},
			)
		}

		// draw number
		numberXCoordinate := barXCoordinate + int((float64(bc.BarWidth) / 2))
		buf.SetString(
			fmt.Sprintf("%d", data),
			image.Pt(numberXCoordinate, bc.Inner.Max.Y-2),
			AttrPair{
				SelectAttr(bc.NumColors, i),
				SelectAttr(bc.BarColors, i),
			},
		)

		barXCoordinate += (bc.BarWidth + bc.BarGap)
	}
}

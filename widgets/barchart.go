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
	MaxHeight   int
}

// NewBarChart returns a new *BarChart with current theme.
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

func (bc *BarChart) Draw(buf Buffer) {
	bc.Block.Draw(buf)

	maxHeight := bc.MaxHeight
	if maxHeight == 0 {
		maxHeight, _ = GetMaxIntFromSlice(bc.Data)
	}

	barXCoordinate := bc.Min.X + 1

	for i, data := range bc.Data {
		// draw bar
		height := int((float64(data) / float64(maxHeight)) * float64(bc.Dy()-3))
		for x := barXCoordinate; x < MinInt(barXCoordinate+bc.BarWidth, bc.Max.X-1); x++ {
			for y := bc.Max.Y - 3; y > (bc.Max.Y-3)-height; y-- {
				c := Cell{' ', AttrPair{ColorDefault, SelectAttr(bc.BarColors, i)}}
				buf.SetCell(c, image.Pt(x, y))
			}
		}

		// draw label
		labelXCoordinate := barXCoordinate +
			int((float64(bc.BarWidth) / 2)) -
			int((float64(rw.StringWidth(bc.Labels[i])) / 2))
		buf.SetString(
			bc.Labels[i],
			image.Pt(labelXCoordinate, bc.Max.Y-2),
			AttrPair{SelectAttr(bc.LabelColors, i), ColorDefault},
		)

		// draw number
		numberXCoordinate := barXCoordinate + int((float64(bc.BarWidth) / 2))
		buf.SetString(
			fmt.Sprintf("%d", data),
			image.Pt(numberXCoordinate, bc.Max.Y-3),
			AttrPair{
				SelectAttr(bc.NumColors, i),
				SelectAttr(bc.BarColors, i),
			},
		)

		barXCoordinate += (bc.BarWidth + bc.BarGap)
	}
}

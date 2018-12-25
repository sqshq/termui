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

type StackedBarChart struct {
	Block
	BarColors   []Attribute
	LabelColors []Attribute
	NumColors   []Attribute
	NumFmt      func(int) string
	Data        [][]int
	Labels      []string
	BarWidth    int
	BarGap      int
	MaxHeight   int
}

// NewStackedBarChart returns a new *StackedBarChart with current theme.
func NewStackedBarChart() *StackedBarChart {
	return &StackedBarChart{
		Block:       *NewBlock(),
		BarColors:   Theme.StackedBarChart.Bars,
		LabelColors: Theme.StackedBarChart.Labels,
		NumColors:   Theme.StackedBarChart.Nums,
		NumFmt:      func(n int) string { return fmt.Sprint(n) },
		BarGap:      1,
		BarWidth:    3,
	}
}

// Buffer implements Bufferer interface.
func (bc *StackedBarChart) Buffer() Buffer {
	buf := bc.Block.Buffer()

	maxHeight := bc.MaxHeight
	if maxHeight == 0 {
		for _, data := range bc.Data {
			maxHeight = MaxInt(maxHeight, SumIntSlice(data))
		}
	}

	barXCoordinate := bc.Min.X + 1

	for i, bar := range bc.Data {
		// draw stacked bars
		stackedBarYCoordinate := 0
		for j, data := range bar {
			// draw each stacked bar
			height := int((float64(data) / float64(maxHeight)) * float64(bc.Dy()-3))
			for x := barXCoordinate; x < MinInt(barXCoordinate+bc.BarWidth, bc.Max.X-1); x++ {
				for y := (bc.Max.Y - 3) - stackedBarYCoordinate; y > (bc.Max.Y-3)-stackedBarYCoordinate-height; y-- {
					c := Cell{' ', AttrPair{ColorDefault, SelectAttr(bc.BarColors, j)}}
					buf.SetCell(c, image.Pt(x, y))
				}
			}

			// draw number
			numberXCoordinate := barXCoordinate + int((float64(bc.BarWidth) / 2)) - 1
			buf.SetString(
				fmt.Sprintf("%d", data),
				image.Pt(numberXCoordinate, (bc.Max.Y-3)-stackedBarYCoordinate),
				AttrPair{
					SelectAttr(bc.NumColors, j+1),
					SelectAttr(bc.BarColors, j),
				},
			)

			stackedBarYCoordinate += height
		}

		// draw label
		labelXCoordinate := barXCoordinate + MaxInt(
			int((float64(bc.BarWidth)/2))-int((float64(rw.StringWidth(bc.Labels[i]))/2)),
			0,
		)
		buf.SetString(
			TrimString(bc.Labels[i], bc.BarWidth),
			image.Pt(labelXCoordinate, bc.Max.Y-2),
			AttrPair{SelectAttr(bc.LabelColors, i), ColorDefault},
		)

		barXCoordinate += (bc.BarWidth + bc.BarGap)
	}

	return buf
}

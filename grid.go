// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"image"
)

// Grid layout
var Grid grid

type gridItemType int

const (
	col gridItemType = 0
	row gridItemType = 1
)

// GridItem represents either a Row or Column in a grid and holds sizing information and other GridItems or widgets
type GridItem struct {
	Type        gridItemType
	XRatio      float64
	YRatio      float64
	WidthRatio  float64
	HeightRatio float64
	Entry       interface{} // Entry.type == GridBufferer if IsLeaf else []GridItem
	IsLeaf      bool
	ratio       float64
}

type grid struct {
	Items []*GridItem
	image.Rectangle
}

func newGrid(r image.Rectangle) grid {
	return grid{
		Rectangle: r,
	}
}

// NewCol takes a height percentage and either a widget or a Row or Column
func NewCol(ratio float64, i ...interface{}) GridItem {
	_, ok := i[0].(Drawable)
	entry := i[0]
	if !ok {
		entry = i
	}
	return GridItem{
		Type:   col,
		Entry:  entry,
		IsLeaf: ok,
		ratio:  ratio,
	}
}

// NewRow takes a width percentage and either a widget or a Row or Column
func NewRow(ratio float64, i ...interface{}) GridItem {
	_, ok := i[0].(Drawable)
	entry := i[0]
	if !ok {
		entry = i
	}
	return GridItem{
		Type:   row,
		Entry:  entry,
		IsLeaf: ok,
		ratio:  ratio,
	}
}

// Set recursively searches the GridItems, adding leaves to the grid and calculating the dimensions of the leaves.
func (g *grid) Set(entries ...GridItem) {
	entry := GridItem{
		Type:   row,
		Entry:  entries,
		IsLeaf: false,
		ratio:  1.0,
	}
	g.setHelper(entry, 1.0, 1.0)
}

func (g *grid) setHelper(item GridItem, parentWidthRatio, parentHeightRatio float64) {
	var HeightRatio float64
	var WidthRatio float64
	if item.Type == col {
		HeightRatio = 1.0
		WidthRatio = item.ratio
	} else {
		HeightRatio = item.ratio
		WidthRatio = 1.0
	}
	item.WidthRatio = parentWidthRatio * WidthRatio
	item.HeightRatio = parentHeightRatio * HeightRatio

	if item.IsLeaf {
		g.Items = append(g.Items, &item)
	} else {
		XRatio := 0.0
		YRatio := 0.0
		cols := false
		rows := false

		children := interfaceSlice(item.Entry)

		for i := 0; i < len(children); i++ {
			child, _ := children[i].(GridItem)

			child.XRatio = item.XRatio + ((1 - item.XRatio) * XRatio)
			child.YRatio = item.YRatio + ((1 - item.YRatio) * YRatio)

			if child.Type == col {
				cols = true
				XRatio += child.ratio
				if rows {
					item.HeightRatio /= 2
				}
			} else {
				rows = true
				YRatio += child.ratio
				if cols {
					item.WidthRatio /= 2
				}
			}

			g.setHelper(child, item.WidthRatio, item.HeightRatio)
		}
	}
}

func (g *grid) Draw(rect image.Rectangle, buf Buffer) {
	width := float64(rect.Dx())
	height := float64(rect.Dy())

	for _, item := range g.Items {
		entry, _ := item.Entry.(Drawable)

		x := int(width * item.XRatio)
		y := int(height * item.YRatio)
		w := int(width*item.WidthRatio) - 1
		h := int(height*item.HeightRatio) - 1

		entry.SetRect(x, y, x+w, x+h)

		entry.Draw(&buf)
	}
}

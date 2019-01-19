// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"image"
)

// Cell represents a terminal cell and is a rune with Fg and Bg Attributes
type Cell struct {
	Rune  rune
	Attrs AttrPair
}

// Buffer represents a section of a terminal and is a renderable rectangle cell data container.
type Buffer struct {
	image.Rectangle
	CellMap map[image.Point]Cell
}

func NewBuffer(r image.Rectangle) *Buffer {
	buf := &Buffer{
		Rectangle: r,
		CellMap:   make(map[image.Point]Cell),
	}
	buf.Fill(Cell{' ', AttrPair{ColorDefault, ColorDefault}}, r) // clears out area
	return buf
}

func (self *Buffer) GetCell(p image.Point) Cell {
	return self.CellMap[p]
}

func (self *Buffer) SetCell(c Cell, p image.Point) {
	self.CellMap[p] = c
}

func (self *Buffer) Fill(c Cell, rect image.Rectangle) {
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			self.SetCell(c, image.Pt(x, y))
		}
	}
}

func (self *Buffer) SetString(s string, pair AttrPair, p image.Point) {
	for i, char := range s {
		self.SetCell(Cell{char, pair}, image.Pt(p.X+i, p.Y))
	}
}

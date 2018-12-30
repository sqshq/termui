// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"image"
)

// Attribute is printable cell's color and style.
type Attribute int

// Define basic terminal colors
const (
	// ColorDefault clears the color
	ColorDefault Attribute = iota - 1
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

// These can be bitwise ored to modify cells
const (
	AttrBold Attribute = 1 << (iota + 9)
	AttrUnderline
	AttrReverse
)

// AttrPair holds a cell's Fg and Bg
type AttrPair struct {
	Fg Attribute
	Bg Attribute
}

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
	return &Buffer{
		Rectangle: r,
		CellMap:   make(map[image.Point]Cell),
	}
}

func (b *Buffer) GetCell(p image.Point) Cell {
	return b.CellMap[p]
}

func (b *Buffer) SetCell(c Cell, p image.Point) {
	b.CellMap[p] = c
}

func (b *Buffer) Fill(c Cell, rect image.Rectangle) {
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			b.SetCell(c, image.Pt(x, y))
		}
	}
}

func (b *Buffer) SetString(s string, p image.Point, pair AttrPair) {
	for i, char := range s {
		b.SetCell(Cell{char, pair}, image.Pt(p.X+i, p.Y))
	}
}

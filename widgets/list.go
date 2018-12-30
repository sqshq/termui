// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

type List struct {
	Block
	Rows      []string
	Wrap      bool
	TextAttrs AttrPair
}

func NewList() *List {
	return &List{
		Block:     *NewBlock(),
		TextAttrs: Theme.List.Text,
	}
}

func (l *List) Draw(buf *Buffer) {
	l.Block.Draw(buf)

	point := l.Min.Add(image.Pt(1, 1))

	for row := 0; row < len(l.Rows) && point.Y < l.Max.Y-1; row++ {
		cells := ParseText(l.Rows[row], l.TextAttrs)
		if l.Wrap {
			cells = WrapText(cells, l.Dx()-2)
		}
		for j := 0; j < len(cells) && point.Y < l.Max.Y-1; j++ {
			if cells[j].Rune == '\n' {
				point = image.Pt(l.Min.X+1, point.Y+1)
			} else {
				if point.X+1 == l.Max.X && len(cells) > l.Dx()-2 {
					buf.SetCell(Cell{DOTS, cells[j].Attrs}, point.Add(image.Pt(-1, 0)))
					break
				} else {
					buf.SetCell(cells[j], point)
					point = point.Add(image.Pt(1, 0))
				}
			}
		}
		point = image.Pt(l.Min.X+1, point.Y+1)
	}
}

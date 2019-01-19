// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

type Paragraph struct {
	Block
	Text      string
	TextAttrs AttrPair
}

func NewParagraph() *Paragraph {
	return &Paragraph{
		Block:     *NewBlock(),
		TextAttrs: Theme.Paragraph.Text,
	}
}

func (self *Paragraph) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	point := self.Inner.Min
	cells := WrapText(ParseText(self.Text, self.TextAttrs), self.Inner.Dx())

	for i := 0; i < len(cells) && point.Y < self.Inner.Max.Y; i++ {
		if cells[i].Rune == '\n' {
			point = image.Pt(self.Inner.Min.X, point.Y+1)
		} else {
			buf.SetCell(cells[i], point)
			point = point.Add(image.Pt(1, 0))
		}
	}
}

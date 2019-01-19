// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

type TabPane struct {
	Block
	TabNames         []string
	ActiveTabIndex   int
	ActiveTabAttrs   AttrPair
	InactiveTabAttrs AttrPair
}

func NewTabPane(names ...string) *TabPane {
	return &TabPane{
		Block:            *NewBlock(),
		TabNames:         names,
		ActiveTabAttrs:   Theme.Tab.Active,
		InactiveTabAttrs: Theme.Tab.Inactive,
	}
}

func (self *TabPane) FocusLeft() {
	if self.ActiveTabIndex > 0 {
		self.ActiveTabIndex--
	}
}

func (self *TabPane) FocusRight() {
	if self.ActiveTabIndex < len(self.TabNames)-1 {
		self.ActiveTabIndex++
	}
}

func (self *TabPane) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	xCoordinate := self.Inner.Min.X
	for i, name := range self.TabNames {
		attrPair := self.InactiveTabAttrs
		if i == self.ActiveTabIndex {
			attrPair = self.ActiveTabAttrs
		}
		buf.SetString(
			TrimString(name, self.Inner.Max.X-xCoordinate),
			attrPair,
			image.Pt(xCoordinate, self.Inner.Min.Y),
		)

		xCoordinate += 1 + len(name)

		if i < len(self.TabNames)-1 && xCoordinate < self.Inner.Max.X {
			buf.SetCell(
				Cell{VERTICAL_LINE, AttrPair{ColorWhite, ColorDefault}},
				image.Pt(xCoordinate, self.Inner.Min.Y),
			)
		}

		xCoordinate += 2
	}
}

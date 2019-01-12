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

func (tp *TabPane) FocusLeft() {
	if tp.ActiveTabIndex > 0 {
		tp.ActiveTabIndex--
	}
}

func (tp *TabPane) FocusRight() {
	if tp.ActiveTabIndex < len(tp.TabNames)-1 {
		tp.ActiveTabIndex++
	}
}

func (tp *TabPane) Draw(buf *Buffer) {
	tp.Block.Draw(buf)

	xCoordinate := tp.Inner.Min.X
	for i, name := range tp.TabNames {
		attrPair := tp.InactiveTabAttrs
		if i == tp.ActiveTabIndex {
			attrPair = tp.ActiveTabAttrs
		}
		buf.SetString(
			TrimString(name, tp.Inner.Max.X-xCoordinate),
			attrPair,
			image.Pt(xCoordinate, tp.Inner.Min.Y),
		)

		xCoordinate += 1 + len(name)

		if i < len(tp.TabNames)-1 && xCoordinate < tp.Inner.Max.X {
			buf.SetCell(
				Cell{VERTICAL_LINE, AttrPair{ColorWhite, ColorDefault}},
				image.Pt(xCoordinate, tp.Inner.Min.Y),
			)
		}

		xCoordinate += 2
	}
}

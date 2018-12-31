// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import "image"

type TabPane struct {
	Block
	Tabs             []*Tab
	ActiveTabIndex   int
	ActiveTabAttrs   AttrPair
	InactiveTabAttrs AttrPair
	clearNextDraw    bool
}

type Tab struct {
	Title  string
	Blocks []Drawable
}

func NewTab(title string, blocks ...Drawable) *Tab {
	return &Tab{
		Title:  title,
		Blocks: blocks,
	}
}

func NewTabPane(tabs ...*Tab) *TabPane {
	return &TabPane{
		Block:            *NewBlock(),
		Tabs:             tabs,
		ActiveTabAttrs:   Theme.Tab.Active,
		InactiveTabAttrs: Theme.Tab.Inactive,
	}
}

func (tp *TabPane) FocusLeft() {
	if tp.ActiveTabIndex > 0 {
		tp.ActiveTabIndex--
		tp.clearNextDraw = true
	}
}

func (tp *TabPane) FocusRight() {
	if tp.ActiveTabIndex < len(tp.Tabs)-1 {
		tp.ActiveTabIndex++
		tp.clearNextDraw = true
	}
}

func (tab *Tab) Draw(buf *Buffer) {
	for _, block := range tab.Blocks {
		buf.Rectangle = buf.Rectangle.Union(block.GetRect())
		block.Draw(buf)
	}
}

func (tp *TabPane) Draw(buf *Buffer) {
	tp.Block.Draw(buf)

	// draw tabpane
	xCoordinate := tp.Inner.Min.X
	for i, tab := range tp.Tabs {
		attrPair := tp.InactiveTabAttrs
		if i == tp.ActiveTabIndex {
			attrPair = tp.ActiveTabAttrs
		}
		buf.SetString(
			TrimString(tab.Title, tp.Inner.Max.X-xCoordinate),
			image.Pt(xCoordinate, tp.Inner.Min.Y),
			attrPair,
		)

		xCoordinate += 1 + len(tab.Title)

		if i < len(tp.Tabs)-1 && xCoordinate < tp.Inner.Max.X {
			buf.SetCell(
				Cell{VERTICAL_LINE, AttrPair{ColorWhite, ColorDefault}},
				image.Pt(xCoordinate, tp.Inner.Min.Y),
			)
		}

		xCoordinate += 2
	}

	// draw tab
	if 0 <= tp.ActiveTabIndex && tp.ActiveTabIndex < len(tp.Tabs) {
		tab := tp.Tabs[tp.ActiveTabIndex]
		tab.Draw(buf)
	}

	tp.clearNextDraw = false
}

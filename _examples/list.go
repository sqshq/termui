// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	ls := widgets.NewList()
	ls.Rows = []string{
		"[0] github.com/gizak/termui",
		"[1] [你好，世界]",
		"[2] [こんにちは世界]",
		"[3] [color output]",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go",
	}
	ls.RowAttrs = ui.AttrPair{ui.ColorYellow, ui.ColorDefault}
	ls.Overflow = widgets.ListOverflowWrap
	ls.Title = "List"
	ls.Height = 7
	ls.Width = 25
	ls.Y = 0
	ls.RowAttributes = map[uint]ui.AttrPair{
		1: ui.AttrPair{ui.ColorRed, ui.ColorDefault},
	}

	ui.Render(ls)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

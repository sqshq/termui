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

	rows1 := [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"你好吗", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"2016", "10", "11"},
	}

	table1 := widgets.NewTable()
	table1.Rows = rows1
	table1.FgColor = ui.ColorWhite
	table1.BgColor = ui.ColorDefault
	table1.SetRect(5, 5, 60, 10)

	ui.Render(table1)

	rows2 := [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"2016", "11", "11"},
	}

	table2 := ui.NewTable()
	table2.Rows = rows2
	table2.FgColor = ui.ColorWhite
	table2.BgColor = ui.ColorDefault
	table2.TextAlign = ui.AlignCenter
	table2.Separator = false
	table2.Analysis()
	table2.SetSize()
	table2.BgColors[2] = ui.ColorRed
	table2.Y = 10
	table2.X = 0
	table2.Border = true

	ui.Render(table2)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

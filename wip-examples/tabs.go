// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	ui "github.com/gizak/termui"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	header := ui.NewParagraph("Press q to quit, Press h or l to switch tabs")
	header.SetRect(0, 0, 50, 1)
	header.Border = false
	header.TextBgColor = ui.ColorBlue

	tab1 := ui.NewTab("pierwszy")
	tab2 := ui.NewTab("drugi")
	tab3 := ui.NewTab("trzeci")
	tab4 := ui.NewTab("żółw")
	tab5 := ui.NewTab("four")
	tab6 := ui.NewTab("five")

	p2 := ui.NewParagraph("Press q to quit\nPress h or l to switch tabs\n")
	p2.Title = "Keys"
	p2.SetRect(0, 0, 37, 5)
	p2.BorderFg = ui.ColorYellow

	bc := ui.NewBarChart()
	bc.Title = "Bar Chart"
	bc.Data = []int{3, 2, 5, 3, 9, 5, 3, 2, 5, 8, 3, 2, 4, 5, 3, 2, 5, 7, 5, 3, 2, 6, 7, 4, 6, 3, 6, 7, 8, 3, 6, 4, 5, 3, 2, 4, 6, 4, 8, 5, 9, 4, 3, 6, 5, 3, 6}
	bc.SetRect(0, 0, 26, 10)
	bc.DataLabels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.TextColor = ui.ColorGreen
	bc.BarColor = ui.ColorRed
	bc.NumColor = ui.ColorYellow

	tab1.Add(p2)
	tab2.Add(bc)

	tabpane := ui.NewTabPane()
	tabpane.SetRect(0, 1, 30, 30)
	tabpane.Border = true

	tabpane.SetTabs(*tab1, *tab2, *tab3, *tab4, *tab5, *tab6)

	ui.Render(header, tabpane)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.SetActiveLeft()
			ui.Clear()
			ui.Render(header, tabpane)
		case "l":
			tabpane.SetActiveRight()
			ui.Clear()
			ui.Render(header, tabpane)
		}
	}
}

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

	bc := widgets.NewBarChart()
	bc.Data = []int{3, 2, 5, 3, 9, 3}
	bc.Labels = []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.Title = "Bar Chart"
	bc.SetRect(5, 5, 100, 25)
	bc.BarWidth = 5
	bc.LabelAttrs = []ui.Attribute{ui.ColorBlue}
	bc.BarAttrs = []ui.Attribute{ui.ColorRed, ui.ColorGreen}
	bc.NumAttrs = []ui.Attribute{ui.ColorYellow}

	ui.Render(bc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

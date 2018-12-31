// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	g0 := widgets.NewGauge()
	g0.Title = "Slim Gauge"
	g0.SetRect(20, 20, 30, 30)
	g0.Percent = 75
	g0.BarAttr = ui.ColorRed
	g0.BorderAttrs.Fg = ui.ColorWhite
	g0.TitleAttrs.Fg = ui.ColorCyan

	g2 := widgets.NewGauge()
	g2.Title = "Slim Gauge"
	g2.SetRect(0, 3, 50, 6)
	g2.Percent = 60
	g2.PercentAttr = ui.ColorBlue
	g2.BarAttr = ui.ColorYellow
	g2.BorderAttrs.Fg = ui.ColorWhite

	g1 := widgets.NewGauge()
	g1.Title = "Big Gauge"
	g1.SetRect(0, 6, 50, 11)
	g1.Percent = 30
	g1.PercentAttr = ui.ColorYellow
	g1.BarAttr = ui.ColorGreen
	g1.TitleAttrs.Fg = ui.ColorMagenta
	g1.BorderAttrs.Fg = ui.ColorWhite

	g3 := widgets.NewGauge()
	g3.Title = "Gauge with custom label"
	g3.SetRect(0, 11, 50, 14)
	g3.Percent = 50
	g3.Label = fmt.Sprintf("%v%% (100MBs free)", g3.Percent)

	g4 := widgets.NewGauge()
	g4.Title = "Gauge"
	g4.SetRect(0, 14, 50, 17)
	g4.Percent = 50
	g4.Label = "Gauge with custom highlighted label"
	g4.PercentAttr = ui.ColorYellow
	g4.BarAttr = ui.ColorGreen

	ui.Render(g0, g1, g2, g3, g4)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

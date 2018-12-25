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

	g0 := widgets.NewGauge()
	g0.Percent = 75
	g0.SetRect(5, 5, 50, 10)
	g0.Title = "Slim Gauge"
	g0.BarColor = ui.ColorRed
	g0.BorderAttrs.Fg = ui.ColorWhite
	g0.TitleAttrs.Fg = ui.ColorCyan

	ui.Render(g0)

	// gg := ui.NewBlock()
	// gg.Width = 50
	// gg.Height = 5
	// gg.Y = 12
	// gg.Title = "TEST"

	// g2 := widgets.NewGauge()
	// g2.Percent = 60
	// g2.Width = 50
	// g2.Height = 3
	// g2.PercentColor = ui.ColorBlue
	// g2.Y = 3
	// g2.Title = "Slim Gauge"
	// g2.BarColor = ui.ColorYellow
	// g2.BorderAttrs.Fg = ui.ColorWhite

	// g1 := widgets.NewGauge()
	// g1.Percent = 30
	// g1.Width = 50
	// g1.Height = 5
	// g1.Y = 6
	// g1.Title = "Big Gauge"
	// g1.PercentColor = ui.ColorYellow
	// g1.BarColor = ui.ColorGreen
	// g1.BorderAttrs.Fg = ui.ColorWhite
	// g1.TitleAttrs.Fg = ui.ColorMagenta

	// g3 := widgets.NewGauge()
	// g3.Percent = 50
	// g3.Width = 50
	// g3.Height = 3
	// g3.Y = 11
	// g3.Title = "Gauge with custom label"
	// // g3.Label = "{{percent}}% (100MBs free)"

	// g4 := widgets.NewGauge()
	// g4.Percent = 50
	// g4.Width = 50
	// g4.Height = 3
	// g4.Y = 14
	// g4.Title = "Gauge"
	// g4.Label = "Gauge with custom highlighted label"
	// g4.PercentColor = ui.ColorYellow
	// g4.BarColor = ui.ColorGreen

	// ui.Render(g0, g1, g2, g3, g4)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

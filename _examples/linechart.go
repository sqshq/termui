// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"math"

	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	sinps := (func() [][]float64 {
		n := 220
		ps := make([][]float64, 2)
		ps[0] = make([]float64, n)
		ps[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			ps[0][i] = 1 + math.Sin(float64(i)/5)
			ps[1][i] = 1 + math.Cos(float64(i)/5)
		}
		return ps
	})()

	lc0 := widgets.NewLineChart()
	lc0.Title = "braille-mode Line Chart"
	lc0.Data = sinps
	lc0.SetRect(0, 0, 50, 12)
	lc0.AxesColor = ui.ColorWhite
	lc0.LineColors[0] = ui.ColorGreen | ui.AttrBold

	lc1 := widgets.NewLineChart()
	lc1.Title = "dot-mode Line Chart"
	lc1.LineType = widgets.Dot
	lc1.Data = sinps
	lc1.SetRect(51, 0, 77, 12)
	lc1.DotChar = '+'
	lc1.AxesColor = ui.ColorWhite
	lc1.LineColors[0] = ui.ColorYellow | ui.AttrBold

	lc2 := widgets.NewLineChart()
	lc2.Title = "dot-mode Line Chart"
	lc2.LineType = widgets.Dot
	lc2.Data = make([][]float64, 2)
	lc2.Data[0] = sinps[0][4:]
	lc2.Data[1] = sinps[1][4:]
	lc2.SetRect(0, 12, 77, 28)
	lc2.AxesColor = ui.ColorWhite
	lc2.LineColors[0] = ui.ColorCyan | ui.AttrBold

	ui.Render(lc0, lc1, lc2)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

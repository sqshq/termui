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

	// sinps := (func() []float64 {
	// 	n := 400
	// 	ps := make([]float64, n)
	// 	for i := range ps {
	// 		ps[i] = 1 + math.Sin(float64(i)/5)
	// 	}
	// 	return ps
	// })()
	// sinpsint := (func() []int {
	// 	ps := make([]int, len(sinps))
	// 	for i, v := range sinps {
	// 		ps[i] = int(100*v + 10)
	// 	}
	// 	return ps
	// })()

	// spark := ui.Sparkline{}
	// spark.Height = 8
	// spdata := sinpsint
	// spark.Data = spdata[:100]
	// spark.LineColor = ui.ColorCyan
	// spark.TitleColor = ui.ColorWhite

	// sp := ui.NewSparklines(spark)
	// sp.Height = 11
	// sp.Title = "Sparkline"

	// lc := ui.NewLineChart()
	// lc.Title = "braille-mode Line Chart"
	// lc.Data["default"] = sinps
	// lc.Height = 11
	// lc.AxesColor = ui.ColorWhite
	// lc.LineColor["default"] = ui.ColorYellow | ui.AttrBold

	// gs := make([]*ui.Gauge, 3)
	// for i := range gs {
	// 	gs[i] = ui.NewGauge()
	// 	//gs[i].LabelAlign = ui.AlignCenter
	// 	gs[i].Height = 2
	// 	gs[i].Border = false
	// 	gs[i].Percent = i * 10
	// 	gs[i].PaddingBottom = 1
	// 	gs[i].BarColor = ui.ColorRed
	// }

	// ls := ui.NewList()
	// ls.Border = false
	// ls.Items = []string{
	// 	"[1] Downloading File 1",
	// 	"", // == \newline
	// 	"[2] Downloading File 2",
	// 	"",
	// 	"[3] Uploading File 3",
	// }
	// ls.Height = 5

	// p := ui.NewParagraph("<> This row has 3 columns\n<- Widgets can be stacked up like left side\n<- Stacked widgets are treated as a single widget")
	// p.Height = 5
	// p.Title = "Demonstration"

	// ui.Body.Set(
	// 	ui.NewRow(1/2,
	// 		ui.NewCol(1/2, sp),
	//         ui.NewCol(1/2, lc)
	//     ),
	// 	ui.NewRow(1/2,
	// 		ui.NewCol(1/4, ls),
	//         ui.NewCol(1/4,
	//             ui.NewRow(1/3, gs[0]),
	//             ui.NewRow(1/3, gs[1]),
	//             ui.NewRow(1/3, gs[2]),
	//         ),
	//         ui.NewCol(1/2, p),
	//     )
	// )

	ui.Grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, ui.NewBlock()),
			ui.NewRow(1.0/2, ui.NewBlock()),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/4, ui.NewBlock()),
			ui.NewCol(1.0/4,
				ui.NewRow(1.0/3, ui.NewBlock()),
				ui.NewRow(0.9/3, ui.NewBlock()),
				ui.NewRow(1.1/3, ui.NewBlock()),
			),
			ui.NewCol(1.0/2, ui.NewBlock()),
		),
	)

	// block := ui.NewBlock()
	// block.Title = " hi "
	// ui.Grid.Set(ui.NewRow(.5, block))

	ui.Render(ui.Grid)
	// block.Width = 2
	// block.Height = 2
	// ui.Render(block)

	// tickerCount := 1
	uiEvents := ui.PollEvents()
	// ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				ui.Grid.Width, ui.Grid.Height = payload.Width, payload.Height
				ui.Clear()
				ui.Render(ui.Grid)
			}
		}
	}
	// 	case <-ticker:
	// 		if tickerCount > 103 {
	// 			return
	// 		}
	// 		for _, g := range gs {
	// 			g.Percent = (g.Percent + 3) % 100
	// 		}
	// 		sp.Lines[0].Data = spdata[:100+tickerCount]
	// 		lc.Data["default"] = sinps[2*tickerCount:]
	// 		ui.Render(ui.Body)
	// 		tickerCount++
	// 	}
	// }
}

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

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalSize()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid2 := ui.NewGrid()
	grid2.Set(ui.NewCol(.5, ui.NewBlock()))

	grid.Set(
		ui.NewRow(.5, ui.NewBlock()),
		ui.NewRow(.5, grid2),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Render(grid)
		}
	}
}

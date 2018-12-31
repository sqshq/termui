// +build ignore

package main

import (
	"image"

	ui "github.com/gizak/termui"
)

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	c := ui.NewCanvas()
	c.SetRect(0, 0, 50, 50)
	c.Line(image.Pt(0, 0), image.Pt(80, 50), ui.ColorDefault)
	c.Line(image.Pt(0, 5), image.Pt(3, 10), ui.ColorDefault)

	ui.Render(c)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}

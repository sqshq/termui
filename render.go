// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"sync"

	tb "github.com/nsf/termbox-go"
)

var renderLock sync.Mutex

// Bufferer should be implemented by all renderable components.
type Bufferer interface {
	Buffer() Buffer
}

func Render(bs ...Bufferer) {
	go func() {
		for _, b := range bs {
			buf := b.Buffer()
			for point, cell := range buf.CellMap {
				if point.In(buf.Rectangle) {
					tb.SetCell(point.X, point.Y, cell.Ch, tb.Attribute(cell.Attributes.Fg)+1, tb.Attribute(cell.Attributes.Bg)+1)
				}
			}
		}
		renderLock.Lock()
		tb.Flush()
		renderLock.Unlock()
	}()
}

func Clear() {
	tb.Clear(tb.ColorDefault, tb.Attribute(tb.Attribute(Theme.Default.Bg))+1)
}

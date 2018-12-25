// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"image"

	tb "github.com/nsf/termbox-go"
)

// Init initializes termui library. This function should be called before any others.
// After initialization, the library must be finalized by 'Close' function.
func Init() error {
	if err := tb.Init(); err != nil {
		return err
	}
	tb.SetInputMode(tb.InputEsc | tb.InputMouse)

	w, h := TerminalSize()
	Grid = newGrid(image.Rect(0, 0, w, h))

	return nil
}

// Close finalizes termui library.
// It should be called after successful initialization when termui's functionality isn't required anymore.
func Close() {
	tb.Close()
}

// OutputMode is used for Termbox display modes
type OutputMode int

// Termbox output modes
const (
	OutputCurrent OutputMode = iota
	OutputNormal
	Output256
	Output216
	OutputGrayscale
)

func SetOutputMode(mode OutputMode) {
	switch mode {
	case OutputCurrent:
		tb.SetOutputMode(tb.OutputCurrent)
	case OutputNormal:
		tb.SetOutputMode(tb.OutputNormal)
	case Output256:
		tb.SetOutputMode(tb.Output256)
	case Output216:
		tb.SetOutputMode(tb.Output216)
	case OutputGrayscale:
		tb.SetOutputMode(tb.OutputGrayscale)
	}
}

func TerminalSize() (int, int) {
	renderLock.Lock()
	tb.Sync()
	width, height := tb.Size()
	renderLock.Unlock()
	return width, height
}

// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

type BrailleCanvas struct {
	image.Rectangle
}

var braillePatterns = map[[2]int]rune{
	[2]int{0, 0}: '⣀',
	[2]int{0, 1}: '⡠',
	[2]int{0, 2}: '⡐',
	[2]int{0, 3}: '⡈',

	[2]int{1, 0}: '⢄',
	[2]int{1, 1}: '⠤',
	[2]int{1, 2}: '⠔',
	[2]int{1, 3}: '⠌',

	[2]int{2, 0}: '⢂',
	[2]int{2, 1}: '⠢',
	[2]int{2, 2}: '⠒',
	[2]int{2, 3}: '⠊',

	[2]int{3, 0}: '⢁',
	[2]int{3, 1}: '⠡',
	[2]int{3, 2}: '⠑',
	[2]int{3, 3}: '⠉',
}

var lSingleBraille = [4]rune{'\u2840', '⠄', '⠂', '⠁'}
var rSingleBraille = [4]rune{'\u2880', '⠠', '⠐', '⠈'}

// LineChart has two modes: braille(default) and dot.
// A single braille character is a 2x4 grid of dots, so Using braille
// gives 2x X resolution and 4x Y resolution over dot mode.
/*
  lc := termui.NewLineChart()
  lc.Border.Label = "braille-mode Line Chart"
  lc.Data["name'] = [1.2, 1.3, 1.5, 1.7, 1.5, 1.6, 1.8, 2.0]
  lc.Width = 50
  lc.Height = 12
  lc.AxesColor = termui.ColorWhite
  lc.LineColor = termui.ColorGreen | termui.AttrBold
  // termui.Render(lc)...
*/
type LineChart struct {
	Block
	Data       [][]float64
	DataLabels []string
	Scale      float64
	LineType   LineStyle
	DotChar    rune
	LineColors []Attribute
	AxesColor  Attribute
	MaxHeight  float64
	ShowAxes   bool
}

const dot = '•'

type LineStyle int

const (
	Braille LineStyle = iota
	Dot
)

// NewLineChart returns a new LineChart with current theme.
func NewLineChart() *LineChart {
	return &LineChart{
		Block:      *NewBlock(),
		LineColors: Theme.LineChart.Lines,
		AxesColor:  Theme.LineChart.Axes,
		LineType:   Braille,
		DotChar:    dot,
	}
}

// one cell contains two data points, so capicity is 2x dot mode
func (lc *LineChart) renderBraille(buf Buffer) {
	// maxHeight := lc.MaxHeight
	// if maxHeight == 0 {
	// 	maxHeight, _ = GetMaxFloat64From2dSlice(lc.Data)
	// }

	// for _, line := range lc.Data {
	// 	previousVal := 0
	// 	for j := 0; j < len(line) && j < lc.Dx()-2; j++ {
	// 		val := line[j]
	// 		height := int((val / maxHeight) * float64(lc.Dy()-3))
	// 		for k := 0; k < height; k++ {
	// 			buf.SetCell(
	// 				Cell{lc.DotChar, AttrPair{ColorBlue, ColorDefault}},
	// 				image.Pt(lc.Min.X+j+1, lc.Max.Y-height-2),
	// 			)
	// 		}
	// 	}
	// }
}

func (lc *LineChart) renderDot(buf Buffer) {
	maxHeight := lc.MaxHeight
	if maxHeight == 0 {
		maxHeight, _ = GetMaxFloat64From2dSlice(lc.Data)
	}

	for _, line := range lc.Data {
		for j := 0; j < len(line) && j < lc.Dx()-2; j++ {
			val := line[j]
			height := int((val / maxHeight) * float64(lc.Dy()-2))
			buf.SetCell(
				Cell{lc.DotChar, AttrPair{ColorBlue, ColorDefault}},
				image.Pt(lc.Min.X+j+1, lc.Max.Y-height-2),
			)
		}
	}
}

func (lc *LineChart) plotAxes(buf Buffer) {
}

// Buffer implements Bufferer interface.
func (lc *LineChart) Buffer() Buffer {
	buf := lc.Block.Buffer()

	if lc.ShowAxes {
		lc.plotAxes(buf)
	}

	if lc.LineType == Braille {
		lc.renderBraille(buf)
	} else {
		lc.renderDot(buf)
	}

	return buf
}

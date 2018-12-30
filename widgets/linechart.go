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
// A single braille character is a 2x4 grid of dots, so using braille
// gives 2x X resolution and 4x Y resolution over dot mode.
type LineChart struct {
	Block
	Data       [][]float64
	DataLabels []string
	Scale      float64
	LineType   LineType
	DotChar    rune
	LineColors []Attribute
	AxesColor  Attribute
	MaxVal     float64
	ShowAxes   bool
}

const DOT = '•'

type LineType int

const (
	BrailleLine LineType = iota
	DotLine
)

func NewLineChart() *LineChart {
	return &LineChart{
		Block:      *NewBlock(),
		LineColors: Theme.LineChart.Lines,
		AxesColor:  Theme.LineChart.Axes,
		LineType:   BrailleLine,
		DotChar:    DOT,
	}
}

// one cell contains two data points, so capicity is 2x dot mode
func (lc *LineChart) renderBraille(buf *Buffer) {
}

func (lc *LineChart) renderDot(buf *Buffer) {
	maxVal := lc.MaxVal
	if maxVal == 0 {
		maxVal, _ = GetMaxFloat64From2dSlice(lc.Data)
	}

	for _, line := range lc.Data {
		for j := 0; j < len(line) && j < lc.Dx()-2; j++ {
			val := line[j]
			height := int((val / maxVal) * float64(lc.Dy()-2))
			buf.SetCell(
				Cell{lc.DotChar, AttrPair{ColorBlue, ColorDefault}},
				image.Pt(lc.Min.X+j+1, lc.Max.Y-height-2),
			)
		}
	}
}

func (lc *LineChart) plotAxes(buf *Buffer) {
}

func (lc *LineChart) Draw(buf *Buffer) {
	lc.Block.Draw(buf)

	if lc.ShowAxes {
		lc.plotAxes(buf)
	}

	if lc.LineType == BrailleLine {
		lc.renderBraille(buf)
	} else {
		lc.renderDot(buf)
	}
}

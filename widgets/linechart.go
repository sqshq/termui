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

// LineChart has two modes: braille(default) and dot.
// A single braille character is a 2x4 grid of dots, so using braille
// gives 2x X resolution and 4x Y resolution over dot mode.
type LineChart struct {
	Block
	Data          [][]float64
	DataLabels    []string
	Scale         int
	LineType      LineType
	DotChar       rune
	LineAttrs     []Attribute
	AxesAttr      Attribute
	MaxVal        float64
	ShowAxes      bool
	DrawDirection DrawDirection
}

type LineType int

const (
	BrailleLine LineType = iota
	DotLine
)

type DrawDirection int

const (
	DrawLeft DrawDirection = iota
	DrawRight
)

func NewLineChart() *LineChart {
	return &LineChart{
		Block:         *NewBlock(),
		LineAttrs:     Theme.LineChart.Lines,
		AxesAttr:      Theme.LineChart.Axes,
		LineType:      BrailleLine,
		DotChar:       DOT,
		Data:          [][]float64{},
		Scale:         3, // TODO
		DrawDirection: DrawRight,
	}
}

// one cell contains two data points, so capicity is 2x dot mode
func (lc *LineChart) renderBraille(buf *Buffer) {
	canvas := NewCanvas()
	canvas.Rectangle = lc.Inner

	maxVal := lc.MaxVal
	if maxVal == 0 {
		maxVal, _ = GetMaxFloat64From2dSlice(lc.Data)
	}

	for i, line := range lc.Data {
		previousHeight := int((line[1] / maxVal) * float64(lc.Inner.Dy()-1))
		for j, val := range line[1:] {
			height := int((val / maxVal) * float64(lc.Inner.Dy()-1))
			canvas.Line(
				image.Pt(
					(lc.Inner.Min.X+(j*lc.Scale))*2,
					(lc.Inner.Max.Y-previousHeight-1)*4,
				),
				image.Pt(
					(lc.Inner.Min.X+((j+1)*lc.Scale))*2,
					(lc.Inner.Max.Y-height-1)*4,
				),
				SelectAttr(lc.LineAttrs, i),
			)
			previousHeight = height
		}
	}

	canvas.Draw(buf)
}

func (lc *LineChart) renderDot(buf *Buffer) {
	maxVal := lc.MaxVal
	if maxVal == 0 {
		maxVal, _ = GetMaxFloat64From2dSlice(lc.Data)
	}

	for i, line := range lc.Data {
		for j := 0; j < len(line) && j*lc.Scale < lc.Inner.Dx(); j++ {
			val := line[j]
			height := int((val / maxVal) * float64(lc.Inner.Dy()-1))
			buf.SetCell(
				Cell{lc.DotChar, AttrPair{SelectAttr(lc.LineAttrs, i), ColorDefault}},
				image.Pt(lc.Inner.Min.X+(j*lc.Scale), lc.Inner.Max.Y-1-height),
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

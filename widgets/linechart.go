// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"fmt"
	"image"

	. "github.com/gizak/termui"
)

// LineChart has two modes: braille(default) and dot.
// A single braille character is a 2x4 grid of dots, so using braille
// gives 2x X resolution and 4x Y resolution over dot mode.
type LineChart struct {
	Block
	Data            [][]float64
	DataLabels      []string
	HorizontalScale int
	LineType        LineType
	DotChar         rune
	LineAttrs       []Attribute
	AxesAttr        Attribute
	MaxVal          float64
	ShowAxes        bool
	DrawDirection   DrawDirection
}

const (
	yAxesWidth  = 4
	xAxesHeight = 1
	yAxesGap    = 1
	xAxesGap    = 2
)

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
		Block:           *NewBlock(),
		LineAttrs:       Theme.LineChart.Lines,
		AxesAttr:        Theme.LineChart.Axes,
		LineType:        BrailleLine,
		DotChar:         DOT,
		Data:            [][]float64{},
		HorizontalScale: 3, // TODO
		DrawDirection:   DrawRight,
		ShowAxes:        true,
	}
}

// one cell contains two data points, so capicity is 2x dot mode
func (lc *LineChart) renderBraille(buf *Buffer, drawArea image.Rectangle, maxVal float64) {
	canvas := NewCanvas()
	canvas.Rectangle = drawArea

	for i, line := range lc.Data {
		previousHeight := int((line[1] / maxVal) * float64(drawArea.Dy()-1))
		for j, val := range line[1:] {
			height := int((val / maxVal) * float64(drawArea.Dy()-1))
			canvas.Line(
				image.Pt(
					(drawArea.Min.X+(j*lc.HorizontalScale))*2,
					(drawArea.Max.Y-previousHeight-1)*4,
				),
				image.Pt(
					(drawArea.Min.X+((j+1)*lc.HorizontalScale))*2,
					(drawArea.Max.Y-height-1)*4,
				),
				SelectAttr(lc.LineAttrs, i),
			)
			previousHeight = height
		}
	}

	canvas.Draw(buf)
}

func (lc *LineChart) renderDot(buf *Buffer, drawArea image.Rectangle, maxVal float64) {
	for i, line := range lc.Data {
		for j := 0; j < len(line) && j*lc.HorizontalScale < drawArea.Dx(); j++ {
			val := line[j]
			height := int((val / maxVal) * float64(drawArea.Dy()-1))
			buf.SetCell(
				Cell{lc.DotChar, AttrPair{SelectAttr(lc.LineAttrs, i), ColorDefault}},
				image.Pt(drawArea.Min.X+(j*lc.HorizontalScale), drawArea.Max.Y-1-height),
			)
		}
	}
}

func (lc *LineChart) plotAxes(buf *Buffer, maxVal float64) {
	// draw origin
	buf.SetCell(
		Cell{BOTTOM_LEFT, AttrPair{ColorWhite, ColorDefault}},
		image.Pt(lc.Inner.Min.X+yAxesWidth, lc.Inner.Max.Y-xAxesHeight-1),
	)
	// draw x axes line
	for i := yAxesWidth + 1; i < lc.Inner.Dx(); i++ {
		buf.SetCell(
			Cell{HORIZONTAL_DASH, AttrPair{ColorWhite, ColorDefault}},
			image.Pt(i+lc.Inner.Min.X, lc.Inner.Max.Y-xAxesHeight-1),
		)
	}
	// draw y axes line
	for i := 0; i < lc.Inner.Dy()-xAxesHeight-1; i++ {
		buf.SetCell(
			Cell{VERTICAL_DASH, AttrPair{ColorWhite, ColorDefault}},
			image.Pt(lc.Inner.Min.X+yAxesWidth, i+lc.Inner.Min.Y),
		)
	}
	// draw x axes labels
	for i := 0; (i*lc.HorizontalScale)+1 < lc.Inner.Dx()-yAxesWidth-1; i++ {
		buf.SetString(
			fmt.Sprintf("%d", i),
			AttrPair{ColorWhite, ColorDefault},
			image.Pt(lc.Inner.Min.X+yAxesWidth+(i*lc.HorizontalScale), lc.Inner.Max.Y-1),
		)
	}
	// draw y axes labels
	verticalScale := maxVal / float64(lc.Inner.Dy()-xAxesHeight-1)
	for i := 0; i*(yAxesGap+1) < lc.Inner.Dy()-1; i++ {
		buf.SetString(
			fmt.Sprintf("%.2f", float64(i)*verticalScale*(yAxesGap+1)),
			AttrPair{ColorWhite, ColorDefault},
			image.Pt(lc.Inner.Min.X, lc.Inner.Max.Y-(i*(yAxesGap+1))-2),
		)
	}
}

func (lc *LineChart) Draw(buf *Buffer) {
	lc.Block.Draw(buf)

	maxVal := lc.MaxVal
	if maxVal == 0 {
		maxVal, _ = GetMaxFloat64From2dSlice(lc.Data)
	}

	if lc.ShowAxes {
		lc.plotAxes(buf, maxVal)
	}

	drawArea := lc.Inner
	if lc.ShowAxes {
		drawArea = image.Rect(
			lc.Inner.Min.X+yAxesWidth+1, lc.Inner.Min.Y,
			lc.Inner.Max.X, lc.Inner.Max.Y-xAxesHeight-1,
		)
	}

	if lc.LineType == BrailleLine {
		lc.renderBraille(buf, drawArea, maxVal)
	} else {
		lc.renderDot(buf, drawArea, maxVal)
	}
}

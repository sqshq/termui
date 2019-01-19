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
	yAxisWidth              = 4
	xAxisHeight             = 1
	yAxisGap                = 1
	xAxisGap                = 2
	horizontalAxisLabelsGap = 2
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
		HorizontalScale: 1,
		DrawDirection:   DrawRight,
		ShowAxes:        true,
	}
}

// one cell contains two data points, so capicity is 2x dot mode
func (self *LineChart) renderBraille(buf *Buffer, drawArea image.Rectangle, maxVal float64) {
	canvas := NewCanvas()
	canvas.Rectangle = drawArea

	for i, line := range self.Data {
		previousHeight := int((line[1] / maxVal) * float64(drawArea.Dy()-1))
		for j, val := range line[1:] {
			height := int((val / maxVal) * float64(drawArea.Dy()-1))
			canvas.Line(
				image.Pt(
					(drawArea.Min.X+(j*self.HorizontalScale))*2,
					(drawArea.Max.Y-previousHeight-1)*4,
				),
				image.Pt(
					(drawArea.Min.X+((j+1)*self.HorizontalScale))*2,
					(drawArea.Max.Y-height-1)*4,
				),
				SelectAttr(self.LineAttrs, i),
			)
			previousHeight = height
		}
	}

	canvas.Draw(buf)
}

func (self *LineChart) renderDot(buf *Buffer, drawArea image.Rectangle, maxVal float64) {
	for i, line := range self.Data {
		for j := 0; j < len(line) && j*self.HorizontalScale < drawArea.Dx(); j++ {
			val := line[j]
			height := int((val / maxVal) * float64(drawArea.Dy()-1))
			buf.SetCell(
				Cell{self.DotChar, AttrPair{SelectAttr(self.LineAttrs, i), ColorDefault}},
				image.Pt(drawArea.Min.X+(j*self.HorizontalScale), drawArea.Max.Y-1-height),
			)
		}
	}
}

func (self *LineChart) plotAxes(buf *Buffer, maxVal float64) {
	// draw origin
	buf.SetCell(
		Cell{BOTTOM_LEFT, AttrPair{ColorWhite, ColorDefault}},
		image.Pt(self.Inner.Min.X+yAxisWidth, self.Inner.Max.Y-xAxisHeight-1),
	)
	// draw x axis line
	for i := yAxisWidth + 1; i < self.Inner.Dx(); i++ {
		buf.SetCell(
			Cell{HORIZONTAL_DASH, AttrPair{ColorWhite, ColorDefault}},
			image.Pt(i+self.Inner.Min.X, self.Inner.Max.Y-xAxisHeight-1),
		)
	}
	// draw y axis line
	for i := 0; i < self.Inner.Dy()-xAxisHeight-1; i++ {
		buf.SetCell(
			Cell{VERTICAL_DASH, AttrPair{ColorWhite, ColorDefault}},
			image.Pt(self.Inner.Min.X+yAxisWidth, i+self.Inner.Min.Y),
		)
	}
	// draw x axis labels
	// draw 0
	buf.SetString(
		"0",
		AttrPair{ColorWhite, ColorDefault},
		image.Pt(self.Inner.Min.X+yAxisWidth, self.Inner.Max.Y-1),
	)
	// draw rest
	for x := self.Inner.Min.X + yAxisWidth + (horizontalAxisLabelsGap)*self.HorizontalScale + 1; x < self.Inner.Max.X-1; {
		label := fmt.Sprintf(
			"%d",
			(x-(self.Inner.Min.X+yAxisWidth)-1)/(self.HorizontalScale)+1,
		)
		buf.SetString(
			label,
			AttrPair{ColorWhite, ColorDefault},
			image.Pt(x, self.Inner.Max.Y-1),
		)
		x += (len(label) + horizontalAxisLabelsGap) * self.HorizontalScale
	}
	// draw y axis labels
	verticalScale := maxVal / float64(self.Inner.Dy()-xAxisHeight-1)
	for i := 0; i*(yAxisGap+1) < self.Inner.Dy()-1; i++ {
		buf.SetString(
			fmt.Sprintf("%.2f", float64(i)*verticalScale*(yAxisGap+1)),
			AttrPair{ColorWhite, ColorDefault},
			image.Pt(self.Inner.Min.X, self.Inner.Max.Y-(i*(yAxisGap+1))-2),
		)
	}
}

func (self *LineChart) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	maxVal := self.MaxVal
	if maxVal == 0 {
		maxVal, _ = GetMaxFloat64From2dSlice(self.Data)
	}

	if self.ShowAxes {
		self.plotAxes(buf, maxVal)
	}

	drawArea := self.Inner
	if self.ShowAxes {
		drawArea = image.Rect(
			self.Inner.Min.X+yAxisWidth+1, self.Inner.Min.Y,
			self.Inner.Max.X, self.Inner.Max.Y-xAxisHeight-1,
		)
	}

	if self.LineType == BrailleLine {
		self.renderBraille(buf, drawArea, maxVal)
	} else {
		self.renderDot(buf, drawArea, maxVal)
	}
}

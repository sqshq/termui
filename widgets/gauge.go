// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"fmt"
	"image"

	. "github.com/gizak/termui"
)

type Gauge struct {
	Block
	Percent     int
	BarAttr     Attribute
	PercentAttr Attribute
	Label       string
}

func NewGauge() *Gauge {
	return &Gauge{
		Block:       *NewBlock(),
		BarAttr:     Theme.Gauge.Bar,
		PercentAttr: Theme.Gauge.Percent,
	}
}

func (self *Gauge) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	label := self.Label
	if label == "" {
		label = fmt.Sprintf("%d%%", self.Percent)
	}

	// plot bar
	barWidth := int((float64(self.Percent) / 100) * float64(self.Inner.Dx()))
	buf.Fill(
		Cell{' ', AttrPair{ColorDefault, self.BarAttr}},
		image.Rect(self.Inner.Min.X, self.Inner.Min.Y, self.Inner.Min.X+barWidth, self.Inner.Max.Y),
	)

	// plot label
	labelXCoordinate := self.Inner.Min.X + (self.Inner.Dx() / 2) - int(float64(len(label))/2)
	labelYCoordinate := self.Inner.Min.Y + ((self.Inner.Dy() - 1) / 2)
	if labelYCoordinate < self.Inner.Max.Y {
		for i, char := range label {
			attrs := AttrPair{self.PercentAttr, ColorDefault}
			if labelXCoordinate+i+1 <= self.Inner.Min.X+barWidth {
				attrs = AttrPair{self.BarAttr, AttrReverse}
			}
			buf.SetCell(Cell{char, attrs}, image.Pt(labelXCoordinate+i, labelYCoordinate))
		}
	}
}

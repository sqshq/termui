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

func (g *Gauge) Draw(buf *Buffer) {
	g.Block.Draw(buf)

	label := g.Label
	if label == "" {
		label = fmt.Sprintf("%d%%", g.Percent)
	}

	// plot bar
	barWidth := int((float64(g.Percent) / 100) * float64(g.Inner.Dx()))
	buf.Fill(
		Cell{' ', AttrPair{ColorDefault, g.BarAttr}},
		image.Rect(g.Inner.Min.X, g.Inner.Min.Y, g.Inner.Min.X+barWidth, g.Inner.Max.Y),
	)

	// plot label
	labelXCoordinate := g.Inner.Min.X + (g.Inner.Dx() / 2) - int(float64(len(label))/2)
	labelYCoordinate := g.Inner.Min.Y + ((g.Inner.Dy() - 1) / 2)
	if labelYCoordinate < g.Inner.Max.Y {
		for i, char := range label {
			attrs := AttrPair{g.PercentAttr, ColorDefault}
			if labelXCoordinate+i+1 <= g.Inner.Min.X+barWidth {
				attrs = AttrPair{g.BarAttr, AttrReverse}
			}
			buf.SetCell(Cell{char, attrs}, image.Pt(labelXCoordinate+i, labelYCoordinate))
		}
	}
}

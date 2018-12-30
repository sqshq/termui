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
	Percent      int
	BarColor     Attribute
	PercentColor Attribute
	Label        string
}

func NewGauge() *Gauge {
	return &Gauge{
		Block:        *NewBlock(),
		BarColor:     Theme.Gauge.Bar,
		PercentColor: Theme.Gauge.Percent,
	}
}

func (g *Gauge) Draw(buf *Buffer) {
	g.Block.Draw(buf)

	label := g.Label
	if label == "" {
		label = fmt.Sprintf("%d%%", g.Percent)
	}

	// plot bar
	barWidth := int((float64(g.Percent) / 100) * float64(g.Dx()-2))
	buf.Fill(
		Cell{' ', AttrPair{ColorDefault, g.BarColor}},
		image.Rect(g.Min.X+1, g.Min.Y+1, g.Min.X+1+barWidth, g.Max.Y-1),
	)

	// plot label
	labelXCoordinate := (g.Min.X + 1) + ((g.Dx() - 2) / 2) - int(float64(len(label))/2)
	labelYCoordinate := (g.Min.Y + 1) + ((g.Dy() - 3) / 2)
	for i, char := range label {
		attrs := AttrPair{g.PercentColor, ColorDefault}
		if labelXCoordinate+i < g.Min.X+barWidth {
			attrs = AttrPair{g.BarColor, AttrReverse}
		}
		buf.SetCell(Cell{char, attrs}, image.Pt(labelXCoordinate+i, labelYCoordinate))
	}
}

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

// NewGauge return a new gauge with current theme.
func NewGauge() *Gauge {
	return &Gauge{
		Block:        *NewBlock(),
		BarColor:     Theme.Gauge.Bar,
		PercentColor: Theme.Gauge.Percent,
	}
}

// Buffer implements Bufferer interface.
func (g *Gauge) Buffer() Buffer {
	buf := g.Block.Buffer()

	label := g.Label
	if label == "" {
		label = fmt.Sprintf("%d%%", g.Percent)
	}

	// plot bar
	barWidth := int((float64(g.Percent) / 100) * float64(g.Dx()-2))
	for x := g.Min.X + 1; x < g.Min.X+barWidth+1; x++ {
		for y := g.Min.Y + 1; y < g.Max.Y-1; y++ {
			buf.SetCell(Cell{' ', AttrPair{ColorDefault, g.BarColor}}, image.Pt(x, y))
		}
	}

	// plot label
	labelXCoordinate := g.Min.X + (g.Dx() / 2) - (len(label) / 2)
	labelYCoordinate := g.Min.Y + (g.Dy() / 2)
	for i, char := range label {
		attrs := AttrPair{g.PercentColor, g.BarColor}
		if labelXCoordinate+i < barWidth {
			attrs = AttrPair{g.BarColor, AttrReverse}
		}
		buf.SetCell(Cell{char, attrs}, image.Pt(labelXCoordinate+i, labelYCoordinate))
	}

	return buf
}

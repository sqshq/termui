// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	rw "github.com/mattn/go-runewidth"

	. "github.com/gizak/termui"
)

type List struct {
	Block
	Rows          []string
	RowAttributes map[uint]AttrPair
	Overflow      ListOverflow
	RowAttrs      AttrPair
}

// NewList returns a new *List with current theme.
func NewList() *List {
	return &List{
		Block:    *NewBlock(),
		Overflow: ListOverflowHidden,
		RowAttrs: Theme.List.RowAttrs,
	}
}

type ListOverflow uint

const (
	ListOverflowWrap ListOverflow = iota
	ListOverflowHidden
)

// Buffer implements Bufferer interface.
func (l *List) Buffer() Buffer {
	buf := l.Block.Buffer()

	switch l.Overflow {
	case ListOverflowWrap:
		yCoordinate := 0
		for i := 0; yCoordinate < l.Height-1 && i < len(l.Rows); i++ {
			row := l.Rows[i]
			attrs := l.RowAttrs
			if entry, ok := l.RowAttributes[uint(i)]; ok {
				attrs = entry
			}
			if rw.StringWidth(row) > l.Width-1 {
				buf.SetString(row[:l.Width-1], image.Pt(1, yCoordinate+1), attrs)
				yCoordinate++
				buf.SetString(row[l.Width-1:], image.Pt(1, yCoordinate+1), attrs)
			} else {
				buf.SetString(row, image.Pt(1, yCoordinate+1), attrs)
			}
			yCoordinate++
		}

	case ListOverflowHidden:
		for i := 0; i < l.Height-1 && i < len(l.Rows); i++ {
			row := l.Rows[i]
			attrs := l.RowAttrs
			if entry, ok := l.RowAttributes[uint(i)]; ok {
				attrs = entry
			}
			buf.SetString(TrimString(row, l.Width-1), image.Pt(1, i+1), attrs)
		}
	}

	return buf
}

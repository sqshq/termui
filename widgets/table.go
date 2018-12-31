// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package widgets

import (
	"image"

	. "github.com/gizak/termui"
)

/* Table is like:

┌Awesome Table ────────────────────────────────────────────────┐
│  Col0          | Col1 | Col2 | Col3  | Col4  | Col5  | Col6  |
│──────────────────────────────────────────────────────────────│
│  Some Item #1  | AAA  | 123  | CCCCC | EEEEE | GGGGG | IIIII |
│──────────────────────────────────────────────────────────────│
│  Some Item #2  | BBB  | 456  | DDDDD | FFFFF | HHHHH | JJJJJ |
└──────────────────────────────────────────────────────────────┘
*/

type Table struct {
	Block
	Rows         [][]string
	ColumnWidths []int
	TextAttrs    AttrPair
	RowSeparator bool
	TextAlign    Alignment
}

func NewTable() *Table {
	return &Table{
		Block:        *NewBlock(),
		TextAttrs:    Theme.Table.Text,
		RowSeparator: true,
	}
}

func (t *Table) Draw(buf *Buffer) {
	t.Block.Draw(buf)

	columnWidths := t.ColumnWidths
	if len(columnWidths) == 0 {
		columnCount := len(t.Rows[0])
		colWidth := t.Inner.Dx() / columnCount
		for i := 0; i < columnCount; i++ {
			columnWidths = append(columnWidths, colWidth)
		}
	}

	yCoordinate := t.Inner.Min.Y

	for i := 0; i < len(t.Rows) && yCoordinate < t.Inner.Max.Y; i++ {
		row := t.Rows[i]
		xCoordinate := t.Inner.Min.X
		for j := 0; j < len(row); j++ {
			col := ParseText(row[j], t.TextAttrs)
			for k, cell := range col {
				if k == columnWidths[j] || xCoordinate+k == t.Inner.Max.X {
					cell.Rune = DOTS
					buf.SetCell(cell, image.Pt(xCoordinate+k-1, yCoordinate))
					break
				} else {
					buf.SetCell(cell, image.Pt(xCoordinate+k, yCoordinate))
				}
			}
			xCoordinate += columnWidths[j] + 1
		}

		// draw vertical separators
		xCoordinate = t.Inner.Min.X
		verticalCell := Cell{VERTICAL_LINE, AttrPair{ColorWhite, ColorDefault}}
		for j := 0; j < len(columnWidths)-1; j++ {
			xCoordinate += columnWidths[j]
			buf.SetCell(verticalCell, image.Pt(xCoordinate, yCoordinate))
			xCoordinate++
		}

		yCoordinate++

		// draw horizontal separator
		horizontalCell := Cell{HORIZONTAL_LINE, AttrPair{ColorWhite, ColorDefault}}
		if t.RowSeparator && yCoordinate < t.Inner.Max.Y && i != len(t.Rows)-1 {
			buf.Fill(horizontalCell, image.Rect(t.Inner.Min.X, yCoordinate, t.Inner.Max.X, yCoordinate+1))
			yCoordinate++
		}
	}
}

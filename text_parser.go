// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"strings"

	wordwrap "github.com/mitchellh/go-wordwrap"
)

const (
	tokenFg       = "fg"
	tokenBg       = "bg"
	tokenModifier = "mod"

	tokenItemSeparator  = ","
	tokenValueSeparator = ":"

	tokenBeginStyledText = '['
	tokenEndStyledText   = ']'

	tokenBeginStyleItems = '('
	tokenEndStyleItems   = ')'
)

type parserState uint

const (
	parserStateDefault parserState = iota
	parserStateStyleItems
	parserStateStyledText
)

var colorMap = map[string]Color{
	"red":     ColorRed,
	"blue":    ColorBlue,
	"black":   ColorBlack,
	"cyan":    ColorCyan,
	"yellow":  ColorYellow,
	"white":   ColorWhite,
	"clear":   ColorClear,
	"green":   ColorGreen,
	"magenta": ColorMagenta,
}

var modifierMap = map[string]Modifier{
	"bold":      ModifierBold,
	"underline": ModifierUnderline,
	"reverse":   ModifierReverse,
}

// AddColorMap allows users to add/override the string to attribute mapping
func AddColorMap(str string, color Color) {
	colorMap[str] = color
}

// readAttributes translates an []rune like `fg:red,mod:bold,bg:white` to a style
func readAttributes(runes []rune, defaultStyle Style) Style {
	style := defaultStyle
	split := strings.Split(string(runes), tokenItemSeparator)
	for _, item := range split {
		pair := strings.Split(item, tokenValueSeparator)
		if len(pair) == 2 {
			switch pair[0] {
			case tokenFg:
				style.Fg = colorMap[pair[1]]
			case tokenBg:
				style.Bg = colorMap[pair[1]]
			case tokenModifier:
				style.Modifier = modifierMap[pair[1]]
			}
		}
	}
	return style
}

func ParseText(s string, defaultStyle Style) []Cell {
	cells := []Cell{}
	runes := []rune(s)
	state := parserStateDefault
	styledText := []rune{}
	styleItems := []rune{}
	squareCount := 0

	runesToStyledCells := func(runes []rune, style Style) []Cell {
		_cells := []Cell{}
		for _, _rune := range runes {
			_cells = append(_cells, Cell{_rune, style})
		}
		return _cells
	}

	reset := func() {
		styledText = []rune{}
		styleItems = []rune{}
		state = parserStateDefault
		squareCount = 0
	}

	rollback := func() {
		cells = append(cells, runesToStyledCells(styledText, defaultStyle)...)
		cells = append(cells, runesToStyledCells(styleItems, defaultStyle)...)
		reset()
	}

	// chop first and last runes
	chop := func(s []rune) []rune {
		return s[1 : len(s)-1]
	}

	for i, _rune := range runes {
		switch state {
		case parserStateDefault:
			if _rune == tokenBeginStyledText {
				state = parserStateStyledText
				squareCount = 1
				styledText = append(styledText, tokenBeginStyledText)
			} else {
				cells = append(cells, Cell{_rune, defaultStyle})
			}
		case parserStateStyledText:
			switch {
			case squareCount == 0:
				switch _rune {
				case tokenBeginStyleItems:
					state = parserStateStyleItems
					styleItems = append(styleItems, _rune)
				default:
					rollback()
					switch _rune {
					case tokenBeginStyledText:
						state = parserStateStyledText
						squareCount = 1
						styleItems = append(styleItems, _rune)
					default:
						cells = append(cells, Cell{_rune, defaultStyle})
					}
				}
			// hit the end
			case len(runes) == i+1:
				styledText = append(styledText, _rune)
				rollback()
			case _rune == tokenBeginStyledText:
				squareCount++
				styledText = append(styledText, _rune)
			case _rune == tokenEndStyledText:
				squareCount--
				styledText = append(styledText, _rune)
			// normal rune
			default:
				styledText = append(styledText, _rune)
			}
		case parserStateStyleItems:
			styleItems = append(styleItems, _rune)
			if _rune == tokenEndStyleItems {
				style := readAttributes(chop(styleItems), defaultStyle)
				cells = append(cells, runesToStyledCells(chop(styledText), style)...)
				reset()
			} else if len(runes) == i+1 {
				rollback()
			}
		}
	}

	return cells
}

func WrapText(cs []Cell, wl int) []Cell {
	tmpCell := make([]Cell, len(cs))
	copy(tmpCell, cs)

	// get the plaintext
	plain := CellsToString(cs)

	// wrap
	plainWrapped := wordwrap.WrapString(plain, uint(wl))

	// find differences and insert
	finalCell := tmpCell // finalcell will get the inserts and is what is returned

	plainRune := []rune(plain)
	plainWrappedRune := []rune(plainWrapped)
	trigger := "go"
	plainRuneNew := plainRune

	for trigger != "stop" {
		plainRune = plainRuneNew
		for i := range plainRune {
			if plainRune[i] == plainWrappedRune[i] {
				trigger = "stop"
			} else if plainRune[i] != plainWrappedRune[i] && plainWrappedRune[i] == 10 {
				trigger = "go"
				cell := NewCell(10)
				j := i - 0

				// insert a cell into the []Cell in correct position
				tmpCell[i] = cell

				// insert the newline into plain so we avoid indexing errors
				plainRuneNew = append(plainRune, 10)
				copy(plainRuneNew[j+1:], plainRuneNew[j:])
				plainRuneNew[j] = plainWrappedRune[j]

				// restart the inner for loop until plain and plain wrapped are
				// the same; yeah, it's inefficient, but the text amounts
				// should be small
				break

			} else if plainRune[i] != plainWrappedRune[i] &&
				plainWrappedRune[i-1] == 10 && // if the prior rune is a newline
				plainRune[i] == 32 { // and this rune is a space
				trigger = "go"
				// need to delete plainRune[i] because it gets rid of an extra
				// space
				plainRuneNew = append(plainRune[:i], plainRune[i+1:]...)
				break

			} else {
				trigger = "stop" // stops the outer for loop
			}
		}
	}

	finalCell = tmpCell

	return finalCell
}
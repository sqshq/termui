package termui

// Color is an integer from -1 to 255
type Color int

// Basic terminal colors
const (
	ColorClear Color = iota - 1
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

type Modifier uint

const (
	ModifierClear Modifier = 0
	ModifierBold  Modifier = 1 << (iota + 9)
	ModifierUnderline
	ModifierReverse
)

// Style represents the look of the text of one terminal cell
type Style struct {
	Fg       Color
	Bg       Color
	Modifier Modifier
}

var StyleClear = Style{
	Fg:       -1,
	Bg:       -1,
	Modifier: 0,
}

func NewStyle(fg Color, args ...interface{}) Style {
	bg := ColorClear
	modifier := ModifierClear
	if len(args) >= 1 {
		bg = args[0].(Color)
	}
	if len(args) >= 2 {
		modifier = args[1].(Modifier)
	}
	return Style{
		fg,
		bg,
		modifier,
	}
}

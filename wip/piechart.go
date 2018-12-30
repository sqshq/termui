package widgets

import (
	"image"
	"math"

	. "github.com/gizak/termui"
)

const (
	piechartOffsetUp = -.5 * math.Pi // the northward angle
	resolutionFactor = .001          // circle resolution: precision vs. performance
	fullCircle       = 2.0 * math.Pi // the full circle angle
	xStretch         = 2.0           // horizontal adjustment
	solidBlock       = '░'
)

// PieChartLabel callback
type PieChartLabel func(dataIndex int, currentValue float64) string

type PieChart struct {
	Block
	Data   []float64     // list of data items
	Colors []Attribute   // colors to by cycled through (see defaultColors)
	Label  PieChartLabel // callback function for labels
	Offset float64       // which angle to start drawing at? (see piechartOffsetUp)
}

// NewPieChart Creates a new pie chart with reasonable defaults and no labels.
func NewPieChart() *PieChart {
	return &PieChart{
		Block:  *NewBlock(),
		Colors: Theme.PieChart,
		Offset: piechartOffsetUp,
	}
}

// Buffer creates the buffer for the pie chart.
func (pc *PieChart) Buffer() Buffer {
	buf := pc.Block.Buffer()

	center := pc.Min.Add(pc.Size().Div(2))
	radius := minFloat64(float64(pc.Dx()/2/xStretch), float64(pc.Dy()/2)) - 1

	// compute slice sizes
	sum := sumFloat64(pc.Data)
	sliceSizes := make([]float64, len(pc.Data))
	for i, v := range pc.Data {
		sliceSizes[i] = v / sum * fullCircle
	}

	borderCircle := &circle{center, radius}
	middleCircle := circle{Point: center, radius: radius / 2.0}

	// draw sectors
	phi := pc.Offset
	for i, size := range sliceSizes {
		for j := 0.0; j < size; j += resolutionFactor {
			borderPoint := borderCircle.at(phi + j)
			line := line{P1: center, P2: borderPoint}
			line.draw(&Cell{solidBlock, AttrPair{SelectAttr(pc.Colors, i), ColorDefault}}, &buf)
		}
		phi += size
	}

	// draw labels
	if pc.Label != nil {
		phi = pc.Offset
		for i, size := range sliceSizes {
			labelPoint := middleCircle.at(phi + size/2.0)
			if len(pc.Data) == 1 {
				labelPoint = center
			}
			buf.SetString(
				pc.Label(i, pc.Data[i]),
				image.Pt(labelPoint.X, labelPoint.Y),
				AttrPair{SelectAttr(pc.Colors, i), ColorDefault},
			)
			phi += size
		}
	}

	return buf
}

type circle struct {
	image.Point
	radius float64
}

// computes the point at a given angle phi
func (c circle) at(phi float64) image.Point {
	x := c.X + int(round(xStretch*c.radius*math.Cos(phi)))
	y := c.Y + int(round(c.radius*math.Sin(phi)))
	return image.Point{X: x, Y: y}
}

// computes the perimeter of a circle
func (c circle) perimeter() float64 {
	return 2.0 * math.Pi * c.radius
}

// a line between two points
type line struct {
	P1, P2 image.Point
}

// draws the line
func (l line) draw(cell *Cell, buf *Buffer) {
	isLeftOf := func(p1, p2 image.Point) bool {
		return p1.X <= p2.X
	}
	isTopOf := func(p1, p2 image.Point) bool {
		return p1.Y <= p2.Y
	}
	p1, p2 := l.P1, l.P2
	buf.SetCell(Cell{'*', cell.Attributes}, l.P2)
	width, height := l.size()
	if width > height { // paint left to right
		if !isLeftOf(p1, p2) {
			p1, p2 = p2, p1
		}
		flip := 1.0
		if !isTopOf(p1, p2) {
			flip = -1.0
		}
		for x := p1.X; x <= p2.X; x++ {
			ratio := float64(height) / float64(width)
			factor := float64(x - p1.X)
			y := ratio * factor * flip
			buf.SetCell(*cell, image.Pt(x, int(round(y))+p1.Y))
		}
	} else { // paint top to bottom
		if !isTopOf(p1, p2) {
			p1, p2 = p2, p1
		}
		flip := 1.0
		if !isLeftOf(p1, p2) {
			flip = -1.0
		}
		for y := p1.Y; y <= p2.Y; y++ {
			ratio := float64(width) / float64(height)
			factor := float64(y - p1.Y)
			x := ratio * factor * flip
			buf.SetCell(*cell, image.Pt(int(round(x))+p1.X, y))
		}
	}
}

// width and height of a line
func (l line) size() (w, h int) {
	return absInt(l.P2.X - l.P1.X), absInt(l.P2.Y - l.P1.Y)
}

// rounds a value
func round(x float64) float64 {
	return math.Floor(x + 0.5)
}

// fold a sum
func sumFloat64(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum
}

// math.Abs for ints
func absInt(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func minFloat64(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

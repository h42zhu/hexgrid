package util

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

// Grid is a 2 by 2 matrix
type Grid struct {
	Position pixel.Vec
	Width    int
	Height   int
	CellSize float64
	Matrix   [][]string
}

// Vec2 is a container struct for (x, y)
type Vec2 struct {
	X int
	Y int
}

// GridRender wraps grid with drawing refs
type GridRender struct {
	Grid      *Grid
	win       *pixelgl.Window
	imd       *imdraw.IMDraw
	textAtlas *text.Atlas
}

// NewGrid returns a Grid
func NewGrid(width int, height int, cellSize float64, position pixel.Vec) *Grid {
	grid := &Grid{
		Position: position,
		Width:    width,
		Height:   height,
		CellSize: cellSize,
	}

	m := make([][]string, height)
	for i := range m {
		m[i] = make([]string, width)
		for j := range m[i] {
			m[i][j] = fmt.Sprintf("%d %d", j+1, i+1)
		}
	}
	grid.Matrix = m

	return grid
}

// GetWorldPosition converts a grid cell to position on screen
func (g *Grid) GetWorldPosition(x int, y int) pixel.Vec {
	return pixel.V(float64(x), float64(y)).Scaled(g.CellSize).Add(g.Position)
}

// GetXY returns x, y index from world coordinate
func (g *Grid) GetXY(v pixel.Vec) Vec2 {
	x := int((v.X - g.Position.X) / g.CellSize)
	y := int((v.Y - g.Position.Y) / g.CellSize)
	return Vec2{X: x, Y: y}
}

// NewGridRender ...
func NewGridRender(width int, height int, cellSize float64, position pixel.Vec, win *pixelgl.Window, imd *imdraw.IMDraw, textAtlas *text.Atlas) *GridRender {
	grid := NewGrid(width, height, cellSize, position)
	return &GridRender{
		Grid:      grid,
		win:       win,
		imd:       imd,
		textAtlas: textAtlas,
	}
}

// ShowGrid draws grid on screen
func (g *GridRender) ShowGrid(color color.RGBA) {

	drawRect := func(x int, y int) {
		v := g.Grid.GetWorldPosition(x, y)
		g.imd.Color = color
		g.imd.Push(v)
		g.imd.Push(v.Add(pixel.V(g.Grid.CellSize, 0)))
		g.imd.Push(v.Add(pixel.V(g.Grid.CellSize, g.Grid.CellSize)))
		g.imd.Push(v.Add(pixel.V(0, g.Grid.CellSize)))
		g.imd.Rectangle(1)
	}

	for i := 0; i < g.Grid.Width; i++ {
		for j := 0; j < g.Grid.Height; j++ {
			drawRect(i, j)
		}
	}
}

// ShowText shows the text in the grid
func (g *GridRender) ShowText(color color.RGBA, scale float64) {

	for i := 0; i < g.Grid.Width; i++ {
		for j := 0; j < g.Grid.Height; j++ {
			v := g.Grid.GetWorldPosition(i, j)
			basicTxt := text.New(v, g.textAtlas)
			basicTxt.Color = color
			fmt.Fprintln(basicTxt, g.Grid.Matrix[i][j])
			basicTxt.Draw(g.win, pixel.IM.Scaled(basicTxt.Orig, scale))
		}
	}
}

// SetValue sets a value in the grid
func (g *GridRender) SetValue(x int, y int, value string) {
	if x >= g.Grid.Width || y >= g.Grid.Height || x < 0 || y < 0 {
		return
	}
	g.Grid.Matrix[x][y] = value
}

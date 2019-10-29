package hexagon

import (
	"image/color"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

// Renderer is a utility type for rendering text and hex shapes
type Renderer struct {
	win *pixelgl.Window
	imd *imdraw.IMDraw
}

// NewRenderer creates a new render instance
func NewRenderer(win *pixelgl.Window, imd *imdraw.IMDraw) *Renderer {
	return &Renderer{
		win,
		imd,
	}
}

// DrawHex draws a single hexagon
func (r *Renderer) DrawHex(hexCell HexCell, border float64, color color.RGBA) {
	r.imd.Color = color

	corners := hexCell.GetAllCorners()
	// connect the original corner for drawing polygon
	corners = append(corners, corners[0])
	r.imd.Push(corners...)
	r.imd.Polygon(border)
}

// DrawHexGrid draws a hex grid
func (r *Renderer) DrawHexGrid(hg HexGrid, border float64, color color.RGBA) {
	for _, hex := range hg.Cells {
		r.DrawHex(hex, border, color)
	}
}

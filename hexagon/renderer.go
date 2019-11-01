package hexagon

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

// Renderer is a utility type for rendering text and hex shapes
type Renderer struct {
	win       *pixelgl.Window
	imd       *imdraw.IMDraw
	textAtlas *text.Atlas
}

// NewRenderer creates a new render instance
func NewRenderer(win *pixelgl.Window, imd *imdraw.IMDraw, textAtlas *text.Atlas) *Renderer {
	return &Renderer{
		win,
		imd,
		textAtlas,
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

// drawText paints the index of the hexCell
func (r *Renderer) drawText(hexCelll HexCell, scale float64, color color.RGBA) {

	basicTxt := text.New(hexCelll.Center.Add(pixel.V(-4*hexCelll.Radius/5, 0)), r.textAtlas)
	basicTxt.Color = color
	fmt.Fprintln(basicTxt, hexCelll.Index.String())
	basicTxt.Draw(r.win, pixel.IM.Scaled(basicTxt.Orig, scale))
}

// DrawTextGrid ...
func (r *Renderer) DrawTextGrid(hg *HexGrid, scale float64, color color.RGBA) {
	for _, hex := range hg.Cells {
		r.drawText(hex, scale, color)
	}
}

// DrawHexGrid draws a hex grid
func (r *Renderer) DrawHexGrid(hg *HexGrid, border float64, color color.RGBA) {

	for _, hex := range hg.Cells {
		r.DrawHex(hex, border, color)
	}
}

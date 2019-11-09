package scene

import (
	"fmt"
	"image/color"
	"tactics/hexagon"

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
func (r *Renderer) DrawHex(hexCell *hexagon.HexCell, border float64, color color.RGBA) {
	r.imd.Color = color

	corners := hexCell.GetAllCorners()
	// connect the original corner for drawing polygon
	corners = append(corners, corners[0])
	r.imd.Push(corners...)
	r.imd.Polygon(border)
}

// drawHexIndex paints the index of the hexCell
func (r *Renderer) drawHexIndex(hexCelll *hexagon.HexCell, scale float64, color color.RGBA) {

	basicTxt := text.New(hexCelll.Center.Add(pixel.V(-4*hexCelll.Radius/5, 0)), r.textAtlas)
	basicTxt.Color = color
	basicTxt.Draw(r.win, pixel.IM.Scaled(basicTxt.Orig, scale))
}

// ShowHexGridIndex ...
func (r *Renderer) ShowHexGridIndex(hg *hexagon.HexGrid, scale float64, color color.RGBA) {
	for _, hex := range hg.Cells {
		r.drawHexIndex(hex, scale, color)
	}
}

// DrawHexGrid draws a hex grid
func (r *Renderer) DrawHexGrid(hg *hexagon.HexGrid, border float64, color color.RGBA) {
	for _, hex := range hg.Cells {
		r.DrawHex(hex, border, color)
	}
}

// DrawHoverCell draws the selected cell
func (r *Renderer) DrawHoverCell(hg *hexagon.HexGrid, border float64, color color.RGBA) {
	if hg.HoverCell != nil {
		if hexCell, ok := hg.Cells[*hg.HoverCell]; ok {
			r.DrawHex(hexCell, border, color)
		}
	}
}

// RenderEntityPosition draws an entity sprite onto the scene
func (r *Renderer) RenderEntityPosition(entity *hexagon.Entity, position pixel.Matrix) {
	entity.Sprite.Draw(r.win, position)
}

// RenderEntityHex ..
func (r *Renderer) RenderEntityHex(entity *hexagon.Entity, hg *hexagon.HexGrid) {
	v, err := hg.GetWorldPosition(entity.GetIndex())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	position := pixel.IM.Moved(v)
	r.RenderEntityPosition(entity, position)
}

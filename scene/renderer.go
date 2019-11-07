package scene

import (
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

// DrawSelectedCell draws the selected cell
func (r *Renderer) DrawSelectedCell(hg *hexagon.HexGrid, border float64, color color.RGBA) {
	if hg.SelectedCell != nil {
		if hexCell, ok := hg.Cells[*hg.SelectedCell]; ok {
			r.DrawHex(hexCell, border, color)
		}
	}
}

// RenderEntity draws an entity sprite onto the scene
func (r *Renderer) RenderEntity(entity *Entity, position pixel.Matrix) {
	entity.Sprite.Draw(r.win, position)
}

// RenderEntityOnHexGrid ..
func (r *Renderer) RenderEntityOnHexGrid(entity *Entity, hg *hexagon.HexGrid) {
	v := hg.GetWorldPosition(entity.Index)
	position := pixel.IM.Moved(v)
	r.RenderEntity(entity, position)
}
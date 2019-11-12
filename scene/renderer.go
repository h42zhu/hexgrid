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

// RenderContext is a utility type for rendering text and hex shapes
type RenderContext struct {
	win       *pixelgl.Window
	imd       *imdraw.IMDraw
	textAtlas *text.Atlas
}

// NewRenderContext creates a new render instance
func NewRenderContext(win *pixelgl.Window, imd *imdraw.IMDraw, textAtlas *text.Atlas) *RenderContext {
	return &RenderContext{
		win,
		imd,
		textAtlas,
	}
}

// DrawHex draws a single hexagon
func DrawHex(r *RenderContext, hexCell *hexagon.HexCell, border float64, color color.RGBA) {
	r.imd.Color = color

	corners := hexCell.GetAllCorners()
	// connect the original corner for drawing polygon
	corners = append(corners, corners[0])
	r.imd.Push(corners...)
	r.imd.Polygon(border)
}

// drawHexIndex paints the index of the hexCell
func drawHexIndex(r *RenderContext, hexCelll *hexagon.HexCell, scale float64, color color.RGBA) {

	basicTxt := text.New(hexCelll.Center.Add(pixel.V(-4*hexCelll.Radius/5, 0)), r.textAtlas)
	basicTxt.Color = color
	basicTxt.Draw(r.win, pixel.IM.Scaled(basicTxt.Orig, scale))
}

// ShowHexGridIndex ...
func ShowHexGridIndex(r *RenderContext, hg *hexagon.HexGrid, scale float64, color color.RGBA) {
	for _, hex := range hg.Cells {
		drawHexIndex(r, hex, scale, color)
	}
}

// DrawHexGrid draws a hex grid
func DrawHexGrid(r *RenderContext, hg *hexagon.HexGrid, border float64, color color.RGBA) {
	for _, hex := range hg.Cells {
		DrawHex(r, hex, border, color)
	}
}

// DrawHoverCell draws the selected cell
func DrawHoverCell(r *RenderContext, hg *hexagon.HexGrid, border float64, color color.RGBA) {
	if hg.HoverCell != nil {
		if hexCell, ok := hg.Cells[*hg.HoverCell]; ok {
			DrawHex(r, hexCell, border, color)
		}
	}
}

// RenderEntityPosition draws an entity sprite onto the scene
func RenderEntityPosition(r *RenderContext, entity *hexagon.Entity, position pixel.Matrix) {
	entity.Sprite.Draw(r.win, position)
}

// RenderEntityHex ..
func RenderEntityHex(r *RenderContext, entity *hexagon.Entity, hg *hexagon.HexGrid) {
	v, err := hg.GetWorldPosition(entity.GetIndex())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	position := pixel.IM.Moved(v)
	RenderEntityPosition(r, entity, position)
}

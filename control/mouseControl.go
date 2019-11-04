package control

import (
	"tactics/hexagon"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// MouseControl is a
type MouseControl struct {
	Win *pixelgl.Window
}

// GetMousePosition returns the mouse position
func (m *MouseControl) GetMousePosition() pixel.Vec {
	return m.Win.MousePosition()
}

// SelectCell sets the SelectedCell in hexGrid
func (m *MouseControl) SelectCell(grid *hexagon.HexGrid) bool {
	p := m.GetMousePosition()
	idx := grid.GetIndex(p)
	grid.SelectedCell = &idx

	return true
}

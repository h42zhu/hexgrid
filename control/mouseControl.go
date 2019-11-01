package control

import (
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

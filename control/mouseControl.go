package util

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// MouseControl is a
type MouseControl struct {
	Win        *pixelgl.Window
	GridRender *GridRender
}

// Update checks
func (m *MouseControl) Update(dt float64) {
	win := m.Win
	if win.JustPressed(pixelgl.MouseButtonLeft) {

		mouse := win.MousePosition()
		idx := m.GridRender.Grid.GetXY(mouse)

		m.GridRender.UpdateSelected(idx.X, idx.Y, 1, colornames.Yellow)
	}
}

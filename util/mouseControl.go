package util

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
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
		fmt.Println(idx)
		m.GridRender.SetValue(idx.X, idx.Y, "")
	}
}

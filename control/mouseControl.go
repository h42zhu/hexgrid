package control

import (
	"fmt"
	"tactics/scene"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// MouseControl is a struct
type MouseControl struct {
	Win *pixelgl.Window
}

// NewMouseControl ..
func NewMouseControl(win *pixelgl.Window) *MouseControl {
	return &MouseControl{
		Win: win,
	}
}

// MouseAction ..
func (m *MouseControl) MouseAction(scene *scene.Scene) {
	pos := m.Win.MousePosition()
	idx := scene.Grid.GetIndex(pos)
	selectedIdx := scene.Grid.SelectedCell

	if selectedIdx == nil {
		scene.Grid.SelectedCell = &idx
		return
	}

	if entity, ok := scene.Grid.Cells[*selectedIdx]; ok {
		fmt.Println(entity)
	}

}

// GetMousePosition returns the mouse position
func (m *MouseControl) GetMousePosition() pixel.Vec {
	return m.Win.MousePosition()
}

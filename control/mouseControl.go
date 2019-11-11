package control

import (
	"sync"
	"tactics/hexagon"
	"tactics/scene"

	"github.com/faiface/pixel/pixelgl"
)

// MouseControl is a struct
type MouseControl struct {
	Win  *pixelgl.Window
	Lock sync.RWMutex
}

// NewMouseControl ..
func NewMouseControl(win *pixelgl.Window) *MouseControl {
	return &MouseControl{
		Win: win,
	}
}

// MouseActionClickLeft ..
func (m *MouseControl) MouseActionClickLeft(scene *scene.Scene) {
	pos := m.Win.MousePosition()
	idx := scene.Grid.GetIndex(pos)

	se := scene.Grid.SelectedEntity

	if se == "" {
		// no entity selected yet
		for k, v := range scene.Ally {
			if v.GetIndex() == idx {
				scene.Grid.SelectedEntity = k
				return
			}
		}
	} else if entity, ok := scene.Ally[se]; ok {
		// cancel selection
		if entity.GetIndex() != idx {
			scene.Grid.SelectedEntity = ""
		}
	} else {
		scene.Grid.SelectedEntity = ""
	}
}

// MouseActionClickRight ..
func (m *MouseControl) MouseActionClickRight(scene *scene.Scene) {
	se := scene.Grid.SelectedEntity
	if se != "" {
		// move entity to mouse position
		hc := scene.Grid.HoverCell
		if entity, ok := scene.Ally[se]; ok && hc != nil {
			entity.SetIndex(hexagon.NewHexIndex(hc.X, hc.Y))
		}

	}
}

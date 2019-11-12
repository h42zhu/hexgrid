package control

import (
	"sync"

	"github.com/faiface/pixel/pixelgl"
)

// MouseControl is a struct
type MouseControl struct {
	Win *pixelgl.Window
	mu  sync.RWMutex
}

// NewMouseControl ..
func NewMouseControl(win *pixelgl.Window) *MouseControl {
	return &MouseControl{
		Win: win,
	}
}

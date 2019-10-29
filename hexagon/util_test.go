package hexagon

import (
	"testing"

	"github.com/faiface/pixel"
)

func TestGenRing(t *testing.T) {
	res := genRing(2)
	if len(res) != 19 {
		t.Errorf("expected length to be %d, got %d", 19, len(res))
	}
}

func TestNewHexGrid(t *testing.T) {
	hg := NewHexGrid(5, pixel.V(0, 0), 2)
	// fmt.Println(hg)
	if hg.Center != pixel.V(0, 0) {
		t.Errorf("expected center to be %+v, got %+v", pixel.V(0, 0), hg.Center)
	}
}

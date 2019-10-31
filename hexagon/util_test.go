package hexagon

import (
	"testing"

	"github.com/faiface/pixel"
)

func TestGenIndex(t *testing.T) {

}

func TestNewHexGrid(t *testing.T) {
	hg := NewHexGrid(5, pixel.V(0, 0), 2)
	// fmt.Println(hg)
	if hg.Center != pixel.V(0, 0) {
		t.Errorf("expected center to be %+v, got %+v", pixel.V(0, 0), hg.Center)
	}
}

package util

import (
	"github.com/faiface/pixel"
)

// Grid is a 2 by 2 matrix
type Grid struct {
	width    int
	height   int
	cellSize float64
	matrix   [][]string
}

// GetWorldPosition converts a grid cell to position on screen
func (g *Grid) GetWorldPosition(x int, y int) pixel.Vec {
	return pixel.V(float64(x), float64(y)).Scaled(g.cellSize)
}

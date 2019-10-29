package hexagon

import (
	"math"

	"github.com/faiface/pixel"
)

// HexCell defines a hexagonal cell
type HexCell struct {
	Center pixel.Vec
	Radius float64
	Index  Vector3
}

// GetCorner returns x, y coord at corner i (0..5)
// i == 0 points to the angle at -30 degrees
func (hc HexCell) GetCorner(i int) pixel.Vec {
	angleDeg := 60*i - 30
	angleRad := degToRad(float64(angleDeg))
	return pixel.V(hc.Center.X+hc.Radius*math.Cos(angleRad), hc.Center.X+hc.Radius*math.Sin(angleRad))
}

// GetAllCorners returns a slice of six corners
func (hc HexCell) GetAllCorners() []pixel.Vec {
	s := make([]pixel.Vec, 6)
	for i := 0; i < 6; i++ {
		s[i] = hc.GetCorner(i)
	}
	return s
}

// IdxDistance ..
func (hc HexCell) IdxDistance(dest HexCell) int {
	h1 := hc.Index
	h2 := dest.Index
	return int(math.Abs(float64(h1.X-h2.X))+math.Abs(float64(h1.Y-h2.Y))+math.Abs(float64(h1.Z-h2.Z))) / 2
}

// HexGrid is a collection of HexCells
type HexGrid struct {
	CellSize float64
	Cells    map[Vector3]HexCell
	Center   pixel.Vec
}

// NewHexGrid creates a new HexGrid
func NewHexGrid(cellSize float64, center pixel.Vec, ringNum int) HexGrid {
	cells := make(map[Vector3]HexCell)
	idxList := genRing(ringNum)

	for _, idx := range idxList {
		cells[idx] = makeCellFromIdx(cellSize, idx, center)
	}

	return HexGrid{
		CellSize: cellSize,
		Center:   center,
		Cells:    cells,
	}

}

// NeighborIdx returns all the index of the neighbor cells
func (hg HexGrid) NeighborIdx(hc HexCell) []Vector3 {
	idx := hc.Index
	c := make([]Vector3, 6)
	for i, offset := range HexDirections {
		idxNeighbor := idx.Add(offset)
		c[i] = idxNeighbor
	}
	return c
}

// GetWorldPosition converts a cell idx to position on screen
func (hg HexGrid) GetWorldPosition(idx Vector3) pixel.Vec {
	return hg.Cells[idx].Center
}

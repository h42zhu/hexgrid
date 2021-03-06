package hexagon

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
)

type HexIndex struct {
	X int
	Y int
}

// String returns a string repr of hex index
func (h HexIndex) String() string {
	return fmt.Sprintf("(%d, %d)", h.X, h.Y)
}

// NewHexIndex ..
func NewHexIndex(x int, y int) HexIndex {
	return HexIndex{
		X: x,
		Y: y,
	}
}

// HexCell defines a hexagonal cell
type HexCell struct {
	Center pixel.Vec
	Radius float64
	Index  HexIndex
}

// GetCorner returns x, y coord at corner i (0..5)
// i == 0 points to the angle at -30 degrees
func (hc *HexCell) GetCorner(i int) pixel.Vec {
	angleDeg := 60*i - 30
	angleRad := degToRad(float64(angleDeg))
	point := pixel.V(hc.Radius*math.Cos(angleRad), hc.Radius*math.Sin(angleRad))

	return hc.Center.Add(point)
}

// GetAllCorners returns a slice of six corners
func (hc *HexCell) GetAllCorners() []pixel.Vec {
	s := make([]pixel.Vec, 6)
	for i := 0; i < 6; i++ {
		s[i] = hc.GetCorner(i)
	}
	return s
}

// IdxDistance ..
func (hc *HexCell) IdxDistance(dest HexCell) int {
	h1 := doubleWidthVec3(hc.Index)
	h2 := doubleWidthVec3(dest.Index)

	return int(math.Abs(float64(h1.X-h2.X))+math.Abs(float64(h1.Y-h2.Y))+math.Abs(float64(h1.Z-h2.Z))) / 2
}

// HexGrid is a collection of HexCells
type HexGrid struct {
	CellSize       float64
	Cells          map[HexIndex]*HexCell
	Center         pixel.Vec
	SelectedEntity string
	HoverCell      *HexIndex
}

// NewHexGrid creates a new HexGrid
func NewHexGrid(cellSize float64, center pixel.Vec, sizeX int, sizeY int) *HexGrid {
	cells := make(map[HexIndex]*HexCell)
	idxList := genIndex(sizeX, sizeY)

	for _, idx := range idxList {
		cells[idx] = makeCellFromIdx(cellSize, idx, center)
	}

	return &HexGrid{
		CellSize:       cellSize,
		Center:         center,
		Cells:          cells,
		SelectedEntity: "",
		HoverCell:      nil,
	}
}

// NeighborIdx returns all the index of the neighbor cells
func (hg *HexGrid) NeighborIdx(hc *HexCell) []pixel.Vec {
	s := []pixel.Vec{}
	for _, offset := range HexDirections {
		s = append(s, hg.Center.Add(offset))
	}

	return s
}

// GetWorldPosition converts a cell idx to position on screen
func (hg *HexGrid) GetWorldPosition(idx HexIndex) (pixel.Vec, error) {
	if hex, ok := hg.Cells[idx]; ok {
		return hex.Center, nil
	}

	return pixel.V(0, 0), fmt.Errorf("unable to find hex index: %s", idx.String())
}

// GetIndex returns the index of a hex cell from world position
func (hg *HexGrid) GetIndex(pos pixel.Vec) HexIndex {
	posAdj := pos.Add(hg.Center.Scaled(-1))

	x := math.Round(posAdj.Dot(InverseDoubleWidthBasisMatrix[0]) / hg.CellSize)
	y := math.Round(posAdj.Dot(InverseDoubleWidthBasisMatrix[1]) / hg.CellSize)

	intX := int(x)
	intY := int(y)

	return NewHexIndex(intX, intY)
}

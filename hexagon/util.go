package hexagon

import (
	"math"

	"github.com/faiface/pixel"
)

var (
	// HexDirections contains the index changes as moving one space in hex coordinates involves changing
	// one of the 3 cube coordinates by +1 and changing another one by -1
	HexDirections = []pixel.Vec{
		pixel.V(2, 0), pixel.V(1, -1), pixel.V(-1, -1),
		pixel.V(-2, 0), pixel.V(-1, 1), pixel.V(1, 1),
	}

	// DoubleWidthBasisMatrix is 2 by 2 matrix of the basis vectors of DD coordinates
	DoubleWidthBasisMatrix = []pixel.Vec{
		pixel.V(math.Sqrt(3)/2, 0),
		pixel.V(0, 3.0/2),
	}

	// InverseDoubleWidthBasisMatrix is 2 by 2 matrix of the basis vectors of DD coordinates
	InverseDoubleWidthBasisMatrix = []pixel.Vec{
		pixel.V(2.0/math.Sqrt(3), 0),
		pixel.V(0, 2.0/3),
	}
)

// Vector3 is a container struct for indices of the form (x, y, z)
type Vector3 struct {
	X int
	Y int
	Z int
}

// Add adds two vec3
func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vec3(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

// Vec3 returns a new Vector3
func Vec3(x int, y int, z int) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func doubleWidthVec3(center HexIndex) Vector3 {
	x := int((center.Y - center.X) / 2)
	z := int(center.X)
	y := int(-x - z)
	return Vec3(x, y, z)
}

func degToRad(deg float64) float64 {
	return math.Pi / 180 * deg
}

func genIndex(sizeX int, sizeY int) []HexIndex {
	s := []HexIndex{}

	for y := 0; y <= sizeY*2; y += 2 {
		for x := 0; x <= sizeX*2; x += 2 {
			s = append(s, NewHexIndex(x, y))
		}

		for x := 1; x <= sizeX*2+1; x += 2 {
			s = append(s, NewHexIndex(x, y+1))
		}

	}
	return s
}

// makeCellFromIdx creates a hex cell from an idx
func makeCellFromIdx(radius float64, idx HexIndex, offset pixel.Vec) *HexCell {

	idxVec := pixel.V(float64(idx.X), float64(idx.Y))
	x := idxVec.Dot(DoubleWidthBasisMatrix[0]) * radius
	y := idxVec.Dot(DoubleWidthBasisMatrix[1]) * radius

	return &HexCell{
		Center: pixel.V(x, y).Add(offset),
		Radius: radius,
		Index:  idx,
	}
}

// CheckValidIndex makes sure the idx is of the double format
func CheckValidIndex(idx HexIndex) bool {
	return (idx.X+idx.Y)%2 == 0
}

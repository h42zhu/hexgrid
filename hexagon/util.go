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
		pixel.V(3/2, 0),
		pixel.V(0, math.Sqrt(3)),
	}

	// InverseDoubleWidthBasisMatrix is 2 by 2 matrix of the basis vectors of DD coordinates
	InverseDoubleWidthBasisMatrix = []pixel.Vec{
		pixel.V(2/3, 0),
		pixel.V(0, math.Sqrt(3)/3),
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

func doubleWidthVec3(center pixel.Vec) Vector3 {
	x := int((center.Y - center.X) / 2)
	z := int(center.X)
	y := int(-x - z)
	return Vec3(x, y, z)
}

func degToRad(deg float64) float64 {
	return math.Pi / 180 * deg
}

func genIndex(size int) []pixel.Vec {
	s := []pixel.Vec{}

	for x := 0; x <= size; x++ {
		m := x * 2
		for y := 0; y <= m; y++ {
			z := m - y
			s = append(s, pixel.V(float64(y), float64(z)))
		}

	}
	return s
}

// makeCellFromIdx creates a hex cell from an idx
func makeCellFromIdx(radius float64, idx pixel.Vec, offset pixel.Vec) HexCell {
	x := idx.Dot(DoubleWidthBasisMatrix[0]) * radius
	y := idx.Dot(DoubleWidthBasisMatrix[1]) * radius

	return HexCell{
		Center: pixel.V(x, y).Add(offset),
		Radius: radius,
		Index:  idx,
	}
}

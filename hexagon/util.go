package hexagon

import (
	"math"

	"github.com/faiface/pixel"
)

var (
	// HexDirections contains the index changes as moving one space in hex coordinates involves changing
	// one of the 3 cube coordinates by +1 and changing another one by -1
	HexDirections = []Vector3{
		Vec3(1, -1, 0), Vec3(1, 0, -1), Vec3(0, 1, -1),
		Vec3(-1, 1, 0), Vec3(-1, 0, 1), Vec3(0, -1, 1),
	}
	// AxialBasisMatrix is 2 by 2 matrix of the basis vectors of Axial coordinates
	AxialBasisMatrix = []pixel.Vec{
		pixel.V(math.Sqrt(3), math.Sqrt(3)/2),
		pixel.V(0, 3/2),
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

func degToRad(deg float64) float64 {
	return math.Pi / 180 * deg
}

func genRing(size int) []Vector3 {
	s := []Vector3{}
	zeroVec := Vec3(0, 0, 0)

	if size == 0 {
		s = append(s, zeroVec)
		return s
	}

	for x := -size; x <= size; x++ {
		for y := -size; y <= size; y++ {
			z := -x - y
			v := Vec3(x, y, z)
			if z <= size && z >= -size {
				s = append(s, v)
			}
		}
	}
	return s
}

// makeCellFromIdx creates a hex cell from an idx
func makeCellFromIdx(radius float64, idx Vector3, offset pixel.Vec) HexCell {
	vec := pixel.V(float64(idx.X), float64(idx.Z))
	x := vec.Dot(AxialBasisMatrix[0]) * radius
	y := vec.Dot(AxialBasisMatrix[1]) * radius

	return HexCell{
		Center: pixel.V(x, y).Add(offset),
		Radius: radius,
		Index:  idx,
	}
}

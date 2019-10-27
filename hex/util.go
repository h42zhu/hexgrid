package hexagon

import (
	"math"
)

var (
	// HexDirections contains the index changes as moving one space in hex coordinates involves changing
	// one of the 3 cube coordinates by +1 and changing another one by -1
	HexDirections = []Vector3{
		Vec3(1, -1, 0), Vec3(1, 0, -1), Vec3(0, 1, -1),
		Vec3(-1, 1, 0), Vec3(-1, 0, 1), Vec3(0, -1, 1),
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

func genIdxRing(idx int) []Vector3 {
	s := []Vector3{}
	s = append(s, Vec3(0, 0, 0))

	if idx == 0 {
		return s
	}

	for i := 0; i <= idx; i++ {

	}
	return s
}

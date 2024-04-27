// Package wire implements wireframe 3d shapes.
package wire

type worldDef struct {
	FL          float64
	CX, CY, CZ  float64
	NearZ, FarZ float64
}

// World contains the parameters for the 3d world.
var World = worldDef{
	FL:    300.0,
	CX:    0.0,
	CY:    0.0,
	CZ:    0.0,
	NearZ: 100.0,
	FarZ:  100000.0,
}

// Visible returns whether or not a point should be visible.
func Visible(point *Point) bool {
	if point.Z+World.CZ < World.NearZ {
		return false
	}
	if point.Z+World.CZ > World.FarZ {
		return false
	}
	return true
}

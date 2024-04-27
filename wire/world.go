// Package wire implements wireframe 3d shapes.
package wire

type worldDef struct {
	FL         float64
	CX, CY, CZ float64
}

// World contains the parameters for the 3d world.
var World = worldDef{
	FL: 300.0,
	CX: 0.0,
	CY: 0.0,
	CZ: 0.0,
}

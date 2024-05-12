// Package wire implements wireframe 3d shapes.
package wire

import (
	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/geom"
)

// Context interface to allow for drawing functions.
// The interface defines only the used methods of cairo.Context.
// This is needed to avoid recursive dependencies between wire and cairo.
type Context interface {
	StrokePath(geom.PointList, bool)
	FillCircle(float64, float64, float64)
	MoveTo(float64, float64)
	LineTo(float64, float64)
	Stroke()
	ClosePath()
	SetLineWidth(float64)
	GetLineWidth() float64
	Save()
	Restore()
	SetSourceColor(blcolor.Color)
	GetSourceRGB() (float64, float64, float64)
}

type worldDef struct {
	FL          float64
	CX, CY, CZ  float64
	NearZ, FarZ float64
	Fog         bool
	NearFog     float64
	FarFog      float64
	R, G, B     float64
	Context     Context
}

// World contains the parameters for the 3d world.
var World = worldDef{
	FL:      300.0,
	CX:      0.0,
	CY:      0.0,
	CZ:      0.0,
	NearZ:   100.0,
	FarZ:    100000.0,
	Fog:     false,
	NearFog: 400.0,
	FarFog:  1200.0,
	R:       1,
	G:       1,
	B:       1,
	Context: nil,
}

// InitWorld initializes the world.
func InitWorld(context Context, cx, cy, cz float64) {
	World.Context = context
	SetRGB(context.GetSourceRGB())
	World.CX = cx
	World.CY = cy
	World.CZ = cz
}

// SetRGB sets the drawing color.
func SetRGB(r, g, b float64) {
	World.R = r
	World.G = g
	World.B = b
}

// FogAmount returns the amount of fog to apply for the given object z.
func FogAmount(objectZ float64) float64 {
	if !World.Fog {
		return 1.0
	}
	f := blmath.Map(objectZ+World.CZ, World.NearFog, World.FarFog, 1, 0)
	return blmath.Clamp(f, 0, 1)
}

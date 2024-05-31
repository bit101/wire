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
	FogActive   bool
	NearFog     float64
	FarFog      float64
	R, G, B     float64
	Context     Context
	Font        FontType
	FontSize    float64
	FontSpacing float64
}

// World contains the parameters for the 3d world.
var world = worldDef{
	FL:          300.0,
	CX:          0.0,
	CY:          0.0,
	CZ:          0.0,
	NearZ:       100.0,
	FarZ:        100000.0,
	FogActive:   false,
	NearFog:     400.0,
	FarFog:      1200.0,
	R:           1,
	G:           1,
	B:           1,
	Context:     nil,
	Font:        FontAsteroid,
	FontSize:    100,
	FontSpacing: 0.2,
}

// InitWorld initializes the world.
func InitWorld(context Context, cx, cy, cz float64) {
	world.Context = context
	SetRGB(context.GetSourceRGB())
	SetCenter(cx, cy, cz)
}

// GetRGB returns the current drawing color.
func GetRGB() (float64, float64, float64) {
	return world.R, world.G, world.B
}

// SetRGB sets the drawing color.
func SetRGB(r, g, b float64) {
	world.R = r
	world.G = g
	world.B = b
}

// SetPerspective sets the amount of perspective to apply.
func SetPerspective(fl float64) {
	world.FL = fl
}

// SetCenter sets the center of the 3d world.
func SetCenter(x, y, z float64) {
	world.CX, world.CY, world.CZ = x, y, z
}

// SetClipping sets the near and far limits of rendering.
func SetClipping(near, far float64) {
	world.NearZ = near
	world.FarZ = far
}

// ApplyFog returns the amount of fog to apply for the given object z.
func ApplyFog(objectZ float64) {
	fog := 1.0
	if world.FogActive {
		fog = blmath.Map(objectZ+world.CZ, world.NearFog, world.FarFog, 1, 0)
		fog = blmath.Clamp(fog, 0, 1)
	}
	color := blcolor.RGBA(world.R, world.G, world.B, fog)
	world.Context.SetSourceColor(color)
}

// SetFog sets the fog parameters, including turning on and off.
func SetFog(active bool, near, far float64) {
	world.FogActive = active
	world.NearFog = near
	world.FarFog = far
}

// SetFont sets the font type, size and spacing for future text objects.
// Size is the width of a single letter. Default 100.
// Spacing is the space between letters, as a percentage of letter width. Defaults to 0.2.
func SetFont(font FontType, size, spacing float64) {
	world.Font = font
	world.FontSize = size
	world.FontSpacing = spacing
}

// SetFontType sets which font type will be used for future text objects.
// Default is wire.FontAsteroid.
func SetFontType(font FontType) {
	world.Font = font
}

// SetFontSize sets the font size (width of one letter) used for future text objects.
// Default is 100.
func SetFontSize(size float64) {
	world.FontSize = size
}

// SetFontSpacing sets the spacing between letters, as a percentage of letter width.
// Default is 0.2.
func SetFontSpacing(spacing float64) {
	world.FontSpacing = spacing
}

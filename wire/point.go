// Package wire implements wireframe 3d shapes.
package wire

import (
	"math"

	"github.com/bit101/bitlib/geom"
	"github.com/bit101/bitlib/random"
)

// Point is a 3d point.
type Point struct {
	X, Y, Z float64
}

// NewPoint creates a new point.
func NewPoint(x, y, z float64) *Point {
	return &Point{x, y, z}
}

// Clone returns a copy of this point.
func (p *Point) Clone() *Point {
	return NewPoint(p.X, p.Y, p.Z)
}

// Project projects this 3d point to a 2d point.
func (p *Point) Project() (*geom.Point, float64) {
	scale := World.FL / (World.CZ + p.Z)
	return geom.NewPoint(World.CX+p.X*scale, World.CY+p.Y*scale), scale
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// TranslateX translates this point on the x-axis in place.
func (p *Point) TranslateX(tx float64) {
	p.X += tx
}

// TranslateY translates this point on the y-axis in place.
func (p *Point) TranslateY(ty float64) {
	p.Y += ty
}

// TranslateZ translates this point on the z-axis in place.
func (p *Point) TranslateZ(tz float64) {
	p.Z += tz
}

// Translate translates this point in place.
func (p *Point) Translate(tx, ty, tz float64) {
	p.X += tx
	p.Y += ty
	p.Z += tz
}

// RotateX rotates this point around the x-axis in place.
func (p *Point) RotateX(angle float64) {
	c := math.Cos(angle)
	s := math.Sin(angle)
	y := c*p.Y + s*p.Z
	z := c*p.Z - s*p.Y
	p.Y = y
	p.Z = z
}

// RotateY rotates this point around the y-axis in place.
func (p *Point) RotateY(angle float64) {
	c := math.Cos(angle)
	s := math.Sin(angle)
	x := c*p.X + s*p.Z
	z := c*p.Z - s*p.X
	p.X = x
	p.Z = z
}

// RotateZ rotates this point around the z-axis in place.
func (p *Point) RotateZ(angle float64) {
	c := math.Cos(angle)
	s := math.Sin(angle)
	y := c*p.Y + s*p.X
	x := c*p.X - s*p.Y
	p.Y = y
	p.X = x
}

// Rotate rotates this point in place.
func (p *Point) Rotate(rx, ry, rz float64) {
	p.RotateX(rx)
	p.RotateY(ry)
	p.RotateZ(rz)
}

// Scale scales this point in place.
func (p *Point) Scale(sx, sy, sz float64) {
	p.X *= sx
	p.Y *= sy
	p.Z *= sz
}

// UniScale scales this point in place.
func (p *Point) UniScale(scale float64) {
	p.X *= scale
	p.Y *= scale
	p.Z *= scale
}

// Randomize randomizes this point in place.
func (p *Point) Randomize(amount float64) {
	p.X += random.FloatRange(-amount, amount)
	p.Y += random.FloatRange(-amount, amount)
	p.Z += random.FloatRange(-amount, amount)
}

//////////////////////////////
// Transform and return new
//////////////////////////////

// TranslatedX returns a copy of this point, translated on the x-axis.
func (p *Point) TranslatedX(tx float64) *Point {
	p1 := p.Clone()
	p1.TranslateX(tx)
	return p1
}

// TranslatedY returns a copy of this point, translated on the y-axis.
func (p *Point) TranslatedY(ty float64) *Point {
	p1 := p.Clone()
	p1.TranslateY(ty)
	return p1
}

// TranslatedZ returns a copy of this point, translated on the z-axis.
func (p *Point) TranslatedZ(tz float64) *Point {
	p1 := p.Clone()
	p1.TranslateZ(tz)
	return p1
}

// Translated returns a copy of this point, translated.
func (p *Point) Translated(tx, ty, tz float64) *Point {
	p1 := p.Clone()
	p1.Translate(tx, ty, tz)
	return p1
}

// RotatedX returns a copy of this point, rotated on the x-axis.
func (p *Point) RotatedX(angle float64) *Point {
	p1 := p.Clone()
	p1.RotateX(angle)
	return p1
}

// RotatedY returns a copy of this point, rotated on the y-axis.
func (p *Point) RotatedY(angle float64) *Point {
	p1 := p.Clone()
	p1.RotateY(angle)
	return p1
}

// RotatedZ returns a copy of this point, rotated on the z-axis.
func (p *Point) RotatedZ(angle float64) *Point {
	p1 := p.Clone()
	p1.RotateZ(angle)
	return p1
}

// Rotated returns a copy of this point, rotated.
func (p *Point) Rotated(rx, ry, rz float64) *Point {
	p1 := p.Clone()
	p1.Rotate(rx, ry, rz)
	return p1
}

// Scaled returns a copy of this point, scaled.
func (p *Point) Scaled(sx, sy, sz float64) *Point {
	p1 := p.Clone()
	p1.Scale(sx, sy, sz)
	return p1
}

// UniScaled returns a copy of this point, scaled.
func (p *Point) UniScaled(scale float64) *Point {
	p1 := p.Clone()
	p1.UniScale(scale)
	return p1
}

// Randomized returns a copy of this point, randomized.
func (p *Point) Randomized(amount float64) *Point {
	p1 := p.Clone()
	p1.Randomize(amount)
	return p1
}

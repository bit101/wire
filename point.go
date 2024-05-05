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

// NewPoint creates a new 3d point.
func NewPoint(x, y, z float64) *Point {
	return &Point{x, y, z}
}

// RandomPointInBox creates a new 3d point within a 3d box of the given dimensions.
// The box is centered on the origin, so points will range from -w/2 to w/2, etc. on each dimension.
func RandomPointInBox(w, h, d float64) *Point {
	return NewPoint(
		random.FloatRange(-w/2, w/2),
		random.FloatRange(-h/2, h/2),
		random.FloatRange(-d/2, d/2),
	)
}

// RandomPointOnSphere creates a random 3d point ON a sphere of the given radius.
func RandomPointOnSphere(radius float64) *Point {
	x := random.FloatRange(-1, 1)
	y := random.FloatRange(-1, 1)
	z := random.FloatRange(-1, 1)
	norm := 1 / math.Sqrt(x*x+y*y+z*z)
	return NewPoint(x*norm*radius, y*norm*radius, z*norm*radius)
}

// RandomPointInSphere creates a random 3d point IN a sphere of the given radius.
func RandomPointInSphere(radius float64) *Point {
	x := random.FloatRange(-1, 1)
	y := random.FloatRange(-1, 1)
	z := random.FloatRange(-1, 1)
	norm := 1 / math.Sqrt(x*x+y*y+z*z)
	radius = math.Sqrt(random.Float()) * radius
	return NewPoint(x*norm*radius, y*norm*radius, z*norm*radius)
}

// RandomPointOnCylinder creates a random 3d point ON a cylinder of the given radius and height.
func RandomPointOnCylinder(height, radius float64) *Point {
	angle := random.Angle()
	x := math.Cos(angle) * radius
	y := random.FloatRange(-height/2, height/2)
	z := math.Sin(angle) * radius
	return NewPoint(x, y, z)
}

// RandomPointInCylinder creates a random 3d point IN a cylinder of the given radius and height.
func RandomPointInCylinder(height, radius float64) *Point {
	radius = math.Sqrt(random.Float()) * radius
	angle := random.Angle()
	x := math.Cos(angle) * radius
	y := random.FloatRange(-height/2, height/2)
	z := math.Sin(angle) * radius
	return NewPoint(x, y, z)
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

// Distance returns the distance from this point to another point.
func (p *Point) Distance(other *Point) float64 {
	dx := other.X - p.X
	dy := other.Y - p.Y
	dz := other.Z - p.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// TranslateX translates this point on the x-axis, in place.
func (p *Point) TranslateX(tx float64) {
	p.X += tx
}

// TranslateY translates this point on the y-axis, in place.
func (p *Point) TranslateY(ty float64) {
	p.Y += ty
}

// TranslateZ translates this point on the z-axis, in place.
func (p *Point) TranslateZ(tz float64) {
	p.Z += tz
}

// Translate translates this point on all three axes, in place.
func (p *Point) Translate(tx, ty, tz float64) {
	p.X += tx
	p.Y += ty
	p.Z += tz
}

// RotateX rotates this point around the x-axis, in place.
func (p *Point) RotateX(angle float64) {
	c := math.Cos(angle)
	s := math.Sin(angle)
	y := c*p.Y + s*p.Z
	z := c*p.Z - s*p.Y
	p.Y = y
	p.Z = z
}

// RotateY rotates this point around the y-axis, in place.
func (p *Point) RotateY(angle float64) {
	c := math.Cos(angle)
	s := math.Sin(angle)
	x := c*p.X + s*p.Z
	z := c*p.Z - s*p.X
	p.X = x
	p.Z = z
}

// RotateZ rotates this point around the z-axis, in place.
func (p *Point) RotateZ(angle float64) {
	c := math.Cos(angle)
	s := math.Sin(angle)
	y := c*p.Y + s*p.X
	x := c*p.X - s*p.Y
	p.Y = y
	p.X = x
}

// Rotate rotates this point around all axes, in place.
func (p *Point) Rotate(rx, ry, rz float64) {
	p.RotateX(rx)
	p.RotateY(ry)
	p.RotateZ(rz)
}

// ScaleX scales this point on the x-axis, in place.
func (p *Point) ScaleX(scale float64) {
	p.X *= scale
}

// ScaleY scales this point on the y-axis, in place.
func (p *Point) ScaleY(scale float64) {
	p.Y *= scale
}

// ScaleZ scales this point on the z-axis, in place.
func (p *Point) ScaleZ(scale float64) {
	p.Z *= scale
}

// Scale scales this point on all axes, in place.
func (p *Point) Scale(sx, sy, sz float64) {
	p.X *= sx
	p.Y *= sy
	p.Z *= sz
}

// UniScale scales this point by the same amount on each axis, in place.
func (p *Point) UniScale(scale float64) {
	p.X *= scale
	p.Y *= scale
	p.Z *= scale
}

// RandomizeX randomizes this point on the x-axis, in place.
func (p *Point) RandomizeX(amount float64) {
	p.X += random.FloatRange(-amount, amount)
}

// RandomizeY randomizes this point on the y-axis, in place.
func (p *Point) RandomizeY(amount float64) {
	p.Y += random.FloatRange(-amount, amount)
}

// RandomizeZ randomizes this point on the z-axis, in place.
func (p *Point) RandomizeZ(amount float64) {
	p.Z += random.FloatRange(-amount, amount)
}

// Randomize randomizes this point on all axes, in place.
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

// Translated returns a copy of this point, translated on all axes.
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

// Rotated returns a copy of this point, rotated around all axes.
func (p *Point) Rotated(rx, ry, rz float64) *Point {
	p1 := p.Clone()
	p1.Rotate(rx, ry, rz)
	return p1
}

// ScaledX returns a copy of this point, scaled on the x-axis.
func (p *Point) ScaledX(scale float64) *Point {
	p1 := p.Clone()
	p1.ScaleX(scale)
	return p1
}

// ScaledY returns a copy of this point, scaled on the y-axis.
func (p *Point) ScaledY(scale float64) *Point {
	p1 := p.Clone()
	p1.ScaleY(scale)
	return p1
}

// ScaledZ returns a copy of this point, scaled on the z-axis.
func (p *Point) ScaledZ(scale float64) *Point {
	p1 := p.Clone()
	p1.ScaleY(scale)
	return p1
}

// Scaled returns a copy of this point, scaled on all axes.
func (p *Point) Scaled(sx, sy, sz float64) *Point {
	p1 := p.Clone()
	p1.Scale(sx, sy, sz)
	return p1
}

// UniScaled returns a copy of this point, scaled by the same amount on each axis.
func (p *Point) UniScaled(scale float64) *Point {
	p1 := p.Clone()
	p1.UniScale(scale)
	return p1
}

// RandomizedX returns a copy of this point, randomized on the x-axis.
func (p *Point) RandomizedX(amount float64) *Point {
	p1 := p.Clone()
	p1.RandomizeX(amount)
	return p1
}

// RandomizedY returns a copy of this point, randomized on the y-axis.
func (p *Point) RandomizedY(amount float64) *Point {
	p1 := p.Clone()
	p1.RandomizeY(amount)
	return p1
}

// RandomizedZ returns a copy of this point, randomized on the z-axis.
func (p *Point) RandomizedZ(amount float64) *Point {
	p1 := p.Clone()
	p1.RandomizeZ(amount)
	return p1
}

// Randomized returns a copy of this point, randomized on all axes.
func (p *Point) Randomized(amount float64) *Point {
	p1 := p.Clone()
	p1.Randomize(amount)
	return p1
}

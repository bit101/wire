// Package wire implements wireframe 3d shapes.
package wire

import (
	"log"

	"github.com/bit101/bitlib/geom"
)

// Context interface to allow for drawing functions.
type Context interface {
	StrokePath(geom.PointList, bool)
	MoveTo(float64, float64)
	LineTo(float64, float64)
	Stroke()
	ClosePath()
	SetLineWidth(float64)
	GetLineWidth() float64
}

// Shape is a 3d shape.
type Shape struct {
	Paths  []PointList
	Closed bool
}

// NewShape creates a new shape.
func NewShape(numPaths int, closed bool) *Shape {
	return &Shape{
		make([]PointList, numPaths),
		closed,
	}
}

// Add adds a new path to the shape.
func (s *Shape) Add(path PointList) {
	s.Paths = append(s.Paths, path)
}

// AddPoint adds a point to the shape at the given index.
func (s *Shape) AddPoint(index int, point *Point) {
	if index >= len(s.Paths) || index < 0 {
		log.Fatal("list index must be from 0 to one less than the size of the list")
	}
	s.Paths[index].Add(point)
}

// AddXYZ adds a point to the shape at the given index.
func (s *Shape) AddXYZ(index int, x, y, z float64) {
	if index >= len(s.Paths) || index < 0 {
		log.Fatal("list index must be from 0 to one less than the size of the list")
	}
	s.Paths[index].AddXYZ(x, y, z)
}

// Clone returns a deep copy of this shape.
func (s *Shape) Clone() *Shape {
	clone := &Shape{
		[]PointList{},
		s.Closed,
	}
	for _, pointList := range s.Paths {
		clone.Paths = append(clone.Paths, pointList.Clone())
	}
	return clone
}

// Stroke strokes each path in a shape.
func (s *Shape) Stroke(context Context) {
	for _, path := range s.Paths {
		path.Stroke(context, s.Closed)
	}
}

// Subdivide puts a new point between each pair of points.
func (s *Shape) Subdivide(times int) {
	for i := 0; i < len(s.Paths); i++ {
		s.Paths[i].Subdivide(times)
	}
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// TranslateX translates this shape on the x-axis in place.
func (s *Shape) TranslateX(tx float64) {
	for _, path := range s.Paths {
		path.TranslateX(tx)
	}
}

// TranslateY translates this shape on the y-axis in place.
func (s *Shape) TranslateY(ty float64) {
	for _, path := range s.Paths {
		path.TranslateY(ty)
	}
}

// TranslateZ translates this shape on the z-axis in place.
func (s *Shape) TranslateZ(tz float64) {
	for _, path := range s.Paths {
		path.TranslateZ(tz)
	}
}

// Translate translates this shape in place.
func (s *Shape) Translate(tx, ty, tz float64) {
	for _, list := range s.Paths {
		list.Translate(tx, ty, tz)
	}
}

// RotateX rotates this shape around the x-axis in place.
func (s *Shape) RotateX(angle float64) {
	for _, list := range s.Paths {
		list.RotateX(angle)
	}
}

// RotateY rotates this shape around the y-axis in place.
func (s *Shape) RotateY(angle float64) {
	for _, list := range s.Paths {
		list.RotateY(angle)
	}
}

// RotateZ rotates this shape around the z-axis in place.
func (s *Shape) RotateZ(angle float64) {
	for _, list := range s.Paths {
		list.RotateZ(angle)
	}
}

// Rotate rotates this shape in place.
func (s *Shape) Rotate(rx, ry, rz float64) {
	for _, list := range s.Paths {
		list.Rotate(rx, ry, rz)
	}
}

// ScaleX scales this shape on the x-axis, in place.
func (s *Shape) ScaleX(scale float64) {
	for _, path := range s.Paths {
		path.ScaleX(scale)
	}
}

// ScaleY scales this shape on the y-axis, in place.
func (s *Shape) ScaleY(scale float64) {
	for _, path := range s.Paths {
		path.ScaleY(scale)
	}
}

// ScaleZ scales this shape on the z-axis, in place.
func (s *Shape) ScaleZ(scale float64) {
	for _, path := range s.Paths {
		path.ScaleZ(scale)
	}
}

// Scale scales this shape in place.
func (s *Shape) Scale(sx, sy, sz float64) {
	for _, list := range s.Paths {
		list.Scale(sx, sy, sz)
	}
}

// UniScale scales this shape in place.
func (s *Shape) UniScale(scale float64) {
	for _, list := range s.Paths {
		list.UniScale(scale)
	}
}

// RandomizeX randomizes this shape on the x-axis, in place.
func (s *Shape) RandomizeX(amount float64) {
	for _, list := range s.Paths {
		list.RandomizeX(amount)
	}
}

// RandomizeY randomizes this shape on the y-axis, in place.
func (s *Shape) RandomizeY(amount float64) {
	for _, list := range s.Paths {
		list.RandomizeY(amount)
	}
}

// RandomizeZ randmizes this shape on the z-axis, in place.
func (s *Shape) RandomizeZ(amount float64) {
	for _, list := range s.Paths {
		list.RandomizeZ(amount)
	}
}

// Randomize randomizes this shape in place.
func (s *Shape) Randomize(amount float64) {
	for _, list := range s.Paths {
		list.Randomize(amount)
	}
}

//////////////////////////////
// Transform and return new
//////////////////////////////

// TranslatedX returns a copy of this shape, translated on the x-axis.
func (s *Shape) TranslatedX(tx float64) *Shape {
	s1 := s.Clone()
	s1.TranslateX(tx)
	return s1
}

// TranslatedY returns a copy of this shape, translated on the y-axis.
func (s *Shape) TranslatedY(ty float64) *Shape {
	s1 := s.Clone()
	s1.TranslateY(ty)
	return s1
}

// TranslatedZ returns a copy of this shape, translated on the z-axis.
func (s *Shape) TranslatedZ(tz float64) *Shape {
	s1 := s.Clone()
	s1.TranslateZ(tz)
	return s1
}

// Translated returns a copy of this shape, translated.
func (s *Shape) Translated(tx, ty, tz float64) *Shape {
	s1 := s.Clone()
	s1.Translate(tx, ty, tz)
	return s1
}

// RotatedX returns a copy of this shape, rotated on the x-axis.
func (s *Shape) RotatedX(angle float64) *Shape {
	s1 := s.Clone()
	s1.RotateX(angle)
	return s1
}

// RotatedY returns a copy of this shape, rotated on the y-axis.
func (s *Shape) RotatedY(angle float64) *Shape {
	s1 := s.Clone()
	s1.RotateY(angle)
	return s1
}

// RotatedZ returns a copy of this shape, rotated on the z-axis.
func (s *Shape) RotatedZ(angle float64) *Shape {
	s1 := s.Clone()
	s1.RotateZ(angle)
	return s1
}

// Rotated returns a copy of this shape, rotated.
func (s *Shape) Rotated(rx, ry, rz float64) *Shape {
	s1 := s.Clone()
	s1.Rotate(rx, ry, rz)
	return s1
}

// ScaledX returns a copy of this shape, scaled on the x-axis.
func (s *Shape) ScaledX(scale float64) *Shape {
	s1 := s.Clone()
	s1.ScaleX(scale)
	return s1
}

// ScaledY returns a copy of this shape, scaled on the y-axis.
func (s *Shape) ScaledY(scale float64) *Shape {
	s1 := s.Clone()
	s1.ScaleY(scale)
	return s1
}

// ScaledZ returns a copy of this shape, scaled on the z-axis.
func (s *Shape) ScaledZ(scale float64) *Shape {
	s1 := s.Clone()
	s1.ScaleZ(scale)
	return s1
}

// Scaled returns a copy of this shape, scaled.
func (s *Shape) Scaled(sx, sy, sz float64) *Shape {
	s1 := s.Clone()
	s1.Scale(sx, sy, sz)
	return s1
}

// UniScaled returns a copy of this shape, scaled.
func (s *Shape) UniScaled(scale float64) *Shape {
	s1 := s.Clone()
	s1.UniScale(scale)
	return s1
}

// RandomizedX returns a copy of this shape, randomized on the x-axis.
func (s *Shape) RandomizedX(amount float64) *Shape {
	s1 := s.Clone()
	s1.RandomizeX(amount)
	return s1
}

// RandomizedY returns a copy of this shape, randomized on the y-axis.
func (s *Shape) RandomizedY(amount float64) *Shape {
	s1 := s.Clone()
	s1.RandomizeY(amount)
	return s1
}

// RandomizedZ returns a copy of this shape, randomized on the z-axis.
func (s *Shape) RandomizedZ(amount float64) *Shape {
	s1 := s.Clone()
	s1.RandomizeZ(amount)
	return s1
}

// Randomized returns a copy of this shape, randomized.
func (s *Shape) Randomized(amount float64) *Shape {
	s1 := s.Clone()
	s1.Randomize(amount)
	return s1
}

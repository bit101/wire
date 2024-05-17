// Package wire implements wireframe 3d shapes.
package wire

import (
	"slices"
)

// Shape is a 3d shape composed of a list of points and segments connecting them.
type Shape struct {
	Points   PointList
	Segments []*Segment
}

// NewShape creates a new shape.
func NewShape() *Shape {
	return &Shape{
		PointList{},
		[]*Segment{},
	}
}

// AddShape adds the points and segments of another shape to this shape.
// Does not clone the original shape, so transforms to this shape
// will affect the added shape as well.
func (s *Shape) AddShape(shape *Shape) {
	s.Points = append(s.Points, shape.Points...)
	s.Segments = append(s.Segments, shape.Segments...)
}

// AddPoint adds a point to the shape.
func (s *Shape) AddPoint(point *Point) {
	s.Points.Add(point)
}

// AddXYZ adds a point to the shape.
func (s *Shape) AddXYZ(x, y, z float64) {
	s.Points.AddXYZ(x, y, z)
}

// AddSegment adds a new segment.
func (s *Shape) AddSegment(seg *Segment) {
	s.Segments = append(s.Segments, seg)
}

// AddSegmentByPoints adds a new segment based on the two points passed.
func (s *Shape) AddSegmentByPoints(a, b *Point) {
	seg := NewSegment(a, b)
	s.Segments = append(s.Segments, seg)
}

// AddSegmentByIndex adds a new segment based on the indexes of the two points passed.
func (s *Shape) AddSegmentByIndex(a, b int) {
	seg := NewSegment(s.Points[a], s.Points[b])
	s.Segments = append(s.Segments, seg)
}

// AddRandomPointInBox creates and adds a new 3d point within a 3d box of the given dimensions.
// The box is centered on the origin, so points will range from -w/2 to w/2, etc. on each dimension.
func (s *Shape) AddRandomPointInBox(w, h, d float64) {
	s.AddPoint(RandomPointInBox(w, h, d))
}

// AddRandomPointOnSphere creates and adds a random 3d point ON a sphere of the given radius.
func (s *Shape) AddRandomPointOnSphere(radius float64) {
	s.AddPoint(RandomPointOnSphere(radius))
}

// AddRandomPointInSphere creates and adds a random 3d point IN a sphere of the given radius.
func (s *Shape) AddRandomPointInSphere(radius float64) {
	s.AddPoint(RandomPointInSphere(radius))
}

// AddRandomPointOnCylinder creates and adds a random 3d point ON a cylinder of the given radius and height.
func (s *Shape) AddRandomPointOnCylinder(height, radius float64) {
	s.AddPoint(RandomPointOnCylinder(height, radius))
}

// AddRandomPointInCylinder creates and adds a random 3d point IN a cylinder of the given radius and height.
func (s *Shape) AddRandomPointInCylinder(height, radius float64) {
	s.AddPoint(RandomPointInCylinder(height, radius))
}

// AddRandomPointOnTorus creates and adds a random 3d point ON a torus.
// radius1 is from the center of the torus to the center of the circle forming the torus.
// radius2 is the radius of the circle forming the torus.
func (s *Shape) AddRandomPointOnTorus(radius1, radius2 float64) {
	s.AddPoint(RandomPointOnTorus(radius1, radius2))
}

// AddRandomPointInTorus creates and adds a random 3d point IN a torus.
// radius1 is from the center of the torus to the center of the circle forming the torus.
// radius2 is the radius of the circle forming the torus.
func (s *Shape) AddRandomPointInTorus(radius1, radius2, arc float64) {
	s.AddPoint(RandomPointInTorus(radius1, radius2, arc))
}

// Clone returns a deep copy of this shape.
func (s *Shape) Clone() *Shape {
	clone := NewShape()
	clone.Points = s.Points.Clone()
	for _, seg := range s.Segments {
		indexA := slices.Index(s.Points, seg.PointA)
		indexB := slices.Index(s.Points, seg.PointB)
		clone.AddSegmentByIndex(indexA, indexB)
	}
	return clone
}

// Stroke strokes each path in a shape.
func (s *Shape) Stroke() {
	s.Points.Project()
	for _, segment := range s.Segments {
		segment.Stroke()
	}
}

// RenderPoints draws a filled circle for each point in the path.
func (s *Shape) RenderPoints(radius float64) {
	s.Points.Project()
	s.Points.RenderPoints(radius)
}

// Subdivide puts a new point between each pair of points.
func (s *Shape) Subdivide(times int) {
	for range times {
		for _, seg := range s.Segments {
			a := seg.PointA
			b := seg.PointB
			x := (a.X + b.X) / 2
			y := (a.Y + b.Y) / 2
			z := (a.Z + b.Z) / 2
			p := NewPoint(x, y, z)
			s.AddPoint(p)
			seg.PointB = p
			newSeg := NewSegment(p, b)
			s.Segments = append(s.Segments, newSeg)
		}
	}
}

// Cull removes points from the shape that do not satisfy the cull function. Modifies shape in place.
// TODO: cull segments not just points
func (s *Shape) Cull(cullFunc func(*Point) bool) {
	s.Points.Cull(cullFunc)
}

// Culled returns a new shape with points removed that do not satisfy the cull function.
// TODO: cull segments not just points
func (s *Shape) Culled(cullFunc func(*Point) bool) *Shape {
	s1 := s.Clone()
	s1.Cull(cullFunc)
	return s1
}

// CullBox removes points that ar not within the defined box. Modifies the shape in place.
// TODO: cull segments not just points
func (s *Shape) CullBox(minX, minY, minZ, maxX, maxY, maxZ float64) {
	s.Points.CullBox(minX, minY, minZ, maxX, maxY, maxZ)
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// TranslateX translates this shape on the x-axis, in place.
func (s *Shape) TranslateX(tx float64) {
	s.Points.TranslateX(tx)
}

// TranslateY translates this shape on the y-axis, in place.
func (s *Shape) TranslateY(ty float64) {
	s.Points.TranslateY(ty)
}

// TranslateZ translates this shape on the z-axis, in place.
func (s *Shape) TranslateZ(tz float64) {
	s.Points.TranslateZ(tz)
}

// Translate translates this shape on all axes, in place.
func (s *Shape) Translate(tx, ty, tz float64) {
	s.Points.Translate(tx, ty, tz)
}

// RotateX rotates this shape around the x-axis, in place.
func (s *Shape) RotateX(angle float64) {
	s.Points.RotateX(angle)
}

// RotateY rotates this shape around the y-axis, in place.
func (s *Shape) RotateY(angle float64) {
	s.Points.RotateY(angle)
}

// RotateZ rotates this shape around the z-axis, in place.
func (s *Shape) RotateZ(angle float64) {
	s.Points.RotateZ(angle)
}

// Rotate rotates this shape around all axes, in place.
func (s *Shape) Rotate(rx, ry, rz float64) {
	s.Points.Rotate(rx, ry, rz)
}

// ScaleX scales this shape on the x-axis, in place.
func (s *Shape) ScaleX(scale float64) {
	s.Points.ScaleX(scale)
}

// ScaleY scales this shape on the y-axis, in place.
func (s *Shape) ScaleY(scale float64) {
	s.Points.ScaleY(scale)
}

// ScaleZ scales this shape on the z-axis, in place.
func (s *Shape) ScaleZ(scale float64) {
	s.Points.ScaleZ(scale)
}

// Scale scales this shape on all axes, in place.
func (s *Shape) Scale(sx, sy, sz float64) {
	s.Points.Scale(sx, sy, sz)
}

// UniScale scales this shape by the same amount on each axis, in place.
func (s *Shape) UniScale(scale float64) {
	s.Points.UniScale(scale)
}

// RandomizeX randomizes this shape on the x-axis, in place.
func (s *Shape) RandomizeX(amount float64) {
	s.Points.RandomizeX(amount)
}

// RandomizeY randomizes this shape on the y-axis, in place.
func (s *Shape) RandomizeY(amount float64) {
	s.Points.RandomizeY(amount)
}

// RandomizeZ randmizes this shape on the z-axis, in place.
func (s *Shape) RandomizeZ(amount float64) {
	s.Points.RandomizeZ(amount)
}

// Randomize randomizes this shape on all axes, in place.
func (s *Shape) Randomize(amount float64) {
	s.Points.Randomize(amount)
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

// Translated returns a copy of this shape, translated on all axes.
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

// Rotated returns a copy of this shape, rotated on all axes.
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

// Scaled returns a copy of this shape, scaled on all axes.
func (s *Shape) Scaled(sx, sy, sz float64) *Shape {
	s1 := s.Clone()
	s1.Scale(sx, sy, sz)
	return s1
}

// UniScaled returns a copy of this shape, scaled by the same amount on each axis.
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

// Randomized returns a copy of this shape, randomized on all axes.
func (s *Shape) Randomized(amount float64) *Shape {
	s1 := s.Clone()
	s1.Randomize(amount)
	return s1
}

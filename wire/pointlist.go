// Package wire implements wireframe 3d shapes.
package wire

import (
	"github.com/bit101/bitlib/geom"
)

// PointList represents a list of 3d points.
type PointList []*Point

// NewPointList creates a new PointList.
func NewPointList() PointList {
	return PointList{}
}

// Clone returns a copy of this pointlist.
func (p PointList) Clone() PointList {
	list := NewPointList()
	for _, point := range p {
		list.Add(point.Clone())
	}
	return list
}

// Add adds a new point to this list.
func (p *PointList) Add(point *Point) {
	*p = append(*p, point)
}

// AddXYZ adds a new point to this list.
func (p *PointList) AddXYZ(x, y, z float64) {
	*p = append(*p, NewPoint(x, y, z))
}

// Project projects this 3d point list to a 2d point list.
func (p PointList) Project() (geom.PointList, []float64) {
	size := len(p)
	list := make(geom.PointList, size)
	scales := make([]float64, size)
	for i, point := range p {
		gp, scale := point.Project()
		list[i] = gp
		scales[i] = scale
	}
	return list, scales
}

// Stroke strokes a path on a point list.
func (p PointList) Stroke(context Context, closed bool) {
	points, _ := p.Project()
	context.StrokePath(points, closed)
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// Translate translates this pointlist in place.
func (p PointList) Translate(tx, ty, tz float64) {
	for _, point := range p {
		point.Translate(tx, ty, tz)
	}
}

// RotateX rotates this pointlist around the x-axis in place.
func (p PointList) RotateX(angle float64) {
	for _, point := range p {
		point.RotateX(angle)
	}
}

// RotateY rotates this pointlist around the y-axis in place.
func (p PointList) RotateY(angle float64) {
	for _, point := range p {
		point.RotateY(angle)
	}
}

// RotateZ rotates this pointlist around the z-axis in place.
func (p PointList) RotateZ(angle float64) {
	for _, point := range p {
		point.RotateZ(angle)
	}
}

// Rotate rotates this pointlist in place.
func (p PointList) Rotate(rx, ry, rz float64) {
	for _, point := range p {
		point.Rotate(rx, ry, rz)
	}
}

// Scale scales this pointlist in place.
func (p PointList) Scale(sx, sy, sz float64) {
	for _, point := range p {
		point.Scale(sx, sy, sz)
	}
}

// UniScale scales this pointlist in place.
func (p PointList) UniScale(scale float64) {
	for _, point := range p {
		point.UniScale(scale)
	}
}

// Randomize randomizes this pointlist in place.
func (p PointList) Randomize(amount float64) {
	for _, point := range p {
		point.Randomize(amount)
	}
}

//////////////////////////////
// Transform and return new
//////////////////////////////

// Translated returns a copy of this pointlist, translated.
func (p PointList) Translated(tx, ty, tz float64) PointList {
	p1 := p.Clone()
	p1.Translate(tx, ty, tz)
	return p1
}

// RotatedX returns a copy of this pointlist, rotated on the x-axis.
func (p PointList) RotatedX(angle float64) PointList {
	p1 := p.Clone()
	p1.RotateX(angle)
	return p1
}

// RotatedY returns a copy of this pointlist, rotated on the y-axis.
func (p PointList) RotatedY(angle float64) PointList {
	p1 := p.Clone()
	p1.RotateY(angle)
	return p1
}

// RotatedZ returns a copy of this pointlist, rotated on the z-axis.
func (p PointList) RotatedZ(angle float64) PointList {
	p1 := p.Clone()
	p1.RotateZ(angle)
	return p1
}

// Rotated returns a copy of this pointlist, rotated.
func (p PointList) Rotated(rx, ry, rz float64) PointList {
	p1 := p.Clone()
	p1.Rotate(rx, ry, rz)
	return p1
}

// Scaled returns a copy of this pointlist, scaled.
func (p PointList) Scaled(sx, sy, sz float64) PointList {
	p1 := p.Clone()
	p1.Scale(sx, sy, sz)
	return p1
}

// UniScaled returns a copy of this pointlist, scaled.
func (p PointList) UniScaled(scale float64) PointList {
	p1 := p.Clone()
	p1.UniScale(scale)
	return p1
}

// Randomized returns a copy of this pointlist, randomized.
func (p PointList) Randomized(amount float64) PointList {
	p1 := p.Clone()
	p1.Randomize(amount)
	return p1
}

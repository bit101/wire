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

func shouldDraw(p0, p1 *Point) bool {
	return Visible(p0) && Visible(p1)
}

// Stroke strokes a path on a point list.
func (p PointList) Stroke(context Context, closed bool) {
	points, _ := p.Project()
	for i := 0; i < len(points)-1; i++ {
		if shouldDraw(p[i], p[i+1]) {
			p0 := points[i]
			p1 := points[i+1]
			context.MoveTo(p0.X, p0.Y)
			context.LineTo(p1.X, p1.Y)
			context.Stroke()
		}
	}
	if closed && shouldDraw(p[0], p.Last()) {
		p0 := points[0]
		p1 := points[len(points)-1]
		context.MoveTo(p0.X, p0.Y)
		context.LineTo(p1.X, p1.Y)
		context.Stroke()
	}
}

// Get returns the point at the given index. Negative indexes go in reverse from end.
func (p PointList) Get(index int) *Point {
	if index < 0 {
		index = len(p) - index
	}
	return p[index]
}

// Last returns the last point in the list.
func (p PointList) Last() *Point {
	return p[len(p)-1]
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// TranslateX translates this pointlist on the x-axis in place.
func (p PointList) TranslateX(tx float64) {
	for _, point := range p {
		point.TranslateX(tx)
	}
}

// TranslateY translates this pointlist on the y-axis in place.
func (p PointList) TranslateY(ty float64) {
	for _, point := range p {
		point.TranslateY(ty)
	}
}

// TranslateZ translates this pointlist on the z-axis in place.
func (p PointList) TranslateZ(tz float64) {
	for _, point := range p {
		point.TranslateX(tz)
	}
}

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

// ScaleX scales this pointlist on the x-axis, in place.
func (p PointList) ScaleX(scale float64) {
	for _, point := range p {
		point.ScaleX(scale)
	}
}

// ScaleY scales this pointlist on the y-axis, in place.
func (p PointList) ScaleY(scale float64) {
	for _, point := range p {
		point.ScaleY(scale)
	}
}

// ScaleZ scales this pointlist on the z-axis, in place.
func (p PointList) ScaleZ(scale float64) {
	for _, point := range p {
		point.ScaleZ(scale)
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

// RandomizeX randomizes this pointlist on the x-axis, in place.
func (p PointList) RandomizeX(amount float64) {
	for _, point := range p {
		point.RandomizeX(amount)
	}
}

// RandomizeY randomizes this pointlist on the y-axis, in place.
func (p PointList) RandomizeY(amount float64) {
	for _, point := range p {
		point.RandomizeY(amount)
	}
}

// RandomizeZ randomizes this pointlist on the z-axis, in place.
func (p PointList) RandomizeZ(amount float64) {
	for _, point := range p {
		point.RandomizeZ(amount)
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

// TranslatedX returns a copy of this pointlist, translated on the x-axis.
func (p *PointList) TranslatedX(tx float64) PointList {
	p1 := p.Clone()
	p1.TranslateX(tx)
	return p1
}

// TranslatedY returns a copy of this pointlist, translated on the y-axis.
func (p *PointList) TranslatedY(ty float64) PointList {
	p1 := p.Clone()
	p1.TranslateY(ty)
	return p1
}

// TranslatedZ returns a copy of this pointlist, translated on the z-axis.
func (p *PointList) TranslatedZ(tz float64) PointList {
	p1 := p.Clone()
	p1.TranslateZ(tz)
	return p1
}

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

// ScaledX returns a copy of this pointlist, scaled on the x-axis.
func (p PointList) ScaledX(scale float64) PointList {
	p1 := p.Clone()
	p1.ScaleX(scale)
	return p1
}

// ScaledY returns a copy of this pointlist, scaled on the x-axis.
func (p PointList) ScaledY(scale float64) PointList {
	p1 := p.Clone()
	p1.ScaleY(scale)
	return p1
}

// ScaledZ returns a copy of this pointlist, scaled on the x-axis.
func (p PointList) ScaledZ(scale float64) PointList {
	p1 := p.Clone()
	p1.ScaleZ(scale)
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

// RandomizedX returns a copy of this pointlist, randomized on the x-axis
func (p PointList) RandomizedX(amount float64) PointList {
	p1 := p.Clone()
	p1.RandomizeX(amount)
	return p1
}

// RandomizedY returns a copy of this pointlist, randomized on the y-axis
func (p PointList) RandomizedY(amount float64) PointList {
	p1 := p.Clone()
	p1.RandomizeY(amount)
	return p1
}

// RandomizedZ returns a copy of this pointlist, randomized on the z-axis
func (p PointList) RandomizedZ(amount float64) PointList {
	p1 := p.Clone()
	p1.RandomizeZ(amount)
	return p1
}

// Randomized returns a copy of this pointlist, randomized.
func (p PointList) Randomized(amount float64) PointList {
	p1 := p.Clone()
	p1.Randomize(amount)
	return p1
}

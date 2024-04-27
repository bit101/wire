// Package wire implements wireframe 3d shapes.
package wire

import "log"

// PathList is a list of lists of 3d points.
type PathList []PointList

// NewPathList creates a new pathlist.
func NewPathList(size int) PathList {
	return make(PathList, size)
}

// Add adds a new path to the list.
func (p *PathList) Add(path PointList) {
	*p = append(*p, path)
}

// AddPoint adds a point to the list at the given index.
func (p *PathList) AddPoint(index int, point *Point) {
	if index >= len(*p) || index < 0 {
		log.Fatal("list index must be from 0 to one less than the size of the list")
	}
	(*p)[index].Add(point)
}

// AddXYZ adds a point to the list at the given index.
func (p *PathList) AddXYZ(index int, x, y, z float64) {
	if index >= len(*p) || index < 0 {
		log.Fatal("list index must be from 0 to one less than the size of the list")
	}
	(*p)[index].AddXYZ(x, y, z)
}

// Clone returns a deep copy of this pathlist.
func (p PathList) Clone() PathList {
	list := NewPathList(0)
	for _, pointList := range p {
		list = append(list, pointList.Clone())
	}
	return list
}

// Stroke strokes each path in a pathlist.
func (p PathList) Stroke(context Context, closed bool) {
	for _, path := range p {
		path.Stroke(context, closed)
	}
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// Translate translates this pathlist in place.
func (p PathList) Translate(tx, ty, tz float64) {
	for _, list := range p {
		list.Translate(tx, ty, tz)
	}
}

// RotateX rotates this pathlist around the x-axis in place.
func (p PathList) RotateX(angle float64) {
	for _, list := range p {
		list.RotateX(angle)
	}
}

// RotateY rotates this pathlist around the y-axis in place.
func (p PathList) RotateY(angle float64) {
	for _, list := range p {
		list.RotateY(angle)
	}
}

// RotateZ rotates this pathlist around the z-axis in place.
func (p PathList) RotateZ(angle float64) {
	for _, list := range p {
		list.RotateZ(angle)
	}
}

// Rotate rotates this pathlist in place.
func (p PathList) Rotate(rx, ry, rz float64) {
	for _, list := range p {
		list.Rotate(rx, ry, rz)
	}
}

// Scale scales this pathlist in place.
func (p PathList) Scale(sx, sy, sz float64) {
	for _, list := range p {
		list.Scale(sx, sy, sz)
	}
}

// UniScale scales this pathlist in place.
func (p PathList) UniScale(scale float64) {
	for _, list := range p {
		list.UniScale(scale)
	}
}

// Randomize randomizes this pathlist in place.
func (p PathList) Randomize(amount float64) {
	for _, list := range p {
		list.Randomize(amount)
	}
}

//////////////////////////////
// Transform and return new
//////////////////////////////

// Translated returns a copy of this pathlist, translated.
func (p PathList) Translated(tx, ty, tz float64) PathList {
	p1 := p.Clone()
	p1.Translate(tx, ty, tz)
	return p1
}

// RotatedX returns a copy of this pathlist, rotated on the x-axis.
func (p PathList) RotatedX(angle float64) PathList {
	p1 := p.Clone()
	p1.RotateX(angle)
	return p1
}

// RotatedY returns a copy of this pathlist, rotated on the y-axis.
func (p PathList) RotatedY(angle float64) PathList {
	p1 := p.Clone()
	p1.RotateY(angle)
	return p1
}

// RotatedZ returns a copy of this pathlist, rotated on the z-axis.
func (p PathList) RotatedZ(angle float64) PathList {
	p1 := p.Clone()
	p1.RotateZ(angle)
	return p1
}

// Rotated returns a copy of this pathlist, rotated.
func (p PathList) Rotated(rx, ry, rz float64) PathList {
	p1 := p.Clone()
	p1.Rotate(rx, ry, rz)
	return p1
}

// Scaled returns a copy of this pathlist, scaled.
func (p PathList) Scaled(sx, sy, sz float64) PathList {
	p1 := p.Clone()
	p1.Scale(sx, sy, sz)
	return p1
}

// UniScaled returns a copy of this pathlist, scaled.
func (p PathList) UniScaled(scale float64) PathList {
	p1 := p.Clone()
	p1.UniScale(scale)
	return p1
}

// Randomized returns a copy of this pathlist, randomized.
func (p PathList) Randomized(amount float64) PathList {
	p1 := p.Clone()
	p1.Randomize(amount)
	return p1
}

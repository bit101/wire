// Package wire implements wireframe 3d shapes.
package wire

import "github.com/bit101/bitlib/noise"

// PointList represents a list of 3d points.
type PointList []*Point

// NewPointList creates a new PointList.
func NewPointList() PointList {
	return PointList{}
}

// Clone returns a deep copy of this pointlist.
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

// AddXYZ creates and adds a new point to this list.
func (p *PointList) AddXYZ(x, y, z float64) {
	*p = append(*p, NewPoint(x, y, z))
}

// AddRandomPointInBox creates and adds a new 3d point within a 3d box of the given dimensions.
// The box is centered on the origin, so points will range from -w/2 to w/2, etc. on each dimension.
func (p *PointList) AddRandomPointInBox(w, h, d float64) {
	p.Add(RandomPointInBox(w, h, d))
}

// AddRandomPointOnSphere creates and adds a random 3d point ON a sphere of the given radius.
func (p *PointList) AddRandomPointOnSphere(radius float64) {
	p.Add(RandomPointOnSphere(radius))
}

// AddRandomPointInSphere creates and adds a random 3d point IN a sphere of the given radius.
func (p *PointList) AddRandomPointInSphere(radius float64) {
	p.Add(RandomPointInSphere(radius))
}

// AddRandomPointOnCylinder creates and adds a random 3d point ON a cylinder of the given radius and height.
func (p *PointList) AddRandomPointOnCylinder(height, radius float64, includeCaps bool) {
	p.Add(RandomPointOnCylinder(height, radius, includeCaps))
}

// AddRandomPointInCylinder creates and adds a random 3d point IN a cylinder of the given radius and height.
func (p *PointList) AddRandomPointInCylinder(height, radius float64) {
	p.Add(RandomPointInCylinder(height, radius))
}

// AddRandomPointOnTorus creates and adds a random 3d point ON a torus.
// radius1 is from the center of the torus to the center of the circle forming the torus.
// radius2 is the radius of the circle forming the torus.
func (p *PointList) AddRandomPointOnTorus(radius1, radius2 float64) {
	p.AddRandomPointOnTorus(radius1, radius2)
}

// AddRandomPointInTorus creates and adds a random 3d point IN a torus.
// radius1 is from the center of the torus to the center of the circle forming the torus.
// radius2 is the radius of the circle forming the torus.
func (p *PointList) AddRandomPointInTorus(radius1, radius2 float64) {
	p.AddRandomPointInTorus(radius1, radius2)
}

// Project projects this 3d point list to a 2d point list.
// This returns a list of 2d points as well as a list of scale values for each point.
func (p PointList) Project() {
	for _, point := range p {
		point.Project()
	}
}

// RenderPoints projects and draws a circle for each point in the list.
func (p PointList) RenderPoints(radius float64) {
	p.Project()
	for _, point := range p {
		if point.Visible() {
			world.Context.Save()
			ApplyFog(point.Z)
			world.Context.FillCircle(point.Px, point.Py, radius*point.Scaling)
			world.Context.Restore()
		}
	}
}

// Get returns the point at the given index. Negative indexes go in reverse from end.
func (p PointList) Get(index int) *Point {
	if index < 0 {
		index = len(p) - index
	}
	return p[index]
}

// First returns the first point in the list.
func (p PointList) First() *Point {
	return p[0]
}

// Last returns the last point in the list.
func (p PointList) Last() *Point {
	return p[len(p)-1]
}

// Cull removes points from the list that do not satisfy the cull function. Modifies list in place.
func (p *PointList) Cull(cullFunc func(*Point) bool) {
	newList := NewPointList()
	for _, point := range *p {
		if cullFunc(point) {
			newList.Add(point)
		}
	}
	*p = newList
}

// Culled returns a new point list with points removed that do not satisfy the cull function.
func (p PointList) Culled(cullFunc func(*Point) bool) PointList {
	p1 := p.Clone()
	p1.Cull(cullFunc)
	return p1
}

// CullBox removes points that ar not within the defined box. Modifies the shape in place.
func (p *PointList) CullBox(minX, minY, minZ, maxX, maxY, maxZ float64) {
	newList := NewPointList()
	for _, point := range *p {
		if point.X >= minX && point.X <= maxX &&
			point.Y >= minY && point.Y <= maxY &&
			point.Z >= minZ && point.Z <= maxZ {
			newList.Add(point)
		}
	}
	*p = newList
}

//////////////////////////////
// Transform in place.
//////////////////////////////

// TranslateX translates each point in this pointlist on the x-axis, in place.
func (p PointList) TranslateX(tx float64) {
	for _, point := range p {
		point.TranslateX(tx)
	}
}

// TranslateY translates each point in this pointlist on the y-axis, in place.
func (p PointList) TranslateY(ty float64) {
	for _, point := range p {
		point.TranslateY(ty)
	}
}

// TranslateZ translates each point in this pointlist on the z-axis, in place.
func (p PointList) TranslateZ(tz float64) {
	for _, point := range p {
		point.TranslateZ(tz)
	}
}

// Translate translates each point in this pointlist on all axes, in place.
func (p PointList) Translate(tx, ty, tz float64) {
	for _, point := range p {
		point.Translate(tx, ty, tz)
	}
}

// RotateX rotates each point in this pointlist around the x-axis, in place.
func (p PointList) RotateX(angle float64) {
	for _, point := range p {
		point.RotateX(angle)
	}
}

// RotateY rotates each point in this pointlist around the y-axis, in place.
func (p PointList) RotateY(angle float64) {
	for _, point := range p {
		point.RotateY(angle)
	}
}

// RotateZ rotates each point in this pointlist around the z-axis, in place.
func (p PointList) RotateZ(angle float64) {
	for _, point := range p {
		point.RotateZ(angle)
	}
}

// Rotate rotates each point in this pointlist around all axes, in place.
func (p PointList) Rotate(rx, ry, rz float64) {
	for _, point := range p {
		point.Rotate(rx, ry, rz)
	}
}

// ScaleX scales each point in this pointlist on the x-axis, in place.
func (p PointList) ScaleX(scale float64) {
	for _, point := range p {
		point.ScaleX(scale)
	}
}

// ScaleY scales each point in this pointlist on the y-axis, in place.
func (p PointList) ScaleY(scale float64) {
	for _, point := range p {
		point.ScaleY(scale)
	}
}

// ScaleZ scales each point in this pointlist on the z-axis, in place.
func (p PointList) ScaleZ(scale float64) {
	for _, point := range p {
		point.ScaleZ(scale)
	}
}

// Scale scales each point in this pointlist on all axes, in place.
func (p PointList) Scale(sx, sy, sz float64) {
	for _, point := range p {
		point.Scale(sx, sy, sz)
	}
}

// UniScale scales each point in this pointlist by the same amount on each axis, in place.
func (p PointList) UniScale(scale float64) {
	for _, point := range p {
		point.UniScale(scale)
	}
}

// RandomizeX randomizes each point in this pointlist on the x-axis, in place.
func (p PointList) RandomizeX(amount float64) {
	for _, point := range p {
		point.RandomizeX(amount)
	}
}

// RandomizeY randomizes each point in this pointlist on the y-axis, in place.
func (p PointList) RandomizeY(amount float64) {
	for _, point := range p {
		point.RandomizeY(amount)
	}
}

// RandomizeZ randomizes each point in this pointlist on the z-axis, in place.
func (p PointList) RandomizeZ(amount float64) {
	for _, point := range p {
		point.RandomizeZ(amount)
	}
}

// Randomize randomizes each point in this pointlist on all axes, in place.
func (p PointList) Randomize(amount float64) {
	for _, point := range p {
		point.Randomize(amount)
	}
}

// Push pushes points away from the specified point.
func (p PointList) Push(pusher *Point, radius float64) {
	for _, point := range p {
		dist := point.Distance(pusher)
		if dist < radius {
			point.X -= pusher.X
			point.Y -= pusher.Y
			point.Z -= pusher.Z
			point.UniScale(radius / dist)
			point.X += pusher.X
			point.Y += pusher.Y
			point.Z += pusher.Z
		}
	}
}

func (p PointList) Noisify(origin *Point, scale, offset float64) {
	for _, point := range p {
		n := noise.Simplex3(
			origin.X+point.X*scale,
			origin.Y+point.Y*scale,
			origin.Z+point.Z*scale,
		)
		point.UniScale(1.0 + n*offset)
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

// Translated returns a copy of this pointlist, translated on all axes.
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

// Rotated returns a copy of this pointlist, rotated on all axes.
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

// ScaledY returns a copy of this pointlist, scaled on the y-axis.
func (p PointList) ScaledY(scale float64) PointList {
	p1 := p.Clone()
	p1.ScaleY(scale)
	return p1
}

// ScaledZ returns a copy of this pointlist, scaled on the z-axis.
func (p PointList) ScaledZ(scale float64) PointList {
	p1 := p.Clone()
	p1.ScaleZ(scale)
	return p1
}

// Scaled returns a copy of this pointlist, scaled on all axes.
func (p PointList) Scaled(sx, sy, sz float64) PointList {
	p1 := p.Clone()
	p1.Scale(sx, sy, sz)
	return p1
}

// UniScaled returns a copy of this pointlist, scaled by the same amount on each axis.
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

// Randomized returns a copy of this pointlist, randomized on all axes.
func (p PointList) Randomized(amount float64) PointList {
	p1 := p.Clone()
	p1.Randomize(amount)
	return p1
}

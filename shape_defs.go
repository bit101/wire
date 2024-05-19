// Package wire implements wireframe 3d shapes.
package wire

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/bitlib/random"
)

// Box creates a 3d box shape.
func Box(w, h, d float64) *Shape {
	shape := NewShape()
	shape.AddXYZ(-1, -1, -1)
	shape.AddXYZ(-1, 1, -1)
	shape.AddXYZ(1, 1, -1)
	shape.AddXYZ(1, -1, -1)
	shape.AddXYZ(-1, -1, 1)
	shape.AddXYZ(-1, 1, 1)
	shape.AddXYZ(1, 1, 1)
	shape.AddXYZ(1, -1, 1)

	shape.AddSegmentByIndex(0, 1)
	shape.AddSegmentByIndex(1, 2)
	shape.AddSegmentByIndex(2, 3)
	shape.AddSegmentByIndex(3, 0)

	shape.AddSegmentByIndex(4, 5)
	shape.AddSegmentByIndex(5, 6)
	shape.AddSegmentByIndex(6, 7)
	shape.AddSegmentByIndex(7, 4)

	shape.AddSegmentByIndex(0, 4)
	shape.AddSegmentByIndex(1, 5)
	shape.AddSegmentByIndex(2, 6)
	shape.AddSegmentByIndex(3, 7)
	shape.Scale(w/2, h/2, d/2)
	return shape
}

// CirclePath creates a single path defining a circle.
func CirclePath(radius float64, res int) (PointList, []*Segment) {
	points := NewPointList()
	segments := []*Segment{}
	for i := 0; i < res; i++ {
		t := blmath.Tau * float64(i) / float64(res)
		p := NewPoint(math.Cos(t)*radius, 0, math.Sin(t)*radius)
		points.Add(p)
		if i > 0 {
			s := NewSegment(points[i-1], points[i])
			segments = append(segments, s)
		}
	}
	segments = append(segments, NewSegment(points.Last(), points.First()))
	return points, segments
}

// Circle creates a 3d cone shape made of a number of circular slices.
func Circle(radius float64, res int) *Shape {
	shape := NewShape()
	p, s := CirclePath(radius, res)
	shape.Points = p
	shape.Segments = s
	return shape
}

// Cone creates a 3d cone shape made of a number of circular slices.
func Cone(height, radius0, radius1 float64, slices, res int, showSlices, showLong bool) *Shape {
	shape := NewShape()
	for i := 0; i < slices; i++ {
		radius := blmath.Map(float64(i), 0, float64(slices-1), radius0, radius1)
		p, s := CirclePath(radius, res)
		y := float64(i)/(float64(slices)-1)*height - height/2
		p.TranslateY(y)
		if showSlices {
			shape.Segments = append(shape.Segments, s...)
		}
		shape.Points = append(shape.Points, p...)
	}
	if showLong {
		for i := range slices - 1 {
			for j := range res {
				index0 := i*res + j
				index1 := (i+1)*res + j
				shape.AddSegmentByIndex(index0, index1)
			}
		}
	}
	return shape
}

// Cylinder creates a 3d cylinder shape made of a number of circular slices.
func Cylinder(height, radius float64, slices, res int, showSlices, showLong bool) *Shape {
	return Cone(height, radius, radius, slices, res, showSlices, showLong)
}

// GridPlane creates a 3d plane containing a grid.
func GridPlane(w, d float64, rows, cols int) *Shape {
	shape := NewShape()
	for x := 0.0; x <= float64(rows); x++ {
		for z := 0.0; z <= float64(cols); z++ {
			shape.Points = append(shape.Points, NewPoint(x, 0, z))
		}
	}
	for x := 0; x <= rows; x++ {
		for z := 0; z <= cols; z++ {
			index := x + z*(rows+1)
			if x < rows {
				shape.AddSegmentByIndex(index, index+1)
			}
			if z < cols {
				shape.AddSegmentByIndex(index, index+rows+1)
			}
		}
	}
	shape.Scale(w/float64(rows), 1, d/float64(cols))
	shape.Translate(-w/2, 0, -d/2)

	return shape
}

// Pyramid creates a 3d pyramid shape.
func Pyramid(height, baseRadius float64, sides int) *Shape {
	return Cone(height, 0, baseRadius, 2, sides, true, true)
}

func RandomInnerBox(w, h, d float64, count int) *Shape {
	shape := NewShape()
	for range count {
		shape.AddPoint(RandomPointInBox(w, h, d))
	}
	return shape
}

func RandomSurfaceBox(w, h, d float64, count int) *Shape {
	shape := NewShape()
	surface := (d*h + w*d + w*h) * 2
	fcount := float64(count)
	// left/right
	for range int(fcount * (d * h / surface)) {
		shape.AddXYZ(-w/2, random.FloatRange(-h/2, h/2), random.FloatRange(-d/2, d/2))
		shape.AddXYZ(w/2, random.FloatRange(-h/2, h/2), random.FloatRange(-d/2, d/2))
	}
	// top/bottom
	for range int(fcount * (w * d / surface)) {
		shape.AddXYZ(random.FloatRange(-w/2, w/2), -h/2, random.FloatRange(-d/2, d/2))
		shape.AddXYZ(random.FloatRange(-w/2, w/2), h/2, random.FloatRange(-d/2, d/2))
	}
	// front/back
	for range int(fcount * (w * h / surface)) {
		shape.AddXYZ(random.FloatRange(-w/2, w/2), random.FloatRange(-h/2, h/2), -d/2)
		shape.AddXYZ(random.FloatRange(-w/2, w/2), random.FloatRange(-h/2, h/2), d/2)
	}
	return shape
}

// RandomInnerSphere creates a 3d sphere made of random points inside the sphere.
func RandomInnerSphere(radius float64, count int) *Shape {
	shape := NewShape()
	for range count {
		shape.AddRandomPointInSphere(radius)
	}
	return shape
}

// RandomSurfaceSphere creates a 3d sphere made of random points on the surface of the sphere.
func RandomSurfaceSphere(radius float64, count int) *Shape {
	shape := NewShape()
	for range count {
		shape.AddRandomPointOnSphere(radius)
	}
	return shape
}

// Sphere creates a 3d sphere of regular points that can be connected longitudinally, lattitudally, or both.
func Sphere(radius float64, long, lat int, showLong, showLat bool) *Shape {
	shape := NewShape()
	fslice := float64(long)
	for i := 0.0; i <= fslice; i++ {
		a := i / fslice * math.Pi
		p, s := CirclePath(math.Sin(a), lat)
		p.TranslateY(math.Cos(a))
		shape.Points = append(shape.Points, p...)
		if showLat {
			shape.Segments = append(shape.Segments, s...)
		}
	}
	if showLong {
		for i := 0; i < long; i++ {
			for j := 0; j < lat; j++ {
				shape.AddSegmentByIndex(i*lat+j, (i+1)*lat+j)
			}
		}
	}
	shape.UniScale(radius)
	return shape
}

// Spring creates a 3d spiral shape.
func Spring(height, r0, r1, turns, res float64) *Shape {
	shape := NewShape()
	totalAngle := blmath.Tau * turns
	for a := 0.0; a <= totalAngle; a += blmath.Tau / res {
		t := a / totalAngle
		radius := blmath.Lerp(t, r0, r1)
		x := math.Cos(a) * radius
		y := blmath.Lerp(t, height/2, -height/2)
		z := math.Sin(a) * radius
		shape.AddXYZ(x, y, z)
	}
	for i := range len(shape.Points) - 1 {
		shape.AddSegmentByIndex(i, i+1)
	}
	return shape
}

// Torus creates a 3d torus made of a number of circular slices.
func Torus(r1, r2, arc float64, slices, res int, showSlices, showLong bool) *Shape {
	shape := NewShape()
	fslice := float64(slices)
	for i := 0.0; i < fslice; i++ {
		angle := i / fslice * arc
		p, s := CirclePath(r2, res)
		p.RotateX(math.Pi / 2)
		p.TranslateX(r1)
		p.RotateY(angle)
		shape.Points = append(shape.Points, p...)
		if showSlices {
			shape.Segments = append(shape.Segments, s...)
		}
	}
	if showLong {
		for i := range slices - 1 {
			for j := range res {
				index0 := i*res + j
				index1 := (i+1)*res + j
				shape.AddSegmentByIndex(index0, index1)
			}
		}
		if arc >= blmath.Tau {
			i := slices - 1
			for j := range res {
				index0 := i*res + j
				index1 := j
				shape.AddSegmentByIndex(index0, index1)
			}
		}
	}
	return shape
}

// TorusKnot creates a 3d torus knot shape made of one long path that wraps around the torus.
func TorusKnot(p, q, r1, r2, res float64) *Shape {
	shape := NewShape()
	for t := 0.0; t < blmath.Tau; t += res {
		r := math.Cos(q*t) + r1/r2
		x := r * math.Cos(p*t)
		y := -math.Sin(q * t)
		z := r * math.Sin(p*t)
		shape.AddXYZ(
			x*r2,
			y*r2,
			z*r2,
		)
	}
	for i := 0; i < len(shape.Points); i++ {
		shape.AddSegmentByIndex(i, (i+1)%len(shape.Points))
	}
	return shape
}

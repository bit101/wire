// Package wire implements wireframe 3d shapes.
package wire

import (
	"math"

	"github.com/bit101/bitlib/blmath"
)

// Box creates a 3d box shape.
func Box(w, h, d float64) *Shape {
	shape := &Shape{
		make([]PointList, 4),
		false,
	}

	shape.AddXYZ(0, -1, -1, -1)
	shape.AddXYZ(0, -1, 1, -1)
	shape.AddXYZ(0, 1, 1, -1)
	shape.AddXYZ(0, 1, -1, -1)

	shape.AddXYZ(1, -1, 1, -1)
	shape.AddXYZ(1, -1, 1, 1)
	shape.AddXYZ(1, 1, 1, 1)
	shape.AddXYZ(1, 1, 1, -1)

	shape.AddXYZ(2, -1, 1, 1)
	shape.AddXYZ(2, -1, -1, 1)
	shape.AddXYZ(2, 1, -1, 1)
	shape.AddXYZ(2, 1, 1, 1)

	shape.AddXYZ(3, -1, -1, 1)
	shape.AddXYZ(3, -1, -1, -1)
	shape.AddXYZ(3, 1, -1, -1)
	shape.AddXYZ(3, 1, -1, 1)
	shape.Scale(w, h, d)
	return shape
}

// Cone creates a 3d cone shape made of a number of slices.
func Cone(height, radius0, radius1 float64, slices, res int) *Shape {
	shape := &Shape{
		make([]PointList, slices),
		true,
	}
	for i := 0; i < slices; i++ {
		radius := blmath.Map(float64(i), 0, float64(slices-1), radius0, radius1)
		for j := 0; j < res; j++ {
			t := blmath.Tau * float64(j) / float64(res)
			y := float64(i)/(float64(slices)-1)*height - height/2
			shape.AddXYZ(i, math.Cos(t)*radius, y, math.Sin(t)*radius)
		}
	}
	return shape
}

// Cylinder creates a 3d cylinder shape made of a number of slices.
func Cylinder(height, radius float64, slices, res int) *Shape {
	return Cone(height, radius, radius, slices, res)
}

// GridPlane creates a 3d plane containing a grid.
func GridPlane(w, d, res float64) *Shape {
	shape := NewShape(0, false)
	for x := -w / 2; x <= w/2; x += res {
		path := NewPointList()
		path.AddXYZ(x, 0, -d/2)
		path.AddXYZ(x, 0, d/2)
		shape.Add(path)
	}
	for z := -d / 2; z <= d/2; z += res {
		path := NewPointList()
		path.AddXYZ(-w/2, 0, z)
		path.AddXYZ(w/2, 0, z)
		shape.Add(path)
	}
	return shape
}

// Pyramid creates a 3d pyramid shape.
func Pyramid(height, baseRadius float64, sides int) *Shape {
	shape := &Shape{
		[]PointList{},
		true,
	}
	for i := 0; i < sides; i++ {
		side := NewPointList()
		a1 := float64(i) / float64(sides) * blmath.Tau
		x := math.Cos(a1) * baseRadius
		y := height / 2
		z := math.Sin(a1) * baseRadius
		side.AddXYZ(x, y, z)
		a2 := float64(i+1) / float64(sides) * blmath.Tau
		x = math.Cos(a2) * baseRadius
		z = math.Sin(a2) * baseRadius
		side.AddXYZ(x, y, z)
		side.AddXYZ(0, -height/2, 0)
		shape.Add(side)
	}
	return shape
}

// Sphere creates a 3d sphere made of a number of slices.
func Sphere(radius float64, slices, res int) *Shape {
	shape := &Shape{
		[]PointList{},
		true,
	}
	fslice := float64(slices)
	dt := blmath.Tau / float64(res)
	for i := 0.0; i < fslice; i++ {
		path := NewPointList()
		a := i / fslice * math.Pi
		for t := 0.0; t <= blmath.Tau-dt; t += dt {
			y := math.Cos(a)
			r := math.Sin(a)
			path.AddXYZ(math.Cos(t)*r, y, math.Sin(t)*r)
		}
		shape.Add(path)
	}
	shape.UniScale(radius)
	return shape
}

// Torus creates a 3d torus made of a number of slices.
func Torus(r1, r2 float64, slices, res int) *Shape {
	shape := &Shape{
		make([]PointList, slices),
		true,
	}
	fslice := float64(slices)
	dt := blmath.Tau / float64(res)
	for i := 0.0; i < fslice; i++ {
		angle := i / fslice * blmath.Tau
		path := NewPointList()
		for t := 0.0; t <= blmath.Tau-dt; t += dt {
			path.AddXYZ(r1+math.Cos(t)*r2, math.Sin(t)*r2, 0)
		}
		path.RotateY(angle)
		shape.Add(path)
	}
	return shape
}

// TorusKnot creates a 3d torus knot shape made of one long path that wraps around the torus.
func TorusKnot(p, q, scale, res float64) *Shape {
	shape := NewShape(1, false)
	res = 1.0 / res
	for t := 0.0; t < blmath.Tau; t += res {
		r := math.Cos(q*t) + 3
		x := r * math.Cos(p*t)
		y := r * math.Sin(p*t)
		z := -math.Sin(q * t)
		shape.AddXYZ(
			0,
			x*scale,
			y*scale,
			z*scale,
		)
	}
	return shape
}

// Package wire implements wireframe 3d shapes.
package wire

import (
	"math"

	"github.com/bit101/bitlib/blmath"
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

func Box(w, h, d float64) PathList {
	box := NewPathList(4)

	box.AddXYZ(0, -1, -1, -1)
	box.AddXYZ(0, -1, 1, -1)
	box.AddXYZ(0, 1, 1, -1)
	box.AddXYZ(0, 1, -1, -1)

	box.AddXYZ(1, -1, 1, -1)
	box.AddXYZ(1, -1, 1, 1)
	box.AddXYZ(1, 1, 1, 1)
	box.AddXYZ(1, 1, 1, -1)

	box.AddXYZ(2, -1, 1, 1)
	box.AddXYZ(2, -1, -1, 1)
	box.AddXYZ(2, 1, -1, 1)
	box.AddXYZ(2, 1, 1, 1)

	box.AddXYZ(3, -1, -1, 1)
	box.AddXYZ(3, -1, -1, -1)
	box.AddXYZ(3, 1, -1, -1)
	box.AddXYZ(3, 1, -1, 1)
	box.Scale(w, h, d)
	return box
}

// Cone creates a cone shape.
func Cone(height, radius0, radius1 float64, slices, res int) PathList {
	cyl := NewPathList(slices)
	for i := 0; i < slices; i++ {
		radius := blmath.Map(float64(i), 0, float64(slices-1), radius0, radius1)
		for j := 0; j < res; j++ {
			t := blmath.Tau * float64(j) / float64(res)
			y := float64(i)/(float64(slices)-1)*height - height/2
			cyl.AddXYZ(i, math.Cos(t)*radius, y, math.Sin(t)*radius)
		}
	}
	return cyl
}

// Cylinder creates a cylindar shape.
func Cylinder(height, radius float64, slices, res int) PathList {
	return Cone(height, radius, radius, slices, res)
}

// Torus creates a 3d torus.
func Torus(r1, r2 float64, slices, res int) PathList {
	torus := NewPathList(slices)
	fslice := float64(slices)
	dt := blmath.Tau / float64(res)
	for i := 0.0; i < fslice; i++ {
		angle := i / fslice * blmath.Tau
		path := NewPointList()
		for t := 0.0; t <= blmath.Tau-dt; t += dt {
			path.AddXYZ(r1+math.Cos(t)*r2, math.Sin(t)*r2, 0)
		}
		path.RotateY(angle)
		torus.Add(path)
	}
	return torus
}

// Sphere creates a 3d sphere.
func Sphere(radius float64, slices, res int) PathList {
	s := NewPathList(0)
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
		s.Add(path)
	}
	s.UniScale(radius)
	return s
}

func Pyramid(height, baseRadius float64, sides int) PathList {
	p := NewPathList(0)
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
		p.Add(side)
	}
	return p
}

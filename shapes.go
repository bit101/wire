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
	box := NewPathList(6)

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

// Cylinder creates a cylindar shape.
func Cylinder(height, radius float64, slices, res int) PathList {
	cyl := NewPathList(slices)
	dt := blmath.Tau / float64(res)
	for i := 0; i < slices; i++ {
		for t := 0.0; t <= blmath.Tau-dt; t += dt {
			y := float64(i)/float64(slices)*height - height/2
			cyl.AddXYZ(i, math.Cos(t)*radius, y, math.Sin(t)*radius)
		}
	}
	return cyl
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

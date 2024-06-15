// Package wire implements wireframe 3d shapes.
package wire

import (
	"math"

	"github.com/bit101/bitlib/blmath"
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

// GridBox creates a 3d box shape where each surface is a grid.
// If inner is true, it will create a full lattice.
func GridBox(w, h, d float64, xCount, yCount, zCount int, inner bool) *Shape {
	shape := NewShape()
	fx, fy, fz := float64(xCount), float64(yCount), float64(zCount)

	// points
	for z := 0.0; z <= fz; z++ {
		for y := 0.0; y <= fy; y++ {
			for x := 0.0; x <= fx; x++ {
				if inner ||
					x < 1 || x >= fx ||
					y < 1 || y >= fy ||
					z < 1 || z >= fz {
					shape.AddXYZ(x, y, z)
				}
			}
		}
	}

	// segments - it's O(N^2), but deliciously simple, works for inner and outer.
	for i := 0; i < len(shape.Points)-1; i++ {
		for j := i + 1; j < len(shape.Points); j++ {
			a := shape.Points[i]
			b := shape.Points[j]
			// adjacent, connected points will be exactly 1 unit apart.
			// non-adjacent will be at least Sqrt2 apart (diagonal on the same plane)
			if a.Distance(b) < math.Sqrt2 {
				shape.AddSegmentByPoints(a, b)
			}
		}
	}
	shape.Scale(w/fx, h/fy, d/fz)
	shape.Translate(-w/2, -h/2, -d/2)

	return shape
}

// GridPlane creates a 3d plane containing a grid.
func GridPlane(w, d float64, rows, cols int) *Shape {
	shape := NewShape()
	fx, fz := float64(rows), float64(cols)

	// points
	for z := 0.0; z <= fz; z++ {
		for x := 0.0; x <= fx; x++ {
			shape.AddXYZ(x, 0, z)
		}
	}

	// segments - it's O(N^2), but deliciously simple, works for inner and outer.
	for i := 0; i < len(shape.Points)-1; i++ {
		for j := i + 1; j < len(shape.Points); j++ {
			a := shape.Points[i]
			b := shape.Points[j]
			// adjacent, connected points will be exactly 1 unit apart.
			// non-adjacent will be at least Sqrt2 apart (diagonal on the same plane)
			if a.Distance(b) < math.Sqrt2 {
				shape.AddSegmentByPoints(a, b)
			}
		}
	}
	shape.Scale(w/fx, 1, d/fz)
	shape.Translate(-w/2, 0, -d/2)

	return shape
}

// Pyramid creates a 3d pyramid shape.
func Pyramid(height, baseRadius float64, sides int) *Shape {
	return Cone(height, 0, baseRadius, 2, sides, true, true)
}

// RandomInnerBox creates a 3d box filled with random points.
func RandomInnerBox(w, h, d float64, count int) *Shape {
	shape := NewShape()
	for range count {
		shape.AddPoint(RandomPointInBox(w, h, d))
	}
	return shape
}

// RandomSurfaceBox creates a 3d box made of random points on the surface of the box.
func RandomSurfaceBox(w, h, d float64, count int) *Shape {
	shape := NewShape()
	for range count {
		shape.AddPoint(RandomPointOnBox(w, h, d))
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

// RandomInnerCylinder creates a 3d cylinder made of random point inside the cylinder.
func RandomInnerCylinder(height, radius float64, count int) *Shape {
	shape := NewShape()
	for range count {
		shape.AddRandomPointInCylinder(height, radius)
	}
	return shape
}

// RandomSurfaceCylinder creates a 3d cylinder made of random point  on the surface of the cylinder.
func RandomSurfaceCylinder(height, radius float64, count int, includeCaps bool) *Shape {
	shape := NewShape()
	for range count {
		shape.AddRandomPointOnCylinder(height, radius, includeCaps)
	}
	return shape
}

// RandomInnerTorus creates a 3d torus made of random point inside the torus.
func RandomInnerTorus(radius1, radius2, arc float64, count int) *Shape {
	shape := NewShape()
	for range count {
		shape.AddRandomPointInTorus(radius1, radius2, arc)
	}
	return shape
}

// RandomSurfaceTorus creates a 3d torus made of random point  on the surface of the torus.
func RandomSurfaceTorus(radius1, radius2, arc float64, count int) *Shape {
	shape := NewShape()
	for range count {
		shape.AddRandomPointOnTorus(radius1, radius2, arc)
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

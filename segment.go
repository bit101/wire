// Package wire implements wireframe 3d shapes.
package wire

// Segment represents a line segment between two points.
type Segment struct {
	PointA, PointB *Point
}

// NewSegment creates a new segment from two points.
func NewSegment(a, b *Point) *Segment {
	return &Segment{a, b}
}

// Stroke draws a line between the two points of this segment.
func (s *Segment) Stroke(width float64) {
	world.Context.Save()
	scale := (s.PointA.Scaling + s.PointB.Scaling) / 2
	if s.PointA.Visible() && s.PointB.Visible() {
		ApplyFogAndWaterLevel((s.PointA.Y+s.PointB.Y)/2, (s.PointA.Z+s.PointB.Z)/2)
		world.Context.SetLineWidth(width * scale)
		world.Context.MoveTo(s.PointA.Px, s.PointA.Py)
		world.Context.LineTo(s.PointB.Px, s.PointB.Py)
		world.Context.Stroke()
	}
	world.Context.Restore()
}

// Length returns the length of this segment.
func (s *Segment) Length() float64 {
	return s.PointA.Distance(s.PointB)
}

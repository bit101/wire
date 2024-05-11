// Package wire implements wireframe 3d shapes.
package wire

// Segment represents a line segment between two points.
type Segment struct {
	PointA, PointB *Point
}

// NewSegment created a new segment from two points.
func NewSegment(a, b *Point) *Segment {
	return &Segment{a, b}
}

// Clone returns a deep clone of this segment.
func (s *Segment) Clone() *Segment {
	return NewSegment(s.PointA.Clone(), s.PointB.Clone())
}

// Stroke draws a line between the two points of this segment.
func (s *Segment) Stroke(context Context) {
	lineWidth := context.GetLineWidth()
	scale := 1.0
	if World.ScaleLineWidth {
		scale = (s.PointA.Scaling + s.PointB.Scaling) / 2
	}
	if shouldDraw(s.PointA, s.PointB) {
		context.SetLineWidth(lineWidth * scale)
		context.MoveTo(s.PointA.Px, s.PointA.Py)
		context.LineTo(s.PointB.Px, s.PointB.Py)
		context.Stroke()
	}
	context.SetLineWidth(lineWidth)
}

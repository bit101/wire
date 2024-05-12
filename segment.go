// Package wire implements wireframe 3d shapes.
package wire

import (
	"github.com/bit101/bitlib/blcolor"
)

// Segment represents a line segment between two points.
type Segment struct {
	PointA, PointB *Point
}

// NewSegment creates a new segment from two points.
func NewSegment(a, b *Point) *Segment {
	return &Segment{a, b}
}

// Stroke draws a line between the two points of this segment.
func (s *Segment) Stroke() {
	r, g, b := World.R, World.G, World.B
	lineWidth := World.Context.GetLineWidth()
	scale := (s.PointA.Scaling + s.PointB.Scaling) / 2
	if s.PointA.Visible() && s.PointB.Visible() {
		fog := FogAmount((s.PointA.Z + s.PointB.Z) / 2)
		World.Context.SetSourceColor(blcolor.RGBA(r, g, b, fog))
		World.Context.SetLineWidth(lineWidth * scale)
		World.Context.MoveTo(s.PointA.Px, s.PointA.Py)
		World.Context.LineTo(s.PointB.Px, s.PointB.Py)
		World.Context.Stroke()
	}
	World.Context.SetLineWidth(lineWidth)
}

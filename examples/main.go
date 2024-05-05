// Package main renders an image, gif or video
package main

import (
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire"
)

//revive:disable:unused-parameter

func main() {
	renderTarget := target.Video

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/out.png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

var (
	sphere *wire.Shape
)

func init() {
	sphere = wire.Sphere(250, 100, 100)
	sphere.Randomize(10)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	context.SetLineJoin(cairo.LineJoinRound)
	wire.World.CX = width / 2
	wire.World.CY = height / 2
	wire.World.CZ = blmath.LoopSin(percent, -100, 1000)

	s := sphere.Rotated(-percent*blmath.Tau, -percent*blmath.Tau*2, 0)

	context.SetLineWidth(1)
	s.Stroke(context)

	context.GaussianBlur(20)
	context.SetLineWidth(0.5)
	s.Stroke(context)
}

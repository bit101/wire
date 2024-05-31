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
		program.AddSceneWithFrames(scene2, 360)
		program.RenderAndPlayVideo("out/frames", "out/out.mp4")
	}
}

var (
	sphere *wire.Shape
)

func init() {
	sphere = wire.Sphere(250, 100, 100, false, true)
	sphere.Randomize(10)
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	context.SetLineJoin(cairo.LineJoinRound)
	wire.InitWorld(context, width/2, height/2, blmath.LoopSin(percent, -100, 1000))

	s := sphere.Rotated(-percent*blmath.Tau, -percent*blmath.Tau*2, 0)

	s.Stroke(1)

	context.GaussianBlur(20)
	s.Stroke(0.5)
}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	context.SetLineJoin(cairo.LineJoinRound)
	wire.InitWorld(context, width/2, height/2, 800)

	t := wire.NewString("foobarbaz", wire.FontAsteroid).AsCylinder(-400, 20)
	t.RotateY(percent * blmath.Tau * 1)
	t2 := wire.NewString("foobarbaz", wire.FontAsteroid).AsVCylinder(-400, 20)
	t2.RotateX(percent * blmath.Tau * 1)

	t.Stroke(1)
	t2.Stroke(1)
	// context.GaussianBlur(20)
	// t.Stroke(0.5)
	// t2.Stroke(0.5)
}

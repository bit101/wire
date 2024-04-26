// Package main renders an image, gif or video
package main

import (
	"github.com/bit101/bitlib/bllog"
	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/wire/wire"
)

func main() {
	bllog.InitProjectLog("project.log")
	defer bllog.CloseLog()

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
	box    wire.PathList
	cyl    wire.PathList
	torus  wire.PathList
	sphere wire.PathList
)

func init() {
	box = wire.Box()
	box.UniScale(160)
	cyl = wire.Cylinder(20, 24)
	cyl.Scale(50, 200, 50)
	cyl.Randomize(5)

	torus = wire.Torus(200, 100, 60, 32)
	torus.Randomize(5)

	sphere = wire.Sphere(100, 100)
	sphere.UniScale(200)
	sphere.Randomize(20)
}

//revive:disable-next-line:unused-parameter
func scene1(context *cairo.Context, width, height, percent float64) {
	context.BlackOnWhite()
	context.SetLineJoin(cairo.LineJoinRound)
	wire.World.CX = width / 2
	wire.World.CY = height / 2
	wire.World.CZ = 500.0

	b := box.Clone()
	b.Rotate(-percent*blmath.Tau, -percent*blmath.Tau*2, 0)

	c := cyl.Clone()
	c.Rotate(-percent*blmath.Tau, -percent*blmath.Tau*2, 0)

	t := torus.Clone()
	t.Rotate(-percent*blmath.Tau, -percent*blmath.Tau*2, 0)
	// c.Randomize(blmath.LoopSin(percent, 0, 5))

	s := sphere.Clone()
	s.Rotate(-percent*blmath.Tau, -percent*blmath.Tau*2, 0)

	context.SetLineWidth(0.5)
	// c.Stroke(context, true)
	// t.Stroke(context, true)
	s.Stroke(context, true)
	b.Stroke(context, true)

	// context.GaussianBlur(10)
	//
	// context.SetLineWidth(0.5)
	// // c.Stroke(context, true)
	// // t.Stroke(context, true)
	// s.Stroke(context, true)
	// b.Stroke(context, true)
	//
	context.GridFull(20, 0.125)
}

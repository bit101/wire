# wire

`wire` is a niche 3d modeling and rendering library written in Go.

## Dependencies

Direct dependencies are two other Go libraries I've been working on for some years:

- https://github.com/bit101/bitlib
- https://github.com/bit101/blcairo (requires https://cairographics.org/ C library).

## Concept

`wire` models are made of 3d points and, optionally, segments, which connect two points.

A `wire` model can be rendered as 3d paths (wires) formed by stroking all the segments, or as a 3d point cloud by rendering all of the points. Or both.

## Constraints

So many constraints!

- No triangles/polygons/filled surfaces,solid objects
- No z-sorting
- No normals
- No shaders
- No backface culling

Due to some of these constraints, monochromatic images and animations tend to work best. A secondary constraint.

That said, these constraints make space for a good amount of creativity in multiple styles.

## Features

- A decent set of 3d primitives:
  - Sphere
  - Circle
  - Cylinder
  - Cone
  - Box
  - Pyramid
  - Torus
  - Torus knot
  - GridPlane
  - Spring
  - Text
  - Platonic Solids:
    - Tetrahedron
    - Cube
    - Octahedron
    - Dodecahedron
    - Icosohedron
- Where it makes sense, most of these can be rendered as:
  - Wireframe, with configurable longitudinal and latitudinal sections
  - Random points on the surface
  - Random points filling the object
- Objects can be:
  - Transformed: rotated, translated, scaled, or manually manipulated with custom functions
  - Cloned
  - Combined
  - Point-culled with custom functions
  - Subdivided
  - Randomized
  - Noisified (Simplex noise)

## Usage

```
// expected: you have a *cairo.Context from blcairo
// 1. Init wire with the context and origin location:
wire.InitWorld(context, centerX, centerY, centerZ)

// 2. Set other (optional) world settings:
wire.SetPerspective(500)
wire.SetFog(true, 500, 1000)

// 3. Create a shape:
model := wire.Sphere(200, 12, 12, true, true)

// 4. Transform model as desired:
model.ScaleX(2.0)
model.RotateY(math.Pi/2)

// 5. Render points and/or edges:
model.Stroke(1)
model.RenderPoints(2)
```

## Examples

See: [https://www.artfromcode.com/tags/wire/](https://www.artfromcode.com/tags/wire/) for some examples of what is possible.

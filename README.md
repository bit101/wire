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

## Examples

See: [https://www.artfromcode.com/tags/wire/](https://www.artfromcode.com/tags/wire/) for some examples of what is possible.

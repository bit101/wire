// Package wire implements wireframe 3d shapes.
package wire

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// ShapeFromXYZ creates a new point-only shape from an .xyz formatted point cloud file.
func ShapeFromXYZ(fileName string) *Shape {
	model := NewShape()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("could not open model:", err)
	}
	defer file.Close()
	minX, minY, minZ := math.MaxFloat64, math.MaxFloat64, math.MaxFloat64
	maxX, maxY, maxZ := -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		x, err := strconv.ParseFloat(vals[0], 64)
		if err != nil {
			log.Fatal("unable to parse float:", err)
		}
		y, err := strconv.ParseFloat(vals[1], 64)
		if err != nil {
			log.Fatal("unable to parse float:", err)
		}
		z, err := strconv.ParseFloat(vals[2], 64)
		if err != nil {
			log.Fatal("unable to parse float:", err)
		}
		minX = math.Min(minX, x)
		minY = math.Min(minY, y)
		minZ = math.Min(minZ, z)
		maxX = math.Max(maxX, x)
		maxY = math.Max(maxY, y)
		maxZ = math.Max(maxZ, z)
		model.AddXYZ(x, y, z)
	}
	model.Translate(
		-minX-(maxX-minX)/2,
		-minY-(maxY-minY)/2,
		-minZ-(maxZ-minZ)/2,
	)
	model.Rotate(-math.Pi/2, math.Pi, 0)
	return model
}

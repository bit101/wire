// Package wire implements wireframe 3d shapes.
package wire

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

//////////////////////////////////////////////////////////////////////////////////////
// XYZ files are originally for chemical descriptions.
//
// First line optional vertex count.
// Second line optional camment.
// All other lines have vertex data:
//
// <elem> x y z
//
// Element is optional. x, y, z are floats in any format.
//
// Valid lines:
//
// C 0 1 2
// Fe -0.432 0.457 10
// H     1.23e23    2.34E-12    -0.23e-456
// 34.765   45.987  -98.123
//
// This parser ignores the vertex count and comment if they exist,
// and igores the element name if it exists. This works on all files I've tried from
// various sources so far. Will update if I find any breaking ones.
//////////////////////////////////////////////////////////////////////////////////////

var (
	floatExp = "([-+]?[0-9]*\\.?[0-9]+(?:[eE][-+]?[0-9]+)?)"
	exp      = fmt.Sprintf("(?:[a-zA-Z]* +)?%s +%s +%s", floatExp, floatExp, floatExp)
)

// ShapeFromXYZ creates a new point-only shape from an .xyz formatted point cloud file.
func ShapeFromXYZ(fileName string) *Shape {
	pattern, err := regexp.Compile(exp)
	if err != nil {
		fmt.Println(err)
	}

	// open file
	model := NewShape()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("could not open model:", err)
	}
	defer file.Close()

	// read lines
	lineNum := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// do regex magic
		matches := pattern.MatchString(line)
		if matches {
			match := pattern.FindStringSubmatch(line)
			// match[0] is entire match. ignore.
			x := getFloat(match[1], lineNum)
			y := getFloat(match[2], lineNum)
			z := getFloat(match[3], lineNum)
			model.AddXYZ(x, y, z)
		} else if lineNum > 2 {
			// per xyz spec fisrt two lines are optionally:
			// 1. number of vertices
			// 2. comment/space
			log.Fatalf("couldn't parse line %d: %q", lineNum, line)
		}
		lineNum++
	}

	// adjust to wire's coord system
	model.Center()
	model.Rotate(-math.Pi/2, math.Pi, 0)
	return model
}

func getFloat(s string, lineNum int) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("couldn't parse float on line %d: %q", lineNum, s)
	}
	return val

}

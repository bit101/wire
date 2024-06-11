// Package wire implements wireframe 3d shapes.
package wire

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////////////
// Shapes can be saved and loaded. The file format is:
// <number of points>
// x y z
// x y z
// x y z
// ...
// <number of segments>
// indexA indexB
// indexA indexB
// indexA indexB
// ...
//
// Simple example (a triangle):
// 3
// 0 -10 0
// 10 10 0
// -10 10 0
// 3
// 0 1
// 1 2
// 2 3
//////////////////////////////////////////////////////////////

func (s *Shape) Save(fileName string) {
	file, err := os.Create(fileName)
	checkErr(err)
	defer file.Close()

	// write points
	_, err = file.WriteString(strconv.Itoa(len(s.Points)) + "\n")
	checkErr(err)

	for _, p := range s.Points {
		str := fmt.Sprintf("%f %f %f\n", p.X, p.Y, p.Z)
		_, err = file.WriteString(str)
		checkErr(err)
	}

	// write segments
	_, err = file.WriteString(strconv.Itoa(len(s.Segments)) + "\n")
	checkErr(err)

	for _, seg := range s.Segments {
		i := slices.Index(s.Points, seg.PointA)
		j := slices.Index(s.Points, seg.PointB)
		str := fmt.Sprintf("%d %d\n", i, j)
		_, err = file.WriteString(str)
		checkErr(err)
	}
}

func LoadShape(fileName string) (*Shape, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("unable to load shape: " + err.Error())
	}
	defer file.Close()

	shape := NewShape()
	scanner := bufio.NewScanner(file)

	// parse points
	scanner.Scan()
	numPoints, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		return nil, errors.New("unable to parse shape: " + err.Error())
	}
	for range numPoints {
		scanner.Scan()
		line := scanner.Text()
		x, y, z, err := parseCoords(line)
		if err != nil {
			return nil, errors.New("unable to parse shape: " + err.Error())
		}
		shape.AddXYZ(x, y, z)
	}

	// parse segments
	scanner.Scan()
	numSegments, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		return nil, errors.New("unable to parse shape: " + err.Error())
	}
	for range numSegments {
		scanner.Scan()
		line := scanner.Text()
		i, j, err := parseIndices(line)
		if err != nil {
			return nil, errors.New("unable to parse shape: " + err.Error())
		}
		if i < 0 || i >= len(shape.Points) || j < 0 || j >= len(shape.Points) {
			return nil, errors.New("invalid segment index, should be from zero to length of points minus one")
		}
		shape.AddSegmentByIndex(i, j)
	}
	return shape, nil
}

func parseCoords(line string) (float64, float64, float64, error) {
	coords := strings.Split(line, " ")
	x, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return 0, 0, 0, err
	}
	y, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return 0, 0, 0, err
	}
	z, err := strconv.ParseFloat(coords[2], 64)
	if err != nil {
		return 0, 0, 0, err
	}
	return x, y, z, nil
}

func parseIndices(line string) (int, int, error) {
	indices := strings.Split(line, " ")
	i, err := strconv.ParseInt(indices[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	j, err := strconv.ParseInt(indices[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return int(i), int(j), nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

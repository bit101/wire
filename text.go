// Package wire implements wireframe 3d shapes.
package wire

import (
	"math"
	"strconv"
	"strings"
)

const fontHalfWidth = 50.0

//////////////////////////////
// Font data
//////////////////////////////

// FontType defines which font will be used.
type FontType int

const (
	// FontAsteroid uses the Asteroid Font.
	FontAsteroid = iota
	// FontArcade uses the Arcade Font.
	FontArcade
	// maybe more to come...
)

//////////////////////////////
// String type
//////////////////////////////

// String represents a 3d character string.
// Initially this only holds a list of shapes, each on representing one letter in the string.
// Calling one of the `As...` methods returns a single unified shape object.
type String struct {
	Orig    string
	Letters []*Shape
}

// NewString creates a new 3d string object.
func NewString(str string, font FontType) *String {
	str = strings.ToUpper(str)
	paths := []*Shape{}
	for _, s := range str {
		char := ParseChar(string(s), font)
		paths = append(paths, char)
	}
	return &String{str, paths}
}

// ParseChar parses a single character into a single 3d shape.
// The font data is initially sized from -1 to +1 on the x-axis.
// Height will depend on the font.
// Each character is scaled to 100 units wide on creation (-50 to +50).
// The string shape can be scaled further later.
func ParseChar(char string, font FontType) *Shape {
	fontData := arcadeFont
	if font == FontAsteroid {
		fontData = asteroidFont
	}
	charData := fontData[char]
	strokes := strings.Split(charData, ":")
	shape := NewShape()
	index := 0
	for _, stroke := range strokes {
		stroke = strings.TrimSpace(stroke)
		coords := strings.Split(stroke, " ")
		for i, coord := range coords {
			xi, _ := strconv.ParseInt(string(coord[0]), 16, 64)
			yi, _ := strconv.ParseInt(string(coord[1]), 16, 64)
			x := float64(xi)/4.0 - 1.0
			y := 1.0 - float64(yi)/4.0
			shape.AddXYZ(x, y, 0)
			if i > 0 {
				shape.AddSegmentByIndex(index-1, index)
			}
			index++
		}
	}
	shape.UniScale(fontHalfWidth)
	return shape
}

// AsCylinder creates a single shape consisting of the all the chars in the string wrapped around a cylinder.
func (s *String) AsCylinder(radius, spacing float64) *Shape {
	shape := NewShape()
	for _, pl := range s.Letters {
		pl.TranslateZ(-radius)
		shape.Points = append(shape.Points, pl.Points...)
		shape.Segments = append(shape.Segments, pl.Segments...)
		shape.RotateY(math.Atan2(fontHalfWidth+spacing, radius) * 2)
	}
	return shape
}

// AsLine creates a single shape consisting of all the chars in the string laid out in a signle line.
func (s *String) AsLine(spacing float64) *Shape {
	shape := NewShape()
	for i, pl := range s.Letters {
		pl.TranslateX(50 + (100+spacing)*float64(i))
		shape.Points = append(shape.Points, pl.Points...)
		shape.Segments = append(shape.Segments, pl.Segments...)
	}
	mult := float64(len(s.Letters))
	shape.TranslateX(-(fontHalfWidth*2+spacing)*mult/2 + spacing)
	return shape
}

//////////////////////////////
// Font definitions
//////////////////////////////

// arcadeFont is the path data for this font.
// adapted from https://github.com/coolbutuseless/arcadefont/blob/master/data-raw/create-arcade-font.R
// with some changes.
var arcadeFont = map[string]string{
	" ":  "00",
	"!":  "00 01 11 10 00 :  03 08 18 13 03",
	"#":  "03 83 :  05 85 :  30 38 :  50 58",
	"$":  "01 81 84 04 07 87 :  40 48",
	"%":  "00 88 :  18 28 27 17 18 :  70 71 61 60 70",
	"&":  "80 47 58 67 21 30 60 82",
	"'":  "07 08 18 17 07 :  17 05",
	"(":  "28 04 20",
	")":  "08 24 00",
	"*":  "41 47 :  14 74 :  22 66 :  26 62",
	"+":  "41 47 :  14 74",
	",":  "01 02 12 11 01 :  11 00",
	"-":  "14 74",
	".":  "00 01 11 10 00",
	"/":  "00 88",
	"0":  "00 80 88 08 00",
	"1":  "00 80 :  40 48 26",
	"2":  "08 88 84 04 00 80",
	"3":  "08 88 80 00 :  04 84",
	"4":  "08 04 84 :  88 80",
	"5":  "00 80 84 04 08 88",
	"6":  "08 00 80 84 04",
	"7":  "08 88 80",
	"8":  "00 08 88 80 00 :  04 84",
	"9":  "80 88 08 04 84",
	":":  "02 03 13 12 02 :  05 06 16 15 05",
	";":  "02 03 13 12 02 :  05 06 16 15 05 :  1 2 0 0",
	"<":  "87 04 81",
	"=":  "13 73 :  15 75",
	">":  "07 84 01",
	"?":  "06 08 88 84 44 40",
	"@":  "71 60 20 02 06 28 68 86 84 62 22 24 36 66 62",
	"A":  "00 06 48 86 80 :  03 83",
	"B":  "00 08 58 76 54 04 :  64 82 60 00",
	"C":  "88 08 00 80",
	"D":  "00 08 58 85 83 50 00",
	"E":  "88 08 00 80 :  04 64",
	"F":  "88 08 00 :  04 64",
	"G":  "88 08 00 80 83 43",
	"H":  "00 08 :  80 88 :  04 84",
	"I":  "00 80 :  08 88 :  40 48",
	"J":  "88 80 40 03",
	"K":  "00 08 :  88 04 80",
	"L":  "08 00 80",
	"M":  "00 08 45 88 80",
	"N":  "00 08 80 88",
	"O":  "00 80 88 08 00",
	"P":  "00 08 88 84 04",
	"Q":  "00 08 88 83 40 00 :  43 80",
	"R":  "00 08 88 84 04 80",
	"S":  "00 80 84 04 08 88",
	"T":  "08 88 :  40 48",
	"U":  "08 00 80 88",
	"V":  "08 40 88",
	"W":  "08 00 43 80 88",
	"X":  "00 88 :  08 80",
	"Y":  "08 45 88 :  45 40",
	"Z":  "08 88 00 80",
	"[":  "28 08 00 20",
	"\"": "07 08 18 17 07 :  17 05 :  27 28 38 37 27 :  37 25",
	"\\": "08 80",
	"]":  "08 28 20 00",
	"^":  "26 48 66",
	"_":  "00 80",
	"`":  "18 26",
	"{":  "28 04 20",
	"|":  "40 48",
	"}":  "08 24 00",
	"~":  "04 26 64 86",
}

// asteroidFont is the path data for this font.
// adapted from https://github.com/osresearch/vst/blob/master/teensyv/asteroids_font.c
// with some changes.
var asteroidFont = map[string]string{
	" ":  "00",
	"!":  "40 32 52 40 : 44 4C",
	"#":  "04 84 62 6A 88 08 2A 22",
	"$":  "62 26 6A : 4C 40",
	"%":  "00 8C : 2A 28 : 64 62",
	"&":  "80 4C 88 04 40 84",
	"'":  "26 6A",
	"(":  "60 24 28 6C",
	")":  "20 64 68 2C",
	"*":  "00 4C 80 08 88 00",
	"+":  "16 76 : 49 43",
	",":  "30 42",
	"-":  "26 66",
	".":  "30 40",
	"/":  "00 8C",
	"0":  "00 80 8C 0C 00 8C",
	"1":  "40 4C 3A",
	"2":  "0C 8C 87 05 00 80",
	"3":  "0C 8C 80 00 : 06 86",
	"4":  "0C 06 86 : 8C 80",
	"5":  "00 80 86 07 0C 8C",
	"6":  "0C 00 80 85 07",
	"7":  "0C 8C 86 40",
	"8":  "00 80 8C 0C 00 : 06 86",
	"9":  "80 8C 0C 07 85",
	":":  "49 47 : 45 43",
	";":  "49 47 : 45 12",
	"<":  "60 26 6C",
	"=":  "14 74 : 18 78",
	">":  "20 66 2C",
	"?":  "08 4C 88 44 : 41 40",
	"@":  "84 40 04 08 4C 88 44 36",
	"A":  "00 08 4C 88 80 : 04 84",
	"B":  "00 0C 4C 8A 46 82 40 00",
	"C":  "80 00 0C 8C",
	"D":  "00 0C 4C 88 84 40 00",
	"E":  "80 00 0C 8C : 06 66",
	"F":  "00 0C 8C : 06 66",
	"G":  "66 84 80 00 0C 8C",
	"H":  "00 0C : 06 86 : 8C 80",
	"I":  "00 80 : 40 4C : 0C 8C",
	"J":  "04 40 80 8C",
	"K":  "00 0C : 8C 06 60",
	"L":  "80 00 0C",
	"M":  "00 0C 48 8C 80",
	"N":  "00 0C 80 8C",
	"O":  "00 0C 8C 80 00",
	"P":  "00 0C 8C 86 05",
	"Q":  "00 0C 8C 84 00 : 44 80",
	"R":  "00 0C 8C 86 05 : 45 80",
	"S":  "02 20 80 85 07 0C 6C 8A",
	"T":  "0C 8C : 4C 40",
	"U":  "0C 02 40 82 8C",
	"V":  "0C 40 8C",
	"W":  "0C 20 44 60 8C",
	"X":  "00 8C : 0C 80",
	"Y":  "0C 46 8C : 46 40",
	"Z":  "0C 8C 00 80 : 26 66",
	"[":  "60 20 2C 6C",
	"\"": "2A 26 : 6A 66",
	"\\": "0C 80",
	"]":  "20 60 6C 2C",
	"^":  "26 4C 66",
	"_":  "00 80",
	"`":  "2A 66",
	"{":  "60 42 4A 6C : 26 46",
	"|":  "40 45 : 46 4C",
	"}":  "40 62 6A 4C : 66 86",
	"~":  "04 28 64 88",
}

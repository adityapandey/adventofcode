package util

import (
	"image"
)

var Neighbors4 = []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var Neighbors8 = []image.Point{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func Manhattan(p, q image.Point) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
}

func Bounds(p []image.Point) image.Rectangle {
	var r image.Rectangle
	for _, pp := range p {
		r = r.Union(image.Rectangle{pp, pp.Add(image.Pt(1, 1))})
	}
	return r.Bounds()
}

type Pt3 struct {
	X, Y, Z int
}

func (p1 Pt3) Add(p2 Pt3) Pt3 {
	return Pt3{p1.X + p2.X, p1.Y + p2.Y, p1.Z + p2.Z}
}
func (p Pt3) Mul(k int) Pt3 {
	return Pt3{p.X * k, p.Y * k, p.Z * k}
}

func Manhattan3(p1, p2 Pt3) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y) + Abs(p1.Z-p2.Z)
}

type Dir int

const (
	N Dir = iota
	E
	S
	W
)

func (d Dir) Add(n int) Dir {
	return Dir((int(d) + n) % 4)
}

func (d Dir) Next() Dir {
	return (d + 1) % 4
}

func (d Dir) Prev() Dir {
	return (d + 3) % 4
}

func (d Dir) Reverse() Dir {
	return (d + 2) % 4
}

var point = map[Dir]image.Point{N: {0, 1}, E: {1, 0}, S: {0, -1}, W: {-1, 0}}
var pointReversed = map[Dir]image.Point{N: {0, -1}, E: {1, 0}, S: {0, 1}, W: {-1, 0}}

// Y-axis goes up.
func (d Dir) Point() image.Point {
	return point[d]
}

// Y-axis goes down.
func (d Dir) PointR() image.Point {
	return pointReversed[d]
}

var fromPoint = map[image.Point]Dir{{0, 1}: N, {1, 0}: E, {0, -1}: S, {-1, 0}: W}
var fromPointReversed = map[image.Point]Dir{{0, -1}: N, {1, 0}: E, {0, 1}: S, {-1, 0}: W}

// Y-axis goes up.
func DirFromPoint(p image.Point) Dir {
	return fromPoint[p]
}

// Y-axis goes down.
func DirFromPointR(p image.Point) Dir {
	return fromPointReversed[p]
}

var fromByte = map[byte]Dir{
	'N': N,
	'E': E,
	'S': S,
	'W': W,
	'U': N,
	'R': E,
	'D': S,
	'L': W,
	'^': N,
	'>': E,
	'v': S,
	'<': W,
}

func DirFromByte(b byte) Dir {
	return fromByte[b]
}

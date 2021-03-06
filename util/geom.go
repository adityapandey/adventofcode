package util

import "image"

var Neighbors4 = []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var Neighbors8 = []image.Point{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func Manhattan(p, q image.Point) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
}

type Pt3 struct {
	X, Y, Z int
}

func (p1 Pt3) Add(p2 Pt3) Pt3 {
	return Pt3{p1.X + p2.X, p1.Y + p2.Y, p1.Z + p2.Z}
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

func (d Dir) Point() image.Point {
	return point[d]
}

func (d Dir) PointR() image.Point {
	return pointReversed[d]
}

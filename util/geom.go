package util

import "image"

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

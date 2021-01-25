package main

import (
	"fmt"
	"image"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

var dirmap = map[int]image.Point{
	0: {1, 0},
	1: {0, 1},
	2: {-1, 0},
	3: {0, -1},
}

var neighbors = []image.Point{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
}

func main() {
	var steps int
	fmt.Fscanf(os.Stdin, "%d", &steps)
	dir := 0
	radius := 0
	var p image.Point
	m := map[image.Point]int{{0, 0}: 1}
	found, foundVal := false, 0
	for i := 2; i <= steps; i++ {
		nxt := p.Add(dirmap[dir])
		switch {
		case nxt.X > radius:
			radius++
			dir = (dir + 1) % 4
			p = nxt
		case nxt.Y > radius:
			fallthrough
		case nxt.X < -radius:
			fallthrough
		case nxt.Y < -radius:
			dir = (dir + 1) % 4
			p = p.Add(dirmap[dir])
		default:
			p = nxt
		}

		if !found {
			for _, n := range neighbors {
				if v, ok := m[p.Add(n)]; ok {
					m[p] += v
				}
			}
			if m[p] > steps {
				found, foundVal = true, m[p]
			}
		}
	}
	fmt.Println(util.Manhattan(p, image.Pt(0, 0)))
	fmt.Println(foundVal)
}

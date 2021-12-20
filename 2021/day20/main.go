package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var square = []image.Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {0, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

type grid struct {
	cells        map[image.Point]byte
	defaultValue byte
}

func newGrid() grid {
	return grid{map[image.Point]byte{}, '.'}
}

func (g grid) at(p image.Point) byte {
	if b, ok := g.cells[p]; ok {
		return b
	}
	return g.defaultValue
}

func (g *grid) enhance(enhancement string) {
	next := map[image.Point]byte{}
	var ps []image.Point
	for p := range g.cells {
		ps = append(ps, p)
	}
	b := util.Bounds(ps).Inset(-1)
	for x := b.Min.X; x <= b.Max.X; x++ {
		for y := b.Min.Y; y <= b.Max.Y; y++ {
			var sum int
			for _, n := range square {
				sum <<= 1
				if g.at(image.Pt(x, y).Add(n)) == '#' {
					sum++
				}
			}
			next[image.Pt(x, y)] = enhancement[sum]
		}
	}
	g.cells = next
	if g.defaultValue == '.' {
		g.defaultValue = enhancement[0]
	} else {
		g.defaultValue = enhancement[511]
	}
}

func (g grid) litCount() int {
	var sum int
	for _, b := range g.cells {
		if b == '#' {
			sum++
		}
	}
	return sum
}

func main() {
	sp := strings.Split(util.ReadAll(), "\n\n")
	enhancement := sp[0]
	g := newGrid()
	for y, line := range strings.Split(sp[1], "\n") {
		for x := 0; x < len(line); x++ {
			g.cells[image.Pt(x, y)] = line[x]
		}
	}

	for i := 1; i <= 50; i++ {
		g.enhance(enhancement)
		if i == 2 || i == 50 {
			fmt.Println(g.litCount())
		}
	}
}

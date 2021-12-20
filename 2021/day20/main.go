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
	m            map[int]byte
}

func newGrid(m map[int]byte) grid {
	return grid{map[image.Point]byte{}, '.', m}
}

func (g grid) at(p image.Point) byte {
	if b, ok := g.cells[p]; ok {
		return b
	}
	return g.defaultValue
}

func (g *grid) enhance() {
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
			next[image.Pt(x, y)] = g.m[sum]
		}
	}
	g.cells = next
	if g.defaultValue == '.' {
		g.defaultValue = g.m[0]
	} else {
		g.defaultValue = g.m[len(g.m)-1]
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

func (g grid) String() string {
	var sb strings.Builder
	var ps []image.Point
	for p := range g.cells {
		ps = append(ps, p)
	}
	b := util.Bounds(ps)
	for y := b.Min.Y; y <= b.Max.Y; y++ {
		for x := b.Min.X; x <= b.Max.X; x++ {
			fmt.Fprintf(&sb, "%c", g.at(image.Pt(x, y)))
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}

func main() {
	m := map[int]byte{}
	sp := strings.Split(util.ReadAll(), "\n\n")
	for i := 0; i < len(sp[0]); i++ {
		m[i] = sp[0][i]
	}
	g := newGrid(m)
	for y, line := range strings.Split(sp[1], "\n") {
		for x := 0; x < len(line); x++ {
			g.cells[image.Pt(x, y)] = line[x]
		}
	}

	for i := 1; i <= 50; i++ {
		g.enhance()
		if i == 2 || i == 50 {
			fmt.Println(g.litCount())
		}
	}
}

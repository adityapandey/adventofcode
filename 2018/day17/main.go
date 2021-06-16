package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type Pt image.Point

func (p Pt) Below() Pt {
	return Pt(image.Point(p).Add(util.DirFromByte('D').PointR()))
}
func (p Pt) Left() Pt {
	return Pt(image.Point(p).Add(util.DirFromByte('L').PointR()))
}
func (p Pt) Right() Pt {
	return Pt(image.Point(p).Add(util.DirFromByte('R').PointR()))
}

type grid map[Pt]byte

func (g grid) isWall(p Pt) bool {
	return g[p] == '#' || g[p] == '~'
}

func main() {
	g := make(grid)
	s := util.ScanAll()
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "x") {
			var x, y0, y1 int
			fmt.Sscanf(s.Text(), "x=%d, y=%d..%d", &x, &y0, &y1)
			for y := y0; y <= y1; y++ {
				g[Pt{x, y}] = '#'
			}
		} else {
			var x0, x1, y int
			fmt.Sscanf(s.Text(), "y=%d, x=%d..%d", &y, &x0, &x1)
			for x := x0; x <= x1; x++ {
				g[Pt{x, y}] = '#'
			}
		}
	}
	bounds := getBounds(g)
	fill(g, bounds, Pt{500, 0})

	var nFlow, nSettled int
	for p, c := range g {
		if p.Y < bounds.Min.Y || p.Y > bounds.Max.Y {
			continue
		}
		switch c {
		case '|':
			nFlow++
		case '~':
			nSettled++
		}
	}
	// Part 1
	fmt.Println(nSettled + nFlow)
	// Part 2
	fmt.Println(nSettled)
}

func fill(g grid, bounds image.Rectangle, p Pt) {
	if p.Y > bounds.Max.Y || g.isWall(p) {
		return
	}
	if g.isWall(p.Below()) {
		// Go as left and right as we can, so long as the tile below is impermeable.
		l, r := p, p
		for !g.isWall(l) && g.isWall(l.Below()) {
			g[l] = '|'
			l = l.Left()
		}
		for !g.isWall(r) && g.isWall(r.Below()) {
			g[r] = '|'
			r = r.Right()
		}
		// Both edges are walls - so this is all still water.
		if g.isWall(l) && g.isWall(r) {
			for x := l.X + 1; x < r.X; x++ {
				g[Pt{x, p.Y}] = '~'
			}
		}
		// Water can flow off the edges
		if !g.isWall(l.Below()) {
			fill(g, bounds, l)
		}
		if !g.isWall(r.Below()) {
			fill(g, bounds, r)
		}
		return
	}
	if g[p] == '|' {
		return
	}
	// Flow downwards
	g[p] = '|'
	fill(g, bounds, p.Below())
	// If a pool of settled water is formed, reflow
	if g[p.Below()] == '~' {
		fill(g, bounds, p)
	}
}

func getBounds(g grid) image.Rectangle {
	var ps []image.Point
	for p := range g {
		ps = append(ps, image.Point(p))
	}
	return util.Bounds(ps)
}

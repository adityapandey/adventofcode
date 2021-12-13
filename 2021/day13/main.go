package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type grid map[image.Point]struct{}

func (g grid) foldAlong(dim byte, n int) {
	switch dim {
	case 'y':
		for p := range g {
			if p.Y > n {
				g[image.Pt(p.X, 2*n-p.Y)] = struct{}{}
				delete(g, p)
			}
		}
	case 'x':
		for p := range g {
			if p.X > n {
				g[image.Pt(2*n-p.X, p.Y)] = struct{}{}
				delete(g, p)
			}
		}
	}
}

func (g grid) String() string {
	var ps []image.Point
	for p := range g {
		ps = append(ps, p)
	}
	b := util.Bounds(ps)
	var sb strings.Builder
	for y := 0; y <= b.Max.Y; y++ {
		for x := 0; x <= b.Max.X; x++ {
			_, ok := g[image.Pt(x, y)]
			fmt.Fprintf(&sb, "%c", map[bool]byte{true: '#', false: ' '}[ok])
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}

func main() {
	g := grid{}
	sp := strings.Split(util.ReadAll(), "\n\n")
	for _, s := range strings.Split(sp[0], "\n") {
		var p image.Point
		fmt.Sscanf(s, "%d,%d", &p.X, &p.Y)
		g[p] = struct{}{}
	}
	for i, s := range strings.Split(sp[1], "\n") {
		var dim byte
		var n int
		fmt.Sscanf(s, "fold along %c=%d", &dim, &n)
		g.foldAlong(dim, n)
		if i == 0 {
			fmt.Println(len(g))
		}
	}
	fmt.Println(g)
}

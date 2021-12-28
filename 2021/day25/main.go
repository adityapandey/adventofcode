package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type grid struct {
	pts  map[image.Point]byte
	w, h int
}

func parse(s string) grid {
	g := grid{}
	g.pts = make(map[image.Point]byte)
	lines := strings.Split(s, "\n")
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == '>' || line[x] == 'v' {
				g.pts[image.Pt(x, y)] = line[x]
			}
		}
	}
	g.w, g.h = len(lines[0]), len(lines)
	return g
}

func (g grid) step() int {
	tmp := map[image.Point]byte{}
	var moves int
	for p, b := range g.pts {
		switch b {
		case '>':
			next := image.Pt((p.X+1)%g.w, p.Y)
			if _, ok := g.pts[next]; !ok {
				tmp[next] = '>'
				moves++
			} else {
				tmp[p] = '>'
			}
		case 'v':
			tmp[p] = 'v'
		}
	}
	for p := range g.pts {
		delete(g.pts, p)
	}
	for p, b := range tmp {
		switch b {
		case 'v':
			next := image.Pt(p.X, (p.Y+1)%g.h)
			if _, ok := tmp[next]; !ok {
				g.pts[next] = 'v'
				moves++
			} else {
				g.pts[p] = 'v'
			}
		case '>':
			g.pts[p] = '>'
		}
	}
	return moves
}

func (g grid) String() string {
	var sb strings.Builder
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			if b, ok := g.pts[image.Pt(x, y)]; ok {
				fmt.Fprintf(&sb, "%c", b)
			} else {
				fmt.Fprintf(&sb, ".")
			}
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}

func main() {
	g := parse(util.ReadAll())
	var steps int
	for g.step() > 0 {
		steps++
	}
	fmt.Println(steps + 1)
}

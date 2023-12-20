package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

type ray struct {
	p image.Point
	d util.Dir
}

func main() {
	grid := map[image.Point]byte{}
	for y, line := range strings.Split(util.ReadAll(), "\n") {
		for x := range line {
			grid[image.Pt(x, y)] = line[x]
		}
	}
	fmt.Println(energize(ray{image.Pt(-1, 0), util.E}, grid))

	b := util.Bounds(maps.Keys(grid))
	max := 0
	for x := 0; x < b.Max.X; x++ {
		max = util.Max(max, energize(ray{image.Pt(x, 0), util.S}, grid), energize(ray{image.Pt(x, b.Max.Y), util.N}, grid))
	}
	for y := 0; y < b.Max.Y; y++ {
		max = util.Max(max, energize(ray{image.Pt(0, y), util.E}, grid), energize(ray{image.Pt(b.Max.X, y), util.W}, grid))
	}
	fmt.Println(max)
}

func energize(r ray, grid map[image.Point]byte) int {
	out := map[ray]struct{}{}
	beam(r, grid, out, 0)

	visited := map[image.Point]struct{}{}
	for r := range out {
		visited[r.p] = struct{}{}
	}
	return len(visited)

}

func beam(r ray, grid map[image.Point]byte, out map[ray]struct{}, n int) {
	// fmt.Println(strings.Repeat("  ", n), "NEW")
	for {
		r.p = r.p.Add(r.d.PointR())
		b, ok := grid[r.p]
		if !ok {
			// fmt.Println(strings.Repeat("  ", n), "out")
			return
		}
		if _, ok := out[r]; ok {
			// fmt.Println(strings.Repeat("  ", n), "seen")
			return
		}
		out[r] = struct{}{}
		// fmt.Printf("%s%v %v %c\n", strings.Repeat("  ", n), r.p, r.d, b)
		switch b {
		case '/':
			switch r.d {
			case util.N:
				r.d = util.E
			case util.E:
				r.d = util.N
			case util.S:
				r.d = util.W
			case util.W:
				r.d = util.S
			}
		case '\\':
			switch r.d {
			case util.N:
				r.d = util.W
			case util.W:
				r.d = util.N
			case util.S:
				r.d = util.E
			case util.E:
				r.d = util.S
			}
		case '-':
			if r.d == util.N || r.d == util.S {
				beam(ray{r.p, util.E}, grid, out, n+1)
				beam(ray{r.p, util.W}, grid, out, n+1)
				return
			}
		case '|':
			if r.d == util.E || r.d == util.W {
				beam(ray{r.p, util.N}, grid, out, n+1)
				beam(ray{r.p, util.S}, grid, out, n+1)
				return
			}
		}
	}
}

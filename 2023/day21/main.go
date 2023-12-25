package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

func main() {
	grid := map[image.Point]byte{}
	var start image.Point
	for y, line := range strings.Split(util.ReadAll(), "\n") {
		for x := range line {
			grid[image.Pt(x, y)] = line[x]
			if line[x] == 'S' {
				start = image.Pt(x, y)
				grid[image.Pt(x, y)] = '.'
			}
		}
	}
	b := util.Bounds(maps.Keys(grid))
	w, h := b.Dx(), b.Dy()
	if w != h {
		panic("broken assumption: square grid")
	}
	loc := []image.Point{start}

	fmt.Println(stepCount(grid, loc, 64))
	fmt.Println(solveQuadratic(grid, loc, 26501365/w))
}

func solveQuadratic(grid map[image.Point]byte, loc []image.Point, x int) int {
	w := util.Bounds(maps.Keys(grid)).Dx()
	a := stepCount(grid, loc, w/2)
	b := stepCount(grid, loc, 3*w/2)
	c := stepCount(grid, loc, 5*w/2)
	return a + (b-a)*x + (c-2*b+a)*(x*(x-1)/2)
}

func stepCount(grid map[image.Point]byte, loc []image.Point, n int) int {
	for i := 0; i < n; i++ {
		loc = step(grid, loc)
	}
	return len(loc)
}

func step(grid map[image.Point]byte, loc []image.Point) []image.Point {
	b := util.Bounds(maps.Keys(grid))
	ret := map[image.Point]struct{}{}
	for _, p := range loc {
		for _, n := range util.Neighbors4 {
			if grid[p.Add(n).Mod(b)] != '#' {
				ret[p.Add(n)] = struct{}{}
			}
		}
	}
	return maps.Keys(ret)
}

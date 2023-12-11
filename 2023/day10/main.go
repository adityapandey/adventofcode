package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/adityapandey/adventofcode/util"
)

var pipes = map[byte][2]util.Dir{
	'|': {util.N, util.S},
	'-': {util.E, util.W},
	'L': {util.N, util.E},
	'J': {util.N, util.W},
	'7': {util.S, util.W},
	'F': {util.S, util.E},
}

func main() {
	grid := map[image.Point]byte{}
	var start image.Point
	s := util.ScanAll()
	for y := 0; s.Scan(); y++ {
		for x := range s.Text() {
			grid[image.Pt(x, y)] = s.Text()[x]
			if s.Text()[x] == 'S' {
				start = image.Pt(x, y)
			}
		}
	}

	grid[start] = getStartTile(start, grid)
	var path []image.Point
	sumArea := 0
	for curr, next := start, start; curr == start || next != start; path = append(path, curr) {
		curr, next = next, start
		for _, d := range pipes[grid[curr]] {
			if _, ok := grid[curr.Add(d.PointR())]; ok && !slices.Contains(path, curr.Add(d.PointR())) {
				next = curr.Add(d.PointR())
			}
		}
		// https://en.wikipedia.org/wiki/Shoelace_formula
		sumArea += curr.X*next.Y - curr.Y*next.X
	}

	fmt.Println(len(path) / 2)
	fmt.Println((util.Abs(sumArea)-len(path))/2 + 1)
}

func getStartTile(start image.Point, grid map[image.Point]byte) byte {
	var dirs [2]util.Dir
	i := 0
	for _, n := range util.Neighbors4 {
		nn := start.Add(n)
		if c, ok := grid[nn]; ok {
			for _, dir := range pipes[c] {
				if nn.Add(dir.PointR()) == start {
					dirs[i] = dir.Reverse()
					i++
				}
			}
		}
	}
	for p, d := range pipes {
		if d == dirs || d == [2]util.Dir{dirs[1], dirs[0]} {
			return p
		}
	}
	panic("Unable to determine start tile")
}

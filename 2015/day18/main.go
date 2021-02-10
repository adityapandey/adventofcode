package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	grid := make(map[image.Point]struct{})
	var w, h int
	s := util.ScanAll()
	for s.Scan() {
		line := s.Text()
		w = len(line)
		for x := range line {
			if line[x] == '#' {
				grid[image.Pt(x, h)] = struct{}{}
			}
		}
		h++
	}
	gridcopy := copymap(grid)
	for i := 0; i < 100; i++ {
		gridcopy = iterate(gridcopy, w, h, false)
	}
	fmt.Println(len(gridcopy))
	gridcopy = copymap(grid)
	for i := 0; i < 100; i++ {
		gridcopy = iterate(gridcopy, w, h, true)
	}
	fmt.Println(len(gridcopy))
}

func iterate(grid map[image.Point]struct{}, w, h int, stuck bool) map[image.Point]struct{} {
	stuckPixels := []image.Point{{0, 0}, {0, h - 1}, {w - 1, 0}, {w - 1, h - 1}}
	if stuck {
		for _, p := range stuckPixels {
			grid[p] = struct{}{}
		}
	}
	next := make(map[image.Point]struct{})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			p := image.Pt(x, y)
			var neighbours int
			for _, n := range util.Neighbors8 {
				q := p.Add(n)
				if q.X < 0 || q.Y < 0 || q.X >= w || q.Y >= h {
					continue
				}
				if _, ok := grid[q]; ok {
					neighbours++
				}
			}
			_, on := grid[p]
			if (on && (neighbours == 2 || neighbours == 3)) || (!on && neighbours == 3) {
				next[p] = struct{}{}
			}
		}
	}
	if stuck {
		for _, p := range stuckPixels {
			next[p] = struct{}{}
		}
	}
	return next
}

func copymap(m map[image.Point]struct{}) map[image.Point]struct{} {
	n := make(map[image.Point]struct{})
	for k := range m {
		n[k] = struct{}{}
	}
	return n
}

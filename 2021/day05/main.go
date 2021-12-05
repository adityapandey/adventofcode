package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	grid := map[image.Point]int{}
	gridWithDiags := map[image.Point]int{}
	s := util.ScanAll()
	for s.Scan() {
		var from, to image.Point
		fmt.Sscanf(s.Text(), "%d,%d -> %d,%d", &from.X, &from.Y, &to.X, &to.Y)
		if from.X == to.X || from.Y == to.Y {
			for _, p := range between(from, to) {
				grid[p]++
				gridWithDiags[p]++
			}
		} else {
			for _, p := range between(from, to) {
				gridWithDiags[p]++
			}
		}
	}
	sum, sumWithDiags := 0, 0
	for _, n := range grid {
		if n >= 2 {
			sum++
		}
	}
	fmt.Println(sum)
	for _, n := range gridWithDiags {
		if n >= 2 {
			sumWithDiags++
		}
	}
	fmt.Println(sumWithDiags)
}

func between(from, to image.Point) []image.Point {
	var ps []image.Point
	xstep, ystep := util.Sign(to.X-from.X), util.Sign(to.Y-from.Y)
	for p := from; p != to; p = p.Add(image.Pt(xstep, ystep)) {
		ps = append(ps, p)
	}
	ps = append(ps, to)
	return ps
}

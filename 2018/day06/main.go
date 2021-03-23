package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

type closest struct {
	p        image.Point
	d        int
	multiple bool
}

func main() {
	var points []image.Point
	s := util.ScanAll()
	for s.Scan() {
		var p image.Point
		fmt.Sscanf(s.Text(), "%d, %d", &p.X, &p.Y)
		points = append(points, p)
	}
	bounds := util.Bounds(points)

	grid := make(map[image.Point]closest)
	for _, p := range points {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				curr := image.Pt(x, y)
				distance := util.Manhattan(curr, p)
				c, ok := grid[curr]
				if ok && c.d == distance {
					c.multiple = true
					grid[curr] = c
				}
				if !ok || distance < c.d {
					grid[curr] = closest{p: p, d: distance}
				}
			}
		}
	}
	areas := make(map[image.Point]int)
	border := make(map[image.Point]struct{})
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			curr := image.Pt(x, y)
			if grid[curr].multiple {
				continue
			}
			areas[grid[curr].p]++
			if !curr.In(bounds.Inset(1)) {
				border[grid[curr].p] = struct{}{}
			}

		}
	}
	var maxArea int
	for p, area := range areas {
		if _, ok := border[p]; !ok && area > maxArea {
			maxArea = area
		}
	}
	fmt.Println(maxArea)

	totalDistance := make(map[image.Point]int)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			curr := image.Pt(x, y)
			for _, p := range points {
				totalDistance[curr] += util.Manhattan(curr, p)
			}
		}
	}
	var safeArea int
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if totalDistance[image.Pt(x, y)] < 10000 {
				safeArea++
			}
		}
	}
	fmt.Println(safeArea)
}

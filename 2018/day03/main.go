package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

type claim struct {
	id int
	r  image.Rectangle
}

func main() {
	s := util.ScanAll()
	var claims []claim
	grid := make(map[image.Point]int)
	for s.Scan() {
		var id, x, y, w, h int
		fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		c := claim{id, image.Rect(x, y, x+w-1, y+h-1)}
		for x := c.r.Min.X; x <= c.r.Max.X; x++ {
			for y := c.r.Min.Y; y <= c.r.Max.Y; y++ {
				grid[image.Pt(x, y)]++
			}
		}
		claims = append(claims, c)
	}
	var overlap int
	for _, v := range grid {
		if v > 1 {
			overlap++
		}
	}
	fmt.Println(overlap)

	var nonOverlap int
	for _, c := range claims {
		if !hasOverlap(c, grid) {
			nonOverlap = c.id
			break
		}
	}
	fmt.Println(nonOverlap)
}

func hasOverlap(c claim, grid map[image.Point]int) bool {
	for x := c.r.Min.X; x <= c.r.Max.X; x++ {
		for y := c.r.Min.Y; y <= c.r.Max.Y; y++ {
			if grid[image.Pt(x, y)] > 1 {
				return true
			}
		}
	}
	return false
}

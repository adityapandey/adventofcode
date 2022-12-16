package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	grid := map[image.Point]int{}
	visible := map[image.Point]struct{}{}
	s := util.ScanAll()
	y := 0
	for s.Scan() {
		for x, b := range s.Bytes() {
			grid[image.Pt(x, y)] = int(b - '0')
		}
		y++
	}

	maxScore := 0
	for p := range grid {
		score := 1
		for _, n := range util.Neighbors4 {
			next, view := p, 0
			for {
				next = next.Add(n)
				if _, ok := grid[next]; ok {
					view++
					if grid[next] >= grid[p] {
						score *= view
						break
					}
				} else {
					visible[p] = struct{}{}
					score *= view
					break
				}
			}
		}

		if score > maxScore {
			maxScore = score
		}
	}
	fmt.Println(len(visible))
	fmt.Println(maxScore)
}

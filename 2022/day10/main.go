package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	x := []int{1}
	for _, line := range strings.Split(util.ReadAll(), "\n") {
		switch line {
		case "noop":
			x = append(x, x[len(x)-1])
		default:
			var n int
			fmt.Sscanf(line, "addx %d", &n)
			x = append(x, x[len(x)-1])
			x = append(x, x[len(x)-1]+n)
		}
	}

	sum := 0
	for i := range x {
		if (i-19)%40 == 0 {
			sum += (i + 1) * x[i]
		}
	}
	fmt.Println(sum)

	grid := map[image.Point]struct{}{}
	for i := range x {
		crtx, crty := i%40, i/40
		if util.Abs(crtx-x[i]) <= 1 {
			grid[image.Pt(crtx, crty)] = struct{}{}
		} else {
			delete(grid, image.Pt(crtx, crty))
		}
	}

	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if _, ok := grid[image.Pt(x, y)]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

package util

import (
	"fmt"
	"image"
	"log"
	"math"
	"strconv"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func DrawGrid(grid map[image.Point]byte) {
	minx, maxx := math.MaxInt32, 0
	miny, maxy := math.MaxInt32, 0
	for p := range grid {
		x, y := p.X, p.Y
		if x < minx {
			minx = x

		}
		if y < miny {
			miny = y
		}
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if b, ok := grid[image.Pt(x, y)]; ok {
				fmt.Printf("%c", b)
			} else {
				fmt.Printf("?")
			}
		}
		fmt.Println()
	}
}

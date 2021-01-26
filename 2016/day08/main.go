package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

const (
	width  = 50
	height = 6
)

func main() {
	var screen [height][width]bool
	s := util.ScanAll()
	for s.Scan() {
		line := s.Text()
		switch {
		case strings.HasPrefix(line, "rect"):
			var w, h int
			fmt.Sscanf(line, "rect %dx%d", &w, &h)
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					screen[y][x] = true
				}
			}
		case strings.HasPrefix(line, "rotate row"):
			var row, offset int
			fmt.Sscanf(line, "rotate row y=%d by %d", &row, &offset)
			var next [width]bool
			for i := 0; i < width; i++ {
				next[(i+offset)%width] = screen[row][i]
			}
			screen[row] = next
		case strings.HasPrefix(line, "rotate column"):
			var col, offset int
			fmt.Sscanf(line, "rotate column x=%d by %d", &col, &offset)
			var next [height]bool
			for i := 0; i < height; i++ {
				next[(i+offset)%height] = screen[i][col]
			}
			for i := 0; i < height; i++ {
				screen[i][col] = next[i]
			}
		}
	}
	var sum int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if screen[y][x] {
				sum++
			}
		}
	}
	fmt.Println(sum)
	draw(screen)
}

func draw(screen [height][width]bool) {
	d := map[bool]byte{true: '#', false: ' '}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%c", d[screen[y][x]])
		}
		fmt.Println()
	}
}

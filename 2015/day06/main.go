package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	lights1 := make(map[image.Point]int)
	lights2 := make(map[image.Point]int)
	s := util.ScanAll()
	for s.Scan() {
		apply(s.Text(), lights1, false)
		apply(s.Text(), lights2, true)
	}
	var sum1, sum2 int
	for _, l := range lights1 {
		if l == 1 {
			sum1++
		}
	}
	for _, l := range lights2 {
		sum2 += l
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func apply(rule string, lights map[image.Point]int, elvish bool) {
	var r image.Rectangle
	f := strings.Fields(rule)
	switch f[0] {
	case "turn":
		fmt.Sscanf(f[2], "%d,%d", &r.Min.X, &r.Min.Y)
		fmt.Sscanf(f[4], "%d,%d", &r.Max.X, &r.Max.Y)
		switch f[1] {
		case "on":
			for x := r.Min.X; x <= r.Max.X; x++ {
				for y := r.Min.Y; y <= r.Max.Y; y++ {
					if elvish {
						lights[image.Pt(x, y)]++
					} else {
						lights[image.Pt(x, y)] = 1
					}
				}
			}
		case "off":
			for x := r.Min.X; x <= r.Max.X; x++ {
				for y := r.Min.Y; y <= r.Max.Y; y++ {
					if elvish {
						if lights[image.Pt(x, y)] > 0 {
							lights[image.Pt(x, y)]--
						}
					} else {
						lights[image.Pt(x, y)] = 0
					}
				}
			}
		}
	case "toggle":
		fmt.Sscanf(f[1], "%d,%d", &r.Min.X, &r.Min.Y)
		fmt.Sscanf(f[3], "%d,%d", &r.Max.X, &r.Max.Y)
		for x := r.Min.X; x <= r.Max.X; x++ {
			for y := r.Min.Y; y <= r.Max.Y; y++ {
				if elvish {
					lights[image.Pt(x, y)] += 2
				} else {
					lights[image.Pt(x, y)] = 1 - lights[image.Pt(x, y)]
				}
			}
		}
	}
}

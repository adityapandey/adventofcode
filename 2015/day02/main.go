package main

import (
	"fmt"
	"math"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var paper, ribbon int
	s := util.ScanAll()
	for s.Scan() {
		var l, w, h int
		fmt.Sscanf(s.Text(), "%dx%dx%d", &l, &w, &h)
		paper += (2*l*w + 2*w*h + 2*h*l) + min(l*w, w*h, h*l)
		ribbon += (l * w * h) + 2*min(l+w, w+h, h+l)
	}
	fmt.Println(paper)
	fmt.Println(ribbon)
}

func min(a ...int) int {
	min := math.MaxInt16
	for _, x := range a {
		if x < min {
			min = x
		}
	}
	return min
}

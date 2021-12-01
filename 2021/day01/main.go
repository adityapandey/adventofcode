package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var depths []int
	s := util.ScanAll()
	for s.Scan() {
		depths = append(depths, util.Atoi(s.Text()))
	}

	var increases int
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	fmt.Println(increases)

	windows := make([]int, len(depths)-2)
	for i := range windows {
		windows[i] = depths[i] + depths[i+1] + depths[i+2]
	}

	increases = 0
	for i := 1; i < len(windows); i++ {
		if windows[i] > windows[i-1] {
			increases++
		}
	}
	fmt.Println(increases)
}

package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var deltas []int
	seen := make(map[int]struct{})
	var freq, twice int
	var found bool
	s := util.ScanAll()
	for s.Scan() {
		curr := util.Atoi(s.Text())
		deltas = append(deltas, curr)
		freq += curr
		if _, ok := seen[freq]; ok && !found {
			twice, found = freq, true
		}
		seen[freq] = struct{}{}
	}
	fmt.Println(freq)

	for i := 0; !found; i = (i + 1) % len(deltas) {
		freq += deltas[i]
		if _, ok := seen[freq]; ok {
			twice, found = freq, true
		}
	}
	fmt.Println(twice)
}

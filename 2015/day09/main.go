package main

import (
	"fmt"
	"math"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	paths := make(map[string]map[string]int)
	s := util.ScanAll()
	for s.Scan() {
		var from, to string
		var d int
		fmt.Sscanf(s.Text(), "%s to %s = %d", &from, &to, &d)
		if len(paths[from]) == 0 {
			paths[from] = make(map[string]int)
		}
		paths[from][to] = d
		if len(paths[to]) == 0 {
			paths[to] = make(map[string]int)
		}
		paths[to][from] = d
	}
	min, max := math.MaxInt16, 0
	for start := range paths {
		shortest, longest := pathLengths(paths, start, map[string]struct{}{})
		if shortest < min {
			min = shortest
		}
		if longest > max {
			max = longest
		}
	}
	fmt.Println(min)
	fmt.Println(max)
}

func pathLengths(paths map[string]map[string]int, from string, seen map[string]struct{}) (int, int) {
	seen[from] = struct{}{}
	if len(seen) == len(paths) {
		return 0, 0
	}
	min, max := math.MaxInt16, 0
	for to := range paths[from] {
		if _, ok := seen[to]; ok {
			continue
		}
		newseen := make(map[string]struct{})
		for k := range seen {
			newseen[k] = struct{}{}
		}
		stepsMin, stepsMax := pathLengths(paths, to, newseen)
		stepsMin += paths[from][to]
		stepsMax += paths[from][to]
		if stepsMin < min {
			min = stepsMin
		}
		if stepsMax > max {
			max = stepsMax
		}
	}
	return min, max
}

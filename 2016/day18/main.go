package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := util.ReadAll()
	fmt.Println(numSafe(input, 40))
	fmt.Println(numSafe(input, 400000))
}

func numSafe(start string, size int) int {
	row := make(map[int]bool) // true = trap
	var sum int
	for x := range start {
		switch start[x] {
		case '^':
			row[x] = true
		case '.':
			sum++
		}
	}
	l := len(start)
	for y := 1; y < size; y++ {
		nextrow := make(map[int]bool)
		for x := 0; x < l; x++ {
			l, r := row[x-1], row[x+1]
			//Since (l && c && !r) || (!l && c && r) || (l && !c && !r) || (!l && !c && r)
			// = (!l && r) || (l && !r)
			if (!l && r) || (l && !r) {
				nextrow[x] = true
			} else {
				sum++
			}
		}
		row = nextrow
	}
	return sum
}

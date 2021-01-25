// https://adventofcode.com/2020/day/15
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var input string
	fmt.Fscanf(os.Stdin, "%s", &input)
	var start []int
	for _, s := range strings.Split(input, ",") {
		start = append(start, util.Atoi(s))
	}
	lastSpoken := make(map[int]int)
	last := start[0]
	for i := 1; i < len(start); i++ {
		lastSpoken[last] = i
		last = start[i]
	}

	// Part 1
	for i := len(start); i < 2020; i++ {
		var next int
		if turn, ok := lastSpoken[last]; ok {
			next = i - turn
		} else {
			next = 0
		}
		lastSpoken[last] = i
		last = next
	}
	fmt.Println(last)

	// Part 2
	for i := 2020; i < 30000000; i++ {
		var next int
		if turn, ok := lastSpoken[last]; ok {
			next = i - turn
		} else {
			next = 0
		}
		lastSpoken[last] = i
		last = next
	}
	fmt.Println(last)
}

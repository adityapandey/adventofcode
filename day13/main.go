// https://adventofcode.com/2020/day/13
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var when int
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Sscanf(s.Text(), "%d", &when)
	s.Scan()
	ids := make(map[int]int)
	for i, t := range strings.Split(s.Text(), ",") {
		b, err := strconv.Atoi(t)
		if err == nil {
			ids[b] = i
		}
	}

	// Part 1
	minWait, bus := math.MaxInt32, 0
	for b := range ids {
		wait := b - (when % b)
		if wait < minWait {
			minWait, bus = wait, b
		}
	}
	fmt.Println(minWait * bus)

	// Part 2
	t, multiplier := 0, 1
	for k, v := range ids {
		for (t+v)%k != 0 {
			t += multiplier
		}
		multiplier *= k // actually lcm(multiplier, k)
	}
	fmt.Println(t)
}

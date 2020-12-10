// https://adventofcode.com/2020/day/10
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	joltages := []int{0}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var j int
		fmt.Sscanf(s.Text(), "%d", &j)
		joltages = append(joltages, j)
	}
	sort.Ints(joltages)
	joltages = append(joltages, joltages[len(joltages)-1]+3)

	// Part 1
	m := make(map[int]int)
	for i := 1; i < len(joltages); i++ {
		m[joltages[i]-joltages[i-1]]++
	}
	fmt.Println(m[1] * m[3])

	// Part 2
	m = map[int]int{0: 1}
	for i := 1; i < len(joltages); i++ {
		for j := max(0, i-3); j < i; j++ {
			if joltages[i]-joltages[j] <= 3 {
				m[i] += m[j]
			}
		}
	}
	fmt.Println(m[len(joltages)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := make(map[int]int)
	s := bufio.NewScanner(os.Stdin)
	i := 0
	for s.Scan() {
		var n int
		fmt.Sscanf(s.Text(), "%d", &n)
		m[i] = n
		i++
	}
	fmt.Println(numSteps(mapCopy(m), false))
	fmt.Println(numSteps(mapCopy(m), true))
}

func numSteps(m map[int]int, part2 bool) int {
	i := 0
	for ip := 0; ip < len(m); i++ {
		next := ip + m[ip]
		if part2 && m[ip] >= 3 {
			m[ip]--
		} else {
			m[ip]++
		}
		ip = next
	}
	return i
}

func mapCopy(m map[int]int) map[int]int {
	c := make(map[int]int)
	for k, v := range m {
		c[k] = v
	}
	return c
}

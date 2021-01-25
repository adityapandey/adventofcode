// https://adventofcode.com/2020/day/1
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	m := make(map[int]struct{})
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		m[util.Atoi(s.Text())] = struct{}{}
	}

	// Part 1
	for k := range m {
		if _, ok := m[2020-k]; ok {
			fmt.Println(k * (2020 - k))
			break
		}
	}

	// Part 2:
outer:
	for k1 := range m {
		for k2 := range m {
			if k1 == k2 || k1+k2 > 2020 {
				continue
			}
			if _, ok := m[2020-(k1+k2)]; ok {
				fmt.Println(k1 * k2 * (2020 - (k1 + k2)))
				break outer
			}
		}
	}
}

package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var sum2, sum3 int
	m := make(map[[2]string]struct{})
	var common string
	s := util.ScanAll()
	for s.Scan() {
		id := s.Text()
		twos, threes := repetitions(id)
		sum2 += twos
		sum3 += threes
		for i := range id {
			k := [2]string{id[:i], id[i+1:]}
			if _, ok := m[k]; ok {
				common = k[0] + k[1]
			}
			m[k] = struct{}{}
		}
	}
	fmt.Println(sum2 * sum3)
	fmt.Println(common)
}

func repetitions(s string) (int, int) {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	var twos, threes int
	for _, v := range m {
		if v == 2 {
			twos = 1
		}
		if v == 3 {
			threes = 1
		}
	}
	return twos, threes
}

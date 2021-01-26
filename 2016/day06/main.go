package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	m := make(map[int]map[byte]int)
	s := util.ScanAll()
	for s.Scan() {
		line := s.Text()
		for i := range line {
			if len(m[i]) == 0 {
				m[i] = make(map[byte]int)
			}
			m[i][line[i]]++
		}
	}
	var size int
	for i := range m {
		if i > size {
			size = i
		}
	}
	var mostCommon, leastCommon []byte
	for i := 0; i <= size; i++ {
		max, min := 0, len(m[0])
		var maxb, minb byte
		for b, count := range m[i] {
			if count > max {
				maxb, max = b, count
			}
			if count < min {
				minb, min = b, count
			}
		}
		mostCommon = append(mostCommon, maxb)
		leastCommon = append(leastCommon, minb)
	}
	fmt.Println(string(mostCommon))
	fmt.Println(string(leastCommon))
}

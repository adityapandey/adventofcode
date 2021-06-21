package main

import (
	"fmt"
	"os"
)

func main() {
	var min, max int
	fmt.Fscanf(os.Stdin, "%d-%d", &min, &max)

	var sum1, sum2 int
	for i := min; i <= max; i++ {
		if isValid(i, false) {
			sum1++
		}
		if isValid(i, true) {
			sum2++
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func isValid(n int, part2 bool) bool {
	ds := digits(n)
	for i := 1; i < len(ds); i++ {
		if ds[i] < ds[i-1] {
			return false
		}
	}
	m := make(map[byte]int)
	for _, d := range ds {
		m[d]++
	}
	for _, v := range m {
		if !part2 {
			if v > 1 {
				return true
			}
		} else {
			if v == 2 {
				return true
			}
		}
	}
	return false
}

func digits(n int) []byte {
	return []byte(fmt.Sprint(n))
}

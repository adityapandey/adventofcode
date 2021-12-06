package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	fish := map[int]int{}
	for _, n := range strings.Split(util.ReadAll(), ",") {
		fish[util.Atoi(n)]++
	}
	for i := 0; i < 80; i++ {
		fish = step(fish)
	}
	fmt.Println(count(fish))
	for i := 80; i < 256; i++ {
		fish = step(fish)
	}
	fmt.Println(count(fish))
}

func step(s map[int]int) map[int]int {
	m := map[int]int{}
	for k, v := range s {
		if k == 0 {
			m[6] += v
			m[8] += v
		} else {
			m[k-1] += v
		}
	}
	return m
}

func count(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

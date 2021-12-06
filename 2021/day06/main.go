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
		step(fish)
	}
	fmt.Println(count(fish))
	for i := 80; i < 256; i++ {
		step(fish)
	}
	fmt.Println(count(fish))
}

func step(fish map[int]int) {
	births := fish[0]
	for i := 1; i <= 8; i++ {
		fish[i-1] = fish[i]
	}
	fish[6] += births
	fish[8] = births
}

func count(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

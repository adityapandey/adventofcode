package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var fish [9]int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		fish[util.Atoi(n)]++
	}
	for steps := 1; steps <= 256; steps++ {
		step(&fish)
		if steps == 80 || steps == 256 {
			fmt.Println(count(fish))
		}
	}
}

func step(fish *[9]int) {
	births := fish[0]
	for i := 1; i <= 8; i++ {
		fish[i-1] = fish[i]
	}
	fish[6] += births
	fish[8] = births
}

func count(fish [9]int) int {
	sum := 0
	for _, v := range fish {
		sum += v
	}
	return sum
}

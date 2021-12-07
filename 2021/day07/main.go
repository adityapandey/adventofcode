package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var crabs []int
	min, max := math.MaxInt, 0
	for _, s := range strings.Split(util.ReadAll(), ",") {
		c := util.Atoi(s)
		crabs = append(crabs, c)
		min, max = util.Min(min, c), util.Max(max, c)
	}
	minFuel, minFuelStepup := math.MaxInt, math.MaxInt
	for pos := min; pos <= max; pos++ {
		minFuel = util.Min(minFuel, fuel(crabs, pos, func(n int) int { return n }))
		minFuelStepup = util.Min(minFuelStepup, fuel(crabs, pos, func(n int) int { return n * (n + 1) / 2 }))
	}
	fmt.Println(minFuel)
	fmt.Println(minFuelStepup)
}

func fuel(crabs []int, pos int, cost func(int) int) int {
	var sum int
	for _, c := range crabs {
		sum += cost(util.Abs(c - pos))
	}
	return sum
}

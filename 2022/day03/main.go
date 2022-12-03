package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	sum := 0
	rucksacks := strings.Split(util.ReadAll(), "\n")
	for _, s := range rucksacks {
		b := []byte(s)
		first, second := util.SetOf(b[:len(b)/2]), util.SetOf(b[len(b)/2:])
		c := first.Intersect(second).Values()[0]
		sum += score(c)
	}
	fmt.Println(sum)

	sum = 0
	for i := 0; i < len(rucksacks); i += 3 {
		first := util.SetOf([]byte(rucksacks[i]))
		second := util.SetOf([]byte(rucksacks[i+1]))
		third := util.SetOf([]byte(rucksacks[i+2]))
		c := first.Intersect(second).Intersect(third).Values()[0]
		sum += score(c)
	}
	fmt.Println(sum)
}

func score(c byte) int {
	if c <= 'Z' {
		return int(c) - 'A' + 27
	}
	return int(c) - 'a' + 1
}

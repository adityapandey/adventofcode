package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	elves := strings.Split(util.ReadAll(), "\n\n")
	l := len(elves)
	calories := make([]int, l)
	for i, elf := range elves {
		sum := 0
		for _, line := range strings.Split(elf, "\n") {
			sum += util.Atoi(line)
		}
		calories[i] = sum
	}
	sort.Ints(calories)
	fmt.Println(calories[l-1])
	fmt.Println(calories[l-1] + calories[l-2] + calories[l-3])
}

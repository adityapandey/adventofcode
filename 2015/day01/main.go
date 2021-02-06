package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var floor, basement int
	input := util.ReadAll()
	for i := range input {
		switch input[i] {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 && basement == 0 {
			basement = i + 1
		}
	}
	fmt.Println(floor)
	fmt.Println(basement)
}

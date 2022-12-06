package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	s := util.ReadAll()
	fmt.Println(firstNUnique(s, 4))
	fmt.Println(firstNUnique(s, 14))
}

func firstNUnique(s string, n int) int {
	for i := n; i < len(s); i++ {
		b := []byte(s[i-n : i])
		if len(b) == len(util.SetOf(b)) {
			return i
		}
	}
	return -1
}

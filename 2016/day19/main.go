package main

import (
	"fmt"
	"math/bits"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := util.Atoi(util.ReadAll())
	fmt.Println(whiteElephant1(uint(input)))
	fmt.Println(whiteElephant2(input))
}

func whiteElephant1(n uint) uint {
	// Josephus recurrence: rotate left 1 in binary
	// 1abc => abc1
	mask := 1<<bits.Len(n) - uint(1)
	return n<<1&mask + 1
}

func whiteElephant2(n int) int {
	// Generalized Josephus: In ternary
	// 1000 => 1000
	// 1xxx => 0xxx
	// 2xxx => 0xxx*2 + 1000
	m, p := n, 1
	for m >= 3 {
		m /= 3
		p *= 3
	}
	if n/p == 1 {
		if w := n % p; w > 0 {
			return w
		}
		return n
	}
	return 2*(n%p) + p
}

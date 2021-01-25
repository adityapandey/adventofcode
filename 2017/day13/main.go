package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	m := make(map[int]int)
	var maxdepth int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		matches := strings.Split(s.Text(), ": ")
		depth, rng := util.Atoi(matches[0]), util.Atoi(matches[1])
		m[depth] = rng
		if depth > maxdepth {
			maxdepth = depth
		}
	}
	catches := getCatches(m, maxdepth, 0)
	var sum int
	for i := range catches {
		sum += catches[i] * m[catches[i]]
	}
	fmt.Println(sum)

	delay := 0
	for len(getCatches(m, maxdepth, delay)) != 0 {
		delay++
	}
	fmt.Println(delay)
}

func getCatches(m map[int]int, maxdepth, delay int) []int {
	var catches []int
	for i := delay; i <= maxdepth+delay; i++ {
		if rng, ok := m[i-delay]; ok && pos(rng, i) == 0 {
			catches = append(catches, i)
		}
	}
	return catches
}

// for range = 5, e.g.
// step 0 1 2 3 4 5 6 7  8 9 10
//      0 1 2 3 4 3 2 1  0 1 2
// period of 2*(range-1)
// max of (range-1)
func pos(rng, step int) int {
	return (rng - 1) - util.Abs(step%(2*(rng-1))-(rng-1))
}

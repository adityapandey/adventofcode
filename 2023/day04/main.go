package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	s := util.ReadAll()
	var w []int
	for _, card := range strings.Split(s, "\n") {
		w = append(w, wins(card))
	}

	dups := map[int]int{}
	var sum1, sum2 int
	for i, ww := range w {
		if ww > 0 {
			sum1 += 1 << (ww - 1)
		}
		dups[i]++
		sum2++
		for j := 0; j < ww; j++ {
			dups[i+1+j] += dups[i]
			sum2 += dups[i]
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func wins(card string) int {
	sp := strings.Split(strings.Split(card, ": ")[1], " | ")
	winNums := strings.Fields(sp[0])
	haveNums := strings.Fields(sp[1])
	w := 0
	for _, h := range haveNums {
		if slices.Contains(winNums, h) {
			w++
		}
	}
	return w
}

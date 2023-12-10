package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	sum1, sum2 := 0, 0
	s := util.ScanAll()
	for s.Scan() {
		var history []int
		for _, n := range strings.Fields(s.Text()) {
			history = append(history, util.Atoi(n))
		}
		sum1 += fill(history, true)
		sum2 += fill(history, false)
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func diff(history []int) []int {
	d := make([]int, len(history)-1)
	for i := 0; i < len(history)-1; i++ {
		d[i] = history[i+1] - history[i]
	}
	return d
}

func fill(history []int, forward bool) int {
	d := [][]int{history}
	for !allZeros(d[len(d)-1]) {
		d = append(d, diff(d[len(d)-1]))
	}
	curr := 0
	for i := len(d) - 2; i >= 0; i-- {
		if forward {
			curr += d[i][len(d[i])-1]
		} else {
			curr = d[i][0] - curr
		}
	}
	return curr
}

func allZeros(seq []int) bool {
	for _, n := range seq {
		if n != 0 {
			return false
		}
	}
	return true
}

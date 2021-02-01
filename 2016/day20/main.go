package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type rng struct {
	min, max uint32
}

func main() {
	var rs []rng
	s := util.ScanAll()
	for s.Scan() {
		sp := strings.Split(s.Text(), "-")
		rs = append(rs, rng{uint32(util.Atoi(sp[0])), uint32(util.Atoi(sp[1]))})
	}
	sort.Slice(rs, func(i, j int) bool { return rs[i].min < rs[j].min })
	for i := 0; i < len(rs)-1; i++ {
		if rs[i+1].min-1 <= rs[i].max {
			rs[i].max = max(rs[i].max, rs[i+1].max)
			rs = append(rs[:i+1], rs[i+2:]...)
			i--
		}
	}

	fmt.Println(rs[0].max + 1)

	var allowed int
	for i := 0; i < len(rs)-1; i++ {
		allowed += int(rs[i+1].min - rs[i].max - 1)
	}
	fmt.Println(allowed)
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

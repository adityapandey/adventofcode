package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`Disc #(\d+) has (\d+) positions; at time=0, it is at position (\d+).`)

type disc [2]int

func (d disc) pos(t int) int {
	return (t + d[1]) % d[0]
}

func main() {
	var discs []disc
	s := util.ScanAll()
	for s.Scan() {
		matches := re.FindAllStringSubmatch(s.Text(), -1)[0]
		discs = append(discs, disc{util.Atoi(matches[2]), util.Atoi(matches[1]) + util.Atoi(matches[3])})
	}

	fmt.Println(findSlot(discs))
	discs = append(discs, disc{11, len(discs) + 1})
	fmt.Println(findSlot(discs))
}

func findSlot(discs []disc) int {
	sort.Slice(discs, func(i, j int) bool { return discs[i][0] > discs[j][0] })
	t, multiplier := 0, 1
	for _, d := range discs {
		for (t+d[1])%d[0] != 0 {
			t += multiplier
		}
		multiplier *= d[0] // actually lcm(multiplier, k)
	}
	return t
}

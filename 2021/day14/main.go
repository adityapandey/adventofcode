package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type pair [2]byte

func main() {
	rules := map[pair][2]pair{}
	sp := strings.Split(util.ReadAll(), "\n\n")
	template := strings.Split(sp[0], "\n")[0]
	for _, l := range strings.Split(sp[1], "\n") {
		var from0, from1, to byte
		fmt.Sscanf(l, "%c%c -> %c", &from0, &from1, &to)
		rules[pair{from0, from1}] = [2]pair{{from0, to}, {to, from1}}
	}

	start, end := template[0], template[len(template)-1]
	polymer := map[pair]int{}
	for i := 0; i < len(template)-1; i++ {
		polymer[pair{template[i], template[i+1]}]++
	}
	for step := 1; step <= 40; step++ {
		polymer = next(polymer, rules)
		if step == 10 || step == 40 {
			fmt.Println(diffFreq(polymer, start, end))
		}
	}
}

func next(polymer map[pair]int, rules map[pair][2]pair) map[pair]int {
	next := map[pair]int{}
	for p, v := range polymer {
		next[rules[p][0]] += v
		next[rules[p][1]] += v
	}
	return next
}

func diffFreq(polymer map[pair]int, start, end byte) int {
	freq := map[byte]int{}
	for p, v := range polymer {
		freq[p[0]] += v
		freq[p[1]] += v
	}
	freq[start]++
	freq[end]++

	var sortFreq []int
	for _, v := range freq {
		sortFreq = append(sortFreq, v/2)
	}
	sort.Ints(sortFreq)
	return sortFreq[len(sortFreq)-1] - sortFreq[0]
}

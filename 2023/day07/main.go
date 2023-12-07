package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

type camelcards struct {
	hand string
	bid  int
}

func (c camelcards) strength(jIsJoker bool) int {
	m := map[rune]int{}
	for _, r := range c.hand {
		m[r]++
	}

	numJs := m['J']
	if jIsJoker {
		delete(m, 'J')
	}

	v := maps.Values(m)
	slices.Sort(v)
	slices.Reverse(v)

	// Possible: 5, 41, 32, 311, 221, 2111, 11111

	if jIsJoker {
		if len(v) == 0 {
			v = append(v, 0)
		}
		v[0] += numJs
	}

	if len(v) == 1 {
		v = append(v, 0)
	}

	return 10*v[0] + v[1]
}

func main() {
	var c []camelcards
	for _, line := range strings.Split(util.ReadAll(), "\n") {
		f := strings.Fields(line)
		c = append(c, camelcards{f[0], util.Atoi(f[1])})
	}

	sortCards(c, "23456789TJQKA", false)
	sum := 0
	for i := range c {
		sum += (i + 1) * c[i].bid
	}
	fmt.Println(sum)

	sortCards(c, "J23456789TQKA", true)
	sum = 0
	for i := range c {
		sum += (i + 1) * c[i].bid
	}
	fmt.Println(sum)
}

func sortCards(c []camelcards, cardOrder string, jIsJoker bool) {
	val := map[byte]int{}
	for i := range cardOrder {
		val[cardOrder[i]] = i
	}
	slices.SortFunc(c, func(a, b camelcards) int {
		d := a.strength(jIsJoker) - b.strength(jIsJoker)
		if d != 0 {
			return d
		}
		for i := 0; i < len(a.hand) && i < len(b.hand); i++ {
			d := val[a.hand[i]] - val[b.hand[i]]
			if d != 0 {
				return d
			}
		}
		return 0
	})
}

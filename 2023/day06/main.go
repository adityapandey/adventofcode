package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type race struct {
	t, d int
}

func (r race) wins() int {
	c := 0
	for i := 0; i <= r.t; i++ {
		if i*(r.t-i) > r.d {
			c++
		}
	}
	return c
}

func main() {
	sp := strings.Split(util.ReadAll(), "\n")
	var races []race
	for _, t := range strings.Fields(sp[0])[1:] {
		races = append(races, race{t: util.Atoi(t)})
	}
	for i, d := range strings.Fields(sp[1])[1:] {
		races[i].d = util.Atoi(d)
	}
	prod := 1
	for _, r := range races {
		prod *= r.wins()
	}
	fmt.Println(prod)

	var onerace race
	fmt.Sscanf(strings.Join(strings.Fields(sp[0])[1:], ""), "%d", &onerace.t)
	fmt.Sscanf(strings.Join(strings.Fields(sp[1])[1:], ""), "%d", &onerace.d)
	fmt.Println(onerace.wins())
}

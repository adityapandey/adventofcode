package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type rng struct {
	min, max int
	delta    int
}

func split(r rng, s []rng) []rng {
	slices.SortFunc(s, func(a, b rng) int { return a.min - b.min })
	p := []int{r.min, r.max}
	for i := range s {
		p = append(p, s[i].min, s[i].max)
	}
	slices.Sort(p)
	var ret []rng
	for i := sort.SearchInts(p, r.min); i < sort.SearchInts(p, r.max); i++ {
		delta := 0
		if j, ok := slices.BinarySearchFunc(s, p[i], func(x rng, target int) int {
			if target >= x.min && target < x.max {
				return 0
			}
			return x.min - target
		}); ok {
			delta = s[j].delta
		}
		ret = append(ret, rng{min: delta + p[i], max: delta + p[i+1]})
	}
	return ret
}

func splitmulti(r []rng, s []rng) []rng {
	var ret []rng
	for _, rr := range r {
		ret = append(ret, split(rr, s)...)
	}
	return ret
}

func main() {
	sp := strings.Split(util.ReadAll(), "\n\n")
	var seeds []int
	for _, s := range strings.Fields(sp[0])[1:] {
		seeds = append(seeds, util.Atoi(s))
	}

	var maps [][]rng
	for _, s := range sp[1:] {
		var m []rng
		for _, ss := range strings.Split(s, "\n")[1:] {
			var dest, src, r int
			fmt.Sscanf(ss, "%d %d %d", &dest, &src, &r)
			m = append(m, rng{
				min:   src,
				max:   src + r,
				delta: dest - src,
			})
		}
		maps = append(maps, m)
	}

	var rng1, rng2 []rng
	for i := range seeds {
		rng1 = append(rng1, rng{min: seeds[i], max: seeds[i] + 1})
		if i%2 == 0 {
			rng2 = append(rng2, rng{min: seeds[i], max: seeds[i] + seeds[i+1]})
		}
	}
	for _, m := range maps {
		rng1 = splitmulti(rng1, m)
		rng2 = splitmulti(rng2, m)
	}
	fmt.Println(min(rng1))
	fmt.Println(min(rng2))
}

func min(r []rng) int {
	min := math.MaxInt
	for _, rr := range r {
		min = util.Min(rr.min, min)
	}
	return min
}

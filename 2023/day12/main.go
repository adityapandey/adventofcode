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
		sp := strings.Fields(s.Text())
		condition := sp[0]
		spp := strings.Split(sp[1], ",")
		springs := make([]int, len(spp))
		for i := range spp {
			springs[i] = util.Atoi(spp[i])
		}
		sum1 += arrangements(condition, springs)

		unfoldCondition := strings.Repeat(condition+"?", 5)
		unfoldCondition = unfoldCondition[:len(unfoldCondition)-1]
		unfoldSprings := make([]int, 0, 5*len(springs))
		for i := 0; i < 5; i++ {
			unfoldSprings = append(unfoldSprings, springs...)
		}
		sum2 += arrangements(unfoldCondition, unfoldSprings)
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

// Okay to use same cache for both parts.
var cache = map[string]int{}

func arrangements(condition string, springs []int) int {
	if n, ok := cache[fmt.Sprint(condition, springs)]; ok {
		return n
	}
	sum := 0
	if len(condition) == 0 {
		if len(springs) == 0 {
			return 1
		}
		return 0
	}
	if condition[0] == '.' || condition[0] == '?' {
		sum += arrangements(condition[1:], springs)
	}
	if (condition[0] == '#' || condition[0] == '?') && len(springs) > 0 && len(condition) >= springs[0] && !strings.Contains(condition[:springs[0]], ".") {
		if len(condition) == springs[0] {
			if len(springs) == 1 {
				sum++
			}
		} else if condition[springs[0]] != '#' {
			sum += arrangements(condition[springs[0]+1:], springs[1:])
		}
	}
	cache[fmt.Sprint(condition, springs)] = sum
	return sum
}

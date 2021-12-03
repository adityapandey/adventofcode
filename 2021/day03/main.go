package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	s := strings.Split(util.ReadAll(), "\n")
	counts := bitCounts(s)
	var gamma, epsilon int
	v := 1
	for i := len(s[0]) - 1; i >= 0; i-- {
		if counts[i]['1'] > counts[i]['0'] {
			gamma += v
		} else {
			epsilon += v
		}
		v *= 2
	}
	fmt.Println(gamma * epsilon)

	o2, co2 := repeatKeeping(s, '1', '0'), repeatKeeping(s, '0', '1')
	fmt.Println(o2 * co2)
}

func bitCounts(s []string) map[int]map[byte]int {
	counts := map[int]map[byte]int{}
	for i := 0; i < len(s[0]); i++ {
		counts[i] = map[byte]int{}
	}
	for _, ss := range s {
		for j := 0; j < len(ss); j++ {
			counts[j][ss[j]]++
		}
	}
	return counts
}

func repeatKeeping(ss []string, moreOnes, moreZeros byte) int {
	s := make([]string, len(ss))
	copy(s, ss)
	var keep byte
	i := 0
	for len(s) > 1 {
		counts := bitCounts(s)
		if counts[i]['1'] >= counts[i]['0'] {
			keep = moreOnes
		} else {
			keep = moreZeros
		}
		for j := 0; j < len(s); j++ {
			if s[j][i] != keep {
				s = append(s[:j], s[j+1:]...)
				j--
			}
		}
		i++
	}
	sum, v := 0, 1
	for i := len(s[0]) - 1; i >= 0; i-- {
		sum += v * int(s[0][i]-'0')
		v *= 2
	}
	return sum
}

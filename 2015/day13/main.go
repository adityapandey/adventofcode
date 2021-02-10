package main

import (
	"fmt"
	"regexp"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).`)

func main() {
	m := make(map[string]map[string]int)
	s := util.ScanAll()
	for s.Scan() {
		matches := re.FindAllStringSubmatch(s.Text(), -1)[0]
		from, to, gainLose := matches[1], matches[4], matches[2]
		points := util.Atoi(matches[3])
		if gainLose == "lose" {
			points = -points
		}
		if len(m[from]) == 0 {
			m[from] = make(map[string]int)
		}
		m[from][to] = points
	}
	fmt.Println(optimal(m))
	m["self"] = make(map[string]int)
	for k := range m {
		if k == "self" {
			continue
		}
		m[k]["self"] = 0
		m["self"][k] = 0
	}
	fmt.Println(optimal(m))
}

func optimal(m map[string]map[string]int) int {
	var max int
	for _, p := range circularPermutations(m) {
		var sum int
		for i := range p {
			sum += m[p[i]][p[(i+1)%len(p)]] + m[p[(i+1)%len(p)]][p[i]]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func circularPermutations(m map[string]map[string]int) [][]string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	p := permutations(keys[1:])
	for i := range p {
		p[i] = append(p[i], keys[0])
	}
	return p
}

func permutations(arr []string) [][]string {
	res := [][]string{}
	var helper func([]string, int)
	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

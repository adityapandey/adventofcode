package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var packages []int
	s := util.ScanAll()
	for s.Scan() {
		var p int
		fmt.Sscanf(s.Text(), "%d", &p)
		packages = append(packages, p)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(packages)))
	fmt.Println(minQE(packages, 3))
	fmt.Println(minQE(packages, 4))
}

func minQE(packages []int, splits int) int {
	var sum int
	for _, p := range packages {
		sum += p
	}
	weight := sum / splits
	combinations := combineSum(packages, weight)
	sort.Slice(combinations, func(i, j int) bool { return len(combinations[i]) < len(combinations[j]) })
	var minLen int
	for i := range combinations {
		leftover := diff(packages, combinations[i])
		if len(combineSum(leftover, (sum-weight)/(splits-1))) > 0 {
			minLen = len(combinations[i])
			break
		}
	}
	qe := math.MaxInt64
	for i := range combinations {
		if len(combinations[i]) > minLen {
			continue
		}
		if len(combinations[i]) < minLen {
			break
		}
		prod := 1
		for j := range combinations[i] {
			prod *= combinations[i][j]
		}
		if prod < qe {
			qe = prod
		}
	}
	return qe
}

func diff(a []int, b []int) []int {
	leftover := make([]int, len(a))
	copy(leftover, a)
	for i, j := 0, 0; i < len(b); i, j = i+1, j+1 {
		for leftover[j] != b[i] {
			j++
		}
		leftover = append(leftover[:j], leftover[j+1:]...)
		j--
	}
	return leftover
}

var m = make(map[string][][]int)

func combineSum(a []int, target int) [][]int {
	if v, ok := m[hash(a, target)]; ok {
		return v
	}
	if target == 0 {
		return [][]int{{}}
	}
	if len(a) == 0 {
		return [][]int{}
	}
	var ret [][]int
	for i := range a {
		if a[i] > target {
			continue
		}
		c := make([]int, len(a))
		copy(c, a)
		g := combineSum(c[i+1:], target-a[i])
		for j := range g {
			ret = append(ret, append([]int{a[i]}, g[j]...))
		}
	}
	m[hash(a, target)] = ret
	return ret
}

func hash(a []int, target int) string { return fmt.Sprintf("%v %v", a, target) }

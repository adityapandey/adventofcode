package main

import (
	"fmt"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var boxes []int
	var sum int
	s := util.ScanAll()
	for s.Scan() {
		var b int
		fmt.Sscanf(s.Text(), "%d", &b)
		sum += b
		boxes = append(boxes, b)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(boxes)))

	c := combinations(boxes, 150, sum)
	fmt.Println(len(c))

	sort.Slice(c, func(i, j int) bool { return len(c[i]) < len(c[j]) })
	l := len(c[0])
	var i int
	for i = range c {
		if len(c[i]) > l {
			break
		}
	}
	fmt.Println(i)
}

func combinations(boxes []int, target, sum int) [][]int {
	if sum < target || len(boxes) == 0 {
		return [][]int{}
	}
	var ret [][]int
	if boxes[0] == target {
		ret = append(ret, []int{boxes[0]})
	}
	ret = append(ret, combinations(boxes[1:], target, sum-boxes[0])...)
	for _, c := range combinations(boxes[1:], target-boxes[0], sum-boxes[0]) {
		c = append([]int{boxes[0]}, c...)
		ret = append(ret, c)
	}
	return ret
}

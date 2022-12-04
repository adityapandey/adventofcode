package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	contains, overlaps := 0, 0
	s := util.ScanAll()
	for s.Scan() {
		var start1, end1, start2, end2 int
		fmt.Sscanf(s.Text(), "%d-%d,%d-%d", &start1, &end1, &start2, &end2)
		if isContained(start1, end1, start2, end2) {
			contains++
		}
		if isOverlapping(start1, end1, start2, end2) {
			overlaps++
		}
	}
	fmt.Println(contains)
	fmt.Println(overlaps)
}

func isContained(start1, end1, start2, end2 int) bool {
	return (start1 >= start2 && end1 <= end2) ||
		(start2 >= start1 && end2 <= end1)
}

func isOverlapping(start1, end1, start2, end2 int) bool {
	return isContained(start1, end1, start2, end2) ||
		(start1 >= start2 && start1 <= end2) ||
		(end1 >= start2 && end1 <= end2)
}

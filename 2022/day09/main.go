package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := util.ReadAll()
	fmt.Println(visited(input, 2))
	fmt.Println(visited(input, 10))
}

func visited(input string, ropelen int) int {
	rope := make([]image.Point, ropelen)
	visited := map[image.Point]struct{}{}
	for _, line := range strings.Split(input, "\n") {
		var b byte
		var n int
		fmt.Sscanf(line, "%c %d", &b, &n)
		d := util.DirFromByte(b)
		for i := 0; i < n; i++ {
			rope[0] = rope[0].Add(d.Point())
			for j := 1; j < ropelen; j++ {
				rope[j] = next(rope[j-1], rope[j])
			}
			visited[rope[ropelen-1]] = struct{}{}
		}
	}
	return len(visited)
}

func next(head, tail image.Point) image.Point {
	if util.Abs(head.X-tail.X) <= 1 && util.Abs(head.Y-tail.Y) <= 1 {
		return tail
	}
	return tail.Add(image.Pt(util.Sign(head.X-tail.X), util.Sign(head.Y-tail.Y)))
}

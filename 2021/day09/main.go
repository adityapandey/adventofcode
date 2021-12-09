package main

import (
	"fmt"
	"image"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	grid := map[image.Point]int{}
	lines := strings.Split(util.ReadAll(), "\n")
	for y, line := range lines {
		for x, n := range strings.Split(line, "") {
			grid[image.Pt(x, y)] = util.Atoi(n)
		}
	}
	h, w := len(lines), len(lines[0])

	var lows []image.Point
	sum := 0
	for p := range grid {
		isLow := true
		for _, n := range util.Neighbors4 {
			if p.Add(n).In(image.Rect(0, 0, w, h)) && grid[p.Add(n)] <= grid[p] {
				isLow = false
				break
			}
		}
		if isLow {
			lows = append(lows, p)
			sum += 1 + grid[p]
		}
	}
	fmt.Println(sum)

	var basinSizes []int
	for _, p := range lows {
		basinSizes = append(basinSizes, basinSize(grid, p, w, h))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}

func basinSize(grid map[image.Point]int, low image.Point, w, h int) int {
	visited := map[image.Point]struct{}{low: {}}
	next := []image.Point{low}
	for len(next) > 0 {
		p := next[0]
		next = next[1:]
		for _, n := range util.Neighbors4 {
			nn := p.Add(n)
			if _, ok := visited[nn]; !ok && nn.In(image.Rect(0, 0, w, h)) && grid[nn] != 9 {
				next = append(next, nn)
				visited[nn] = struct{}{}
			}
		}
	}
	return len(visited)
}

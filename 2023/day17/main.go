package main

import (
	"container/heap"
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

type crucible struct {
	p image.Point
	d util.Dir
	n int
}

func main() {
	grid := map[image.Point]int{}
	for y, line := range strings.Split(util.ReadAll(), "\n") {
		for x := range line {
			grid[image.Pt(x, y)] = util.Atoi(line[x : x+1])
		}
	}

	fmt.Println(minHeat(grid, 1, 3))
	fmt.Println(minHeat(grid, 4, 10))
}

func minHeat(grid map[image.Point]int, minStraight int, maxTurn int) int {
	b := util.Bounds(maps.Keys(grid))
	pq := util.PQ{
		&util.Item{crucible{image.Pt(0, 0), util.E, 0}, 0},
		&util.Item{crucible{image.Pt(0, 0), util.S, 0}, 0},
	}
	seen := map[crucible]int{
		{image.Pt(0, 0), util.E, 1}: 0,
		{image.Pt(0, 0), util.S, 1}: 0,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(crucible)
		if curr.p == b.Max.Sub(image.Pt(1, 1)) {
			return seen[curr]
		}
		var nexts []crucible
		if curr.n < minStraight {
			if _, ok := grid[curr.p.Add(curr.d.PointR().Mul(4-curr.n))]; ok {
				nexts = []crucible{{curr.p.Add(curr.d.PointR()), curr.d, curr.n + 1}}
			}
		} else {
			nexts = []crucible{
				{curr.p.Add(curr.d.PointR()), curr.d, curr.n + 1},
				{curr.p.Add(curr.d.Next().PointR()), curr.d.Next(), 1},
				{curr.p.Add(curr.d.Prev().PointR()), curr.d.Prev(), 1},
			}
		}
		for _, next := range nexts {
			if next.n > maxTurn {
				continue
			}
			if _, ok := grid[next.p]; !ok {
				continue
			}
			nextheat := seen[curr] + grid[next.p]
			if heat, ok := seen[next]; !ok || heat > nextheat {
				heap.Push(&pq, &util.Item{next, -nextheat})
				seen[next] = nextheat
			}
		}
	}
	return -1
}

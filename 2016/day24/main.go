package main

import (
	"container/heap"
	"fmt"
	"image"
	"math"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	grid := make(map[image.Point]struct{})
	poi := make(map[int]image.Point)
	y := 0
	s := util.ScanAll()
	for s.Scan() {
		line := s.Text()
		for x := range line {
			switch line[x] {
			case '.':
				grid[image.Pt(x, y)] = struct{}{}
			case '#':
			default:
				grid[image.Pt(x, y)] = struct{}{}
				poi[int(line[x]-'0')] = image.Pt(x, y)
			}
		}
		y++
	}
	paths := shortestPaths(grid, poi)
	fmt.Println(fewestSteps(paths, 0, map[int]struct{}{}, false))
	fmt.Println(fewestSteps(paths, 0, map[int]struct{}{}, true))
}

func shortestPaths(grid map[image.Point]struct{}, poi map[int]image.Point) map[int]map[int]int {
	paths := make(map[int]map[int]int)
	for i, p := range poi {
		paths[i] = make(map[int]int)
		depth := map[image.Point]int{p: 0}
		pq := util.PQ{&util.Item{p, 0}}
		heap.Init(&pq)
		for pq.Len() > 0 {
			curr := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
			currdepth := depth[curr] + 1
			for _, n := range util.Neighbors4 {
				next := curr.Add(n)
				if _, ok := grid[next]; !ok {
					continue
				}
				if d, ok := depth[next]; !ok || currdepth < d {
					depth[next] = currdepth
					heap.Push(&pq, &util.Item{next, -currdepth})
				}
			}
		}
		for j, q := range poi {
			if i == j {
				continue
			}
			paths[i][j] = depth[q]
		}
	}
	return paths
}

func fewestSteps(paths map[int]map[int]int, start int, seen map[int]struct{}, ret bool) int {
	seen[start] = struct{}{}
	var steps []int
	for next := range paths[start] {
		if _, ok := seen[next]; ok {
			continue
		}
		newseen := make(map[int]struct{})
		for k := range seen {
			newseen[k] = struct{}{}
		}
		newseen[next] = struct{}{}
		steps = append(steps, fewestSteps(paths, next, newseen, ret)+paths[start][next])
	}
	if len(steps) == 0 {
		if ret {
			return paths[start][0]
		}
		return 0
	}
	return min(steps)
}

func min(a []int) int {
	min := math.MaxInt16
	for _, x := range a {
		if x < min {
			min = x
		}
	}
	return min
}

package main

import (
	"container/heap"
	"fmt"
	"image"
	"math/bits"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	favnum := util.Atoi(util.ReadAll())
	fmt.Println(findDepth(image.Pt(1, 1), image.Pt(31, 39), favnum))
	fmt.Println(expand(image.Pt(1, 1), 50, favnum))
}

func findDepth(start, end image.Point, favnum int) int {
	depth := map[image.Point]int{start: 0}
	pq := util.PQ{&util.Item{start, 0}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*util.Item)
		p := i.Obj.(image.Point)
		if p == end {
			break
		}
		currdepth := depth[p] + 1
		for _, n := range util.Neighbors4 {
			next := p.Add(n)
			if next.X < 0 || next.Y < 0 {
				continue
			}
			if d, ok := depth[next]; isOpen(next.X, next.Y, favnum) && (!ok || currdepth < d) {
				depth[next] = currdepth
				heap.Push(&pq, &util.Item{next, -currdepth})
			}
		}
	}
	return depth[end]
}

func isOpen(x, y, favnum int) bool {
	return bits.OnesCount(uint(x*x+3*x+2*x*y+y+y*y+favnum))%2 == 0
}

func expand(start image.Point, steps, favnum int) int {
	visited := map[image.Point]struct{}{start: {}}
	boundary := []image.Point{start}
	for i := 0; i < steps; i++ {
		for _, p := range boundary {
			for _, n := range util.Neighbors4 {
				next := p.Add(n)
				if _, ok := visited[next]; ok || next.X < 0 || next.Y < 0 || !isOpen(next.X, next.Y, favnum) {
					continue
				}
				visited[next] = struct{}{}
				boundary = append(boundary, next)
			}
		}
	}
	return len(visited)
}

package main

import (
	"container/heap"
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	grid := map[image.Point]byte{}
	s := util.ScanAll()
	var start, end image.Point
	var as []image.Point
	y := 0
	for s.Scan() {
		for x, b := range s.Bytes() {
			p := image.Pt(x, y)
			grid[p] = b
			if b == 'S' {
				start = p
			} else if b == 'E' {
				end = p
			} else if b == 'a' {
				as = append(as, p)
			}
		}
		y++
	}
	grid[start], grid[end] = 'a', 'z'

	dists := djikstra(grid, end)

	l := dists[start]
	fmt.Println(l)

	for _, a := range as {
		if d, ok := dists[a]; ok {
			l = util.Min(l, d)
		}
	}
	fmt.Println(l)
}

func djikstra(grid map[image.Point]byte, end image.Point) map[image.Point]int {
	pq := util.PQ{&util.Item{end, 0}}
	dist := map[image.Point]int{end: 0}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
		for _, n := range util.Neighbors4 {
			next := curr.Add(n)
			if _, ok := grid[next]; !ok {
				continue
			}
			if int(grid[curr])-int(grid[next]) > 1 {
				continue
			}
			nextdist := dist[curr] + 1
			if d, ok := dist[next]; !ok || nextdist < d {
				dist[next] = nextdist
				heap.Push(&pq, &util.Item{next, nextdist})
			}
		}
	}
	return dist
}

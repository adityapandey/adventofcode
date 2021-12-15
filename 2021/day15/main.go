package main

import (
	"container/heap"
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	cave := map[image.Point]int{}
	lines := strings.Split(util.ReadAll(), "\n")
	for y, line := range lines {
		for x := range line {
			cave[image.Pt(x, y)] = int(line[x] - '0')
		}
	}
	w, h := len(lines[0]), len(lines)
	fmt.Println(minRisk(cave, image.Pt(0, 0), image.Pt(w-1, h-1)))

	for x := 0; x < 5*w; x++ {
		for y := 0; y < 5*h; y++ {
			if x < w && y < h {
				continue
			}
			cave[image.Pt(x, y)] = ((cave[image.Pt(x%w, y%h)] - 1 + x/w + y/h) % 9) + 1
		}
	}
	fmt.Println(minRisk(cave, image.Pt(0, 0), image.Pt(5*w-1, 5*h-1)))
}

func minRisk(cave map[image.Point]int, start, end image.Point) int {
	bounds := image.Rectangle{start, end.Add(image.Pt(1, 1))}
	pq := util.PQ{&util.Item{start, 0}}
	distmap := map[image.Point]int{start: 0}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
		currdist := distmap[curr]
		if curr == end {
			return currdist
		}
		for _, n := range util.Neighbors4 {
			next := curr.Add(n)
			if !next.In(bounds) {
				continue
			}
			nextdist := currdist + cave[next]
			if d, ok := distmap[next]; !ok || nextdist < d {
				distmap[next] = nextdist
				heap.Push(&pq, &util.Item{next, -nextdist})
			}
		}
	}
	panic("no path found")
}

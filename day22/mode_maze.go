package main

import (
	"container/heap"
	"fmt"
	"os"
)

type P struct{ x, y int }

var erosionMap = make(map[P]int)

func erosionLevel(p P, target P, depth int) int {
	if e, ok := erosionMap[p]; ok {
		return e
	}
	var g int
	if p == (P{0, 0}) || p == target {
		g = 0
	} else if p.x == 0 {
		g = p.y * 48271
	} else if p.y == 0 {
		g = p.x * 16807
	} else {
		g = erosionLevel(P{p.x - 1, p.y}, target, depth) * erosionLevel(P{p.x, p.y - 1}, target, depth)
	}
	e := (g + depth) % 20183
	erosionMap[p] = e
	return e
}

type RegionType int

const (
	Rocky RegionType = iota
	Wet
	Narrow
)

func regionType(p P, target P, depth int) RegionType {
	return RegionType(erosionLevel(p, target, depth) % 3)
}

type Tool int

const (
	Neither Tool = iota
	Torch
	Climber
)

type Node struct {
	p    P
	tool Tool
}

type Item struct {
	n        Node
	distance int
	index    int
	parent   *Item
}

type Q []*Item

func (q Q) Len() int { return len(q) }

func (q Q) Less(i, j int) bool { return q[i].distance < q[j].distance }

func (q Q) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *Q) Push(x interface{}) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)
}

func (q *Q) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1
	*q = old[:n-1]
	return item
}

func neighbors(i Node, target P, depth int) []Node {
	n := []Node{
		{P{i.p.x - 1, i.p.y}, i.tool},
		{P{i.p.x + 1, i.p.y}, i.tool},
		{P{i.p.x, i.p.y - 1}, i.tool},
		{P{i.p.x, i.p.y + 1}, i.tool},
	}
	for i := 0; i < len(n); i++ {
		if n[i].p.x < 0 || n[i].p.y < 0 || int(regionType(n[i].p, target, depth)) == int(n[i].tool) {
			n = append(n[:i], n[i+1:]...)
			i--
		}
	}
	return n
}

func main() {
	var depth int
	var target P
	fmt.Fscanf(os.Stdin, "depth: %d\n", &depth)
	fmt.Fscanf(os.Stdin, "target: %d,%d", &target.x, &target.y)

	// Part 1
	riskLevel := 0
	for x := 0; x <= target.x; x++ {
		for y := 0; y <= target.y; y++ {
			riskLevel += int(regionType(P{x, y}, target, depth))
		}
	}
	fmt.Println(riskLevel)

	// Part 2
	distanceMap := make(map[Node]int)
	var q Q
	heap.Init(&q)
	heap.Push(&q, &Item{Node{P{0, 0}, Torch}, 0, 0, nil})
	distanceMap[Node{P{0, 0}, Torch}] = 0
	done := make(map[Node]struct{})
	const BigNum = 9999
	for q.Len() > 0 {
		start := heap.Pop(&q).(*Item)
		if start.n.p == target && start.n.tool == Torch {
			fmt.Println(start.distance)
			break
		}
		for _, n := range neighbors(start.n, target, depth) {
			if _, ok := done[n]; ok {
				continue
			}
			var d int
			var ok bool
			if d, ok = distanceMap[n]; !ok {
				d = BigNum
			}
			if start.distance+1 < d {
				d = start.distance + 1
				heap.Push(&q, &Item{n: n, distance: d})
				distanceMap[n] = d
			}
		}
		n := start.n
		n.tool = Tool((3 - int(regionType(n.p, target, depth))) - int(n.tool))
		if _, ok := done[n]; !ok {
			var d int
			var ok bool
			if d, ok = distanceMap[n]; !ok {
				d = BigNum
			}
			if start.distance+7 < d {
				d = start.distance + 7
				heap.Push(&q, &Item{n: n, distance: d})
				distanceMap[n] = d
			}
		}
		done[start.n] = struct{}{}
	}
}

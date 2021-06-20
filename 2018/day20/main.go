package main

import (
	"container/heap"
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

var nodeAt = make(map[image.Point]*node)

type node struct {
	p         image.Point
	neighbors []*node
}

func (n *node) addAndMove(dir byte) *node {
	p := n.p.Add(util.DirFromByte(dir).Point())
	if _, ok := nodeAt[p]; !ok {
		nodeAt[p] = &node{p: p}
	}
	n.neighbors = append(n.neighbors, nodeAt[p])
	nodeAt[p].neighbors = append(nodeAt[p].neighbors, n)
	return nodeAt[p]
}

type stack []*node

func (s *stack) Top() *node { return (*s)[len(*s)-1] }

func (s *stack) Pop() *node {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *stack) Push(n *node) { *s = append(*s, n) }

func main() {
	regex := util.ReadAll()

	var origin node
	nodeAt[image.Pt(0, 0)] = &origin

	s := &stack{}
	curr := &origin
	for i := 1; i < len(regex)-1; i++ {
		switch regex[i] {
		case 'N', 'E', 'W', 'S':
			curr = curr.addAndMove(regex[i])
		case '(':
			s.Push(curr)
		case '|':
			curr = s.Top()
		case ')':
			curr = s.Pop()
		}
	}

	dist := map[image.Point]int{image.Pt(0, 0): 0}
	pq := util.PQ{&util.Item{image.Pt(0, 0), 0}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
		currdist := dist[curr] + 1
		for _, n := range nodeAt[curr].neighbors {
			if d, ok := dist[n.p]; !ok || currdist < d {
				dist[n.p] = currdist
				heap.Push(&pq, &util.Item{n.p, -currdist})
			}
		}
	}

	// Part 1
	var maxDist int
	for _, d := range dist {
		if d > maxDist {
			maxDist = d
		}
	}
	fmt.Println(maxDist)

	// Part 2
	var n int
	for _, d := range dist {
		if d >= 1000 {
			n++
		}
	}
	fmt.Println(n)
}

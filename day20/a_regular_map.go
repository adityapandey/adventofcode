package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type P struct{ x, y int }

var m = make(map[P]*Node)

type Node struct {
	P
	neighbors []*Node
}

func (n *Node) AddNeighbor(dir byte) *Node {
	p := n.P
	switch dir {
	case 'N':
		p.y--
	case 'E':
		p.x++
	case 'W':
		p.x--
	case 'S':
		p.y++
	}
	if _, ok := m[p]; !ok {
		m[p] = &Node{P: p}
	}
	n.neighbors = append(n.neighbors, m[p])
	m[p].neighbors = append(m[p].neighbors, n)
	return m[p]
}

func (n *Node) String() string {
	var s strings.Builder
	fmt.Fprintf(&s, "[(%d, %d) neighbors: ", n.x, n.y)
	for _, neighbor := range n.neighbors {
		fmt.Fprintf(&s, "(%d, %d) ", neighbor.x, neighbor.y)
	}
	return s.String()
}

type Stack []*Node

func (s *Stack) Top() *Node {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() *Node {
	r := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *Stack) Push(n *Node) {
	*s = append(*s, n)
}

func main() {
	var regex string
	fmt.Fscanln(os.Stdin, &regex)

	var origin Node
	origin.P = P{0, 0}
	m[P{0, 0}] = &origin

	l := len(regex)
	s := &Stack{}
	cur := &origin
	for i := 1; i < l-1; i++ {
		switch regex[i] {
		case 'N', 'E', 'W', 'S':
			cur = cur.AddNeighbor(regex[i])
		case '(':
			s.Push(cur)
		case '|':
			cur = s.Top()
		case ')':
			cur = s.Pop()
		}
	}

	distance := make(map[P]int)
	done := make(map[P]struct{})
	next := []P{origin.P}
	distance[origin.P] = 0

	for len(next) > 0 {
		sort.Slice(next, func(i, j int) bool { return distance[next[i]] < distance[next[j]] })
		start := next[0]
		minDist := distance[start] + 1

		for _, n := range m[start].neighbors {
			if _, ok := done[n.P]; ok {
				continue
			}
			if d, ok := distance[n.P]; !ok || d > minDist {
				distance[n.P] = minDist
				next = append(next, n.P)
			}
		}

		next = next[1:]
		done[start] = struct{}{}
	}

	// Part 1
	maxDist := 0
	for _, d := range distance {
		if d > maxDist {
			maxDist = d
		}
	}
	fmt.Println(maxDist)

	// Part 2
	c := 0
	for _, d := range distance {
		if d >= 1000 {
			c++
		}
	}
	fmt.Println(c)
}

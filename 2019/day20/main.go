package main

import (
	"container/heap"
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type grid struct {
	m           map[image.Point]byte
	portals     map[image.Point]image.Point
	portalDelta map[image.Point]int
	start, end  image.Point
}

func main() {
	g := makegrid(util.ReadAll())

	fmt.Println(path(g))

	fmt.Println(pathrecurse(g))
}

func makegrid(input string) grid {
	g := grid{
		m:           make(map[image.Point]byte),
		portals:     make(map[image.Point]image.Point),
		portalDelta: make(map[image.Point]int),
	}
	sp := strings.Split(input, "\n")
	width, height := len(sp[0]), len(sp)
	for y, line := range sp {
		for x := 0; x < len(line); x++ {
			g.m[image.Pt(x, y)] = line[x]
		}
		y++
	}
	portals := make(map[string]image.Point)
	for p, b := range g.m {
		if b >= 'A' && b <= 'Z' {
			var pt image.Point
			var name string
			for _, n := range util.Neighbors4 {
				pt = p.Add(n)
				if g.m[pt] == '.' {
					d := util.DirFromPointR(n)
					switch d {
					case util.N, util.W:
						name = string([]byte{b, g.m[p.Add(d.Reverse().PointR())]})
					case util.S, util.E:
						name = string([]byte{g.m[p.Add(d.Reverse().PointR())], b})
					}
					break
				}
			}
			if name == "" {
				continue
			}
			switch name {
			case "AA":
				g.start = pt
			case "ZZ":
				g.end = pt
			}
			if _, ok := portals[name]; !ok {
				portals[name] = pt
			} else {
				g.portals[portals[name]] = pt
				g.portals[pt] = portals[name]
				if pt.X == 2 || pt.X == width-3 || pt.Y == 2 || pt.Y == height-3 {
					g.portalDelta[pt] = -1
				} else {
					g.portalDelta[pt] = 1
				}
				g.portalDelta[portals[name]] = -g.portalDelta[pt]
			}
		}
	}
	return g
}

func path(g grid) int {
	pq := util.PQ{&util.Item{g.start, 0}}
	heap.Init(&pq)
	distmap := map[image.Point]int{g.start: 0}
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
		currdist := distmap[curr]
		if curr == g.end {
			return currdist
		}
		nextdist := currdist + 1
		for _, n := range util.Neighbors4 {
			next := curr.Add(n)
			if g.m[next] != '.' {
				continue
			}
			if _, ok := distmap[next]; !ok || nextdist < distmap[next] {
				distmap[next] = nextdist
				heap.Push(&pq, &util.Item{next, -nextdist})
			}
		}
		if next, ok := g.portals[curr]; ok {
			if _, ok := distmap[next]; !ok || nextdist < distmap[next] {
				distmap[next] = nextdist
				heap.Push(&pq, &util.Item{next, -nextdist})
			}
		}
	}
	panic("No path found")
}

type point struct {
	p image.Point
	l int
}

func pathrecurse(g grid) int {
	start, goal := point{g.start, 0}, point{g.end, 0}
	pq := util.PQ{&util.Item{start, 0}}
	heap.Init(&pq)
	distmap := map[point]int{start: 0}
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(point)
		currdist := distmap[curr]
		if curr == goal {
			return currdist
		}
		nextdist := currdist + 1
		for _, n := range util.Neighbors4 {
			next := curr
			next.p = curr.p.Add(n)
			if g.m[next.p] != '.' {
				continue
			}
			if next.l == 0 {
				if g.portalDelta[next.p] == -1 {
					continue
				}
			} else {
				if next.p == g.start || next.p == g.end {
					continue
				}
			}
			if _, ok := distmap[next]; !ok || nextdist < distmap[next] {
				distmap[next] = nextdist
				heap.Push(&pq, &util.Item{next, -nextdist})
			}
		}
		if _, ok := g.portals[curr.p]; ok {
			next := curr
			next.p = g.portals[curr.p]
			next.l += g.portalDelta[curr.p]
			if _, ok := distmap[next]; !ok || nextdist < distmap[next] {
				distmap[next] = nextdist
				heap.Push(&pq, &util.Item{next, -nextdist})
			}
		}
	}
	panic("No path found")
}

package main

import (
	"container/heap"
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

type pathfinder struct {
	m       *machine.Machine
	grid    map[image.Point]byte
	in, out chan int
	dirmap  map[util.Dir]int
	p       image.Point
	oxygen  image.Point
}

func newPathfinder(program []int) *pathfinder {
	pf := &pathfinder{
		grid:   make(map[image.Point]byte),
		in:     make(chan int),
		out:    make(chan int),
		dirmap: map[util.Dir]int{util.N: 1, util.S: 2, util.W: 3, util.E: 4},
	}
	pf.grid[pf.p] = '.'
	pf.m = machine.New(program, pf.in, pf.out)
	go pf.m.Run()
	return pf
}

func (pf *pathfinder) tryMove(dir util.Dir) bool {
	pf.in <- pf.dirmap[dir]
	next := pf.p.Add(dir.Point())
	switch <-pf.out {
	case 0:
		pf.grid[next] = '#'
		return false
	case 1:
		pf.grid[next] = '.'
	case 2:
		pf.grid[next] = 'O'
		pf.oxygen = next
	}
	pf.p = next
	return true
}

func (pf *pathfinder) explore() {
	for len(pf.open()) > 0 {
		if _, ok := pf.open()[pf.p]; !ok {
			min := math.MaxInt32
			var next image.Point
			for to := range pf.open() {
				dist := util.Manhattan(pf.p, to)
				if dist < min {
					min, next = dist, to
				}
			}
			minpath := pf.shortestPath(pf.p, next)
			for _, m := range minpath {
				if !pf.tryMove(m) {
					panic("bad path")
				}
			}
		}
		for {
			var d util.Dir
			for _, n := range util.Neighbors4 {
				if _, ok := pf.grid[pf.p.Add(n)]; !ok {
					d = util.DirFromPoint(n)
					break
				}
			}
			if !pf.tryMove(d) {
				break
			}
		}
	}
}

func (pf *pathfinder) open() map[image.Point]struct{} {
	ps := make(map[image.Point]struct{})
	for p, b := range pf.grid {
		if b == '#' {
			continue
		}
		for _, n := range util.Neighbors4 {
			if _, ok := pf.grid[p.Add(n)]; !ok {
				ps[p] = struct{}{}
				break
			}
		}
	}
	return ps
}

func (pf *pathfinder) shortestPath(from, to image.Point) []util.Dir {
	pq := util.PQ{&util.Item{from, 0}}
	pathmap := map[image.Point][]util.Dir{from: {}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
		currpath := pathmap[curr]
		if curr == to {
			break
		}
		for _, n := range util.Neighbors4 {
			next := curr.Add(n)
			if b, ok := pf.grid[next]; !ok || b == '#' {
				continue
			}
			if path, ok := pathmap[next]; !ok || len(path) > 1+len(currpath) {
				nextpath := make([]util.Dir, 1+len(currpath))
				copy(nextpath, currpath)
				nextpath[len(nextpath)-1] = util.DirFromPoint(n)
				heap.Push(&pq, &util.Item{next, -len(nextpath)})
				pathmap[next] = nextpath
			}
		}
	}
	path, ok := pathmap[to]
	if !ok {
		panic("no path")
	}
	return path
}

func (pf *pathfinder) longestPath(from image.Point) int {
	pq := util.PQ{&util.Item{from, 0}}
	distmap := map[image.Point]int{from: 0}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
		currdist := distmap[curr]
		for _, n := range util.Neighbors4 {
			next := curr.Add(n)
			if b, ok := pf.grid[next]; !ok || b == '#' {
				continue
			}
			if dist, ok := distmap[next]; !ok || dist > 1+currdist {
				nextdist := 1 + currdist
				heap.Push(&pq, &util.Item{next, -nextdist})
				distmap[next] = nextdist
			}
		}
	}
	var max int
	for _, d := range distmap {
		max = util.Max(max, d)
	}
	return max
}

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}

	pf := newPathfinder(program)
	pf.explore()
	fmt.Println(len(pf.shortestPath(image.Pt(0, 0), pf.oxygen)))
	fmt.Println(pf.longestPath(pf.oxygen))
}

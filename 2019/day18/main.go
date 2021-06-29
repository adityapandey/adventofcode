package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"image"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	grid, keys, start := parse(util.ReadAll())

	fmt.Println(optimalPath(grid, keys, []image.Point{start}))

	grid[start] = '#'
	for _, n := range util.Neighbors4 {
		grid[start.Add(n)] = '#'
	}
	var starts []image.Point
	for _, n := range []image.Point{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
		grid[start.Add(n)] = '@'
		starts = append(starts, start.Add(n))
	}
	fmt.Println(optimalPath(grid, keys, starts))
}

func parse(input string) (map[image.Point]byte, map[byte]struct{}, image.Point) {
	grid := make(map[image.Point]byte)
	keys := make(map[byte]struct{})
	var start image.Point
	var y int
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		for x := 0; x < len(line); x++ {
			b, p := line[x], image.Pt(x, y)
			grid[image.Pt(x, y)] = b
			switch {
			case b >= 'a' && b <= 'z':
				keys[b] = struct{}{}
			case b == '@':
				start = p
			}
		}
		y++
	}
	return grid, keys, start
}

type step struct {
	p        []image.Point
	workerId int
	keys     map[byte]struct{}
}

func (g step) clone() step {
	gg := step{
		p:        make([]image.Point, len(g.p)),
		workerId: g.workerId,
		keys:     make(map[byte]struct{}),
	}
	copy(gg.p, g.p)
	for k := range g.keys {
		gg.keys[k] = struct{}{}
	}
	return gg
}

func (g step) String() string {
	var ks []byte
	for k := range g.keys {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool { return ks[i] < ks[j] })
	return fmt.Sprintf("%v_%v_%v", g.p, g.workerId, string(ks))
}

func optimalPath(grid map[image.Point]byte, keys map[byte]struct{}, start []image.Point) int {
	pq := util.PQ{}
	heap.Init(&pq)
	dist := make(map[string]int)

	s := step{start, 0, keys}
	for i := 0; i < len(start); i++ {
		ss := s.clone()
		ss.workerId = i
		heap.Push(&pq, &util.Item{ss, 0})
		dist[ss.String()] = 0
	}

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(step)
		currdist := dist[curr.String()]
		if len(curr.keys) == 0 {
			return currdist
		}
		nextdist := 1 + currdist
		for _, n := range util.Neighbors4 {
			b, ok := grid[curr.p[curr.workerId].Add(n)]
			if !ok || b == '#' {
				continue
			}
			if b >= 'A' && b <= 'Z' {
				if _, ok := curr.keys[b+('a'-'A')]; ok {
					continue
				}
			}
			next := curr.clone()
			next.p[next.workerId] = next.p[next.workerId].Add(n)
			foundNewKey := false
			if b >= 'a' && b <= 'z' {
				if _, ok := next.keys[b]; ok {
					foundNewKey = true
					delete(next.keys, b)
				}
			}
			for i := 0; i < len(next.p); i++ {
				if curr.workerId != i && !foundNewKey {
					continue
				}
				worker := next.clone()
				worker.workerId = i
				if d, ok := dist[worker.String()]; !ok || nextdist < d {
					dist[worker.String()] = nextdist
					heap.Push(&pq, &util.Item{worker, -nextdist})
				}
			}
		}
	}
	panic("No viable path")
}

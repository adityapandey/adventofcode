package main

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	orbits := make(map[string]string)
	s := util.ScanAll()
	for s.Scan() {
		sp := strings.Split(s.Text(), ")")
		orbits[sp[1]] = sp[0]
	}

	var sum int
	for o := range orbits {
		sum += totalOrbits(o, orbits)
	}
	fmt.Println(sum)

	orbitedBy := make(map[string]map[string]struct{})
	for k, v := range orbits {
		if _, ok := orbitedBy[v]; !ok {
			orbitedBy[v] = make(map[string]struct{})
		}
		orbitedBy[v][k] = struct{}{}
	}
	pq := util.PQ{&util.Item{"YOU", 0}}
	dist := map[string]int{"YOU": 0}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(string)
		currdist := dist[curr]
		if curr == "SAN" {
			fmt.Println(currdist - 2)
			break
		}
		if o, ok := orbits[curr]; ok {
			if d, ok := dist[o]; !ok || d < 1+currdist {
				d = 1 + currdist
				dist[o] = d
				heap.Push(&pq, &util.Item{o, -d})
			}
		}
		for o := range orbitedBy[curr] {
			if d, ok := dist[o]; !ok || d < 1+currdist {
				d = 1 + currdist
				dist[o] = d
				heap.Push(&pq, &util.Item{o, -d})
			}
		}
	}
}

func totalOrbits(s string, orbits map[string]string) int {
	var totalOrbitsInternal func(string, map[string]int, map[string]string) int
	totalOrbitsInternal = func(s string, m map[string]int, orbits map[string]string) int {
		if n, ok := m[s]; ok {
			return n
		}
		if _, ok := orbits[s]; !ok {
			m[s] = 0
			return 0
		}
		n := 1 + totalOrbitsInternal(orbits[s], m, orbits)
		m[s] = n
		return n
	}
	m := make(map[string]int)
	return totalOrbitsInternal(s, m, orbits)
}

package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

type edge struct {
	from, to string // invariant: from < to
}

func main() {
	nodes := map[string][]string{}
	s := util.ScanAll()
	for s.Scan() {
		sp := strings.Split(s.Text(), ": ")
		for _, f := range strings.Fields(sp[1]) {
			nodes[sp[0]] = append(nodes[sp[0]], f)
			nodes[f] = append(nodes[f], sp[0])
		}
	}
	nRandomPairs := 50
	for i := 0; i < 3; i++ {
		remove2Edge(nodes, maxEdge(nodes, nRandomPairs))
	}
	c := clusterSize(nodes, maps.Keys(nodes)[0])
	fmt.Println(c * (len(nodes) - c))
}

func clusterSize(nodes map[string][]string, start string) int {
	seen := map[string]struct{}{}
	q := []string{start}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, next := range nodes[curr] {
			if _, ok := seen[next]; ok {
				continue
			}
			q = append(q, next)
			seen[next] = struct{}{}
		}
	}
	return len(seen)
}

// Most traversed edge between n random pairs of nodes.
func maxEdge(nodes map[string][]string, n int) edge {
	edgeCount := map[edge]int{}
	for _, pair := range randomNodePairs(nodes, n) {
		seen := map[string]struct{}{}
		q := []string{pair[0]}
		for len(q) > 0 {
			curr := q[0]
			if curr == pair[1] {
				break
			}
			q = q[1:]
			for _, next := range nodes[curr] {
				if _, ok := seen[next]; ok {
					continue
				}
				q = append(q, next)
				seen[next] = struct{}{}
				var e edge
				if curr < next {
					e = edge{curr, next}
				} else {
					e = edge{next, curr}
				}
				edgeCount[e]++
			}
		}
	}
	var max int
	var maxEdge edge
	for edge, c := range edgeCount {
		if c > max {
			max, maxEdge = c, edge
		}
	}
	return maxEdge
}

func randomNodePairs(nodes map[string][]string, n int) [][2]string {
	res := make([][2]string, n)
	nodeNames := maps.Keys(nodes)
	for i := 0; i < n; i++ {
		res = append(res, [2]string{
			nodeNames[rand.Intn(len(nodeNames))],
			nodeNames[rand.Intn(len(nodeNames))],
		})
	}
	return res
}

func remove2Edge(nodes map[string][]string, e edge) {
	nodes[e.from] = slices.DeleteFunc(nodes[e.from], func(s string) bool { return s == e.to })
	nodes[e.to] = slices.DeleteFunc(nodes[e.to], func(s string) bool { return s == e.from })
}

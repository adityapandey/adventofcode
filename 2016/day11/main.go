package main

import (
	"container/heap"
	"fmt"
	"hash/fnv"
	"regexp"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`a ([a-z]+)(-compatible)* (generator|microchip)`)
var hash = fnv.New64()

type floor []int

func (f floor) legal() bool {
	g := make(map[int]struct{})
	for _, obj := range f {
		if obj > 0 {
			g[obj] = struct{}{}
		}
	}
	if len(g) == 0 {
		return true
	}
	// At least one generator exists
	for _, obj := range f {
		if obj < 0 {
			if _, ok := g[-obj]; !ok {
				return false
			}
		}
	}
	return true
}

func (f floor) move(to floor, i, j int) (floor, floor) {
	fcopy := make(floor, len(f))
	copy(fcopy, f)
	tocopy := make(floor, len(to))
	copy(tocopy, to)
	m := []int{f[i]}
	if i == j { // Actually moving only one object
		return append(fcopy[:i], fcopy[i+1:]...), append(tocopy, m...)
	}
	m = append(m, f[j])
	return append(fcopy[:i], append(fcopy[i+1:j], fcopy[j+1:]...)...), append(tocopy, m...)
}

type state struct {
	e      int
	floors [4]floor
}

func (s state) hash() uint64 {
	hash.Reset()
	fmt.Fprintf(hash, "%v", s)
	return hash.Sum64()
}

func (s state) sortfloors() {
	for i := 0; i < 4; i++ {
		sort.Ints(s.floors[i])
	}
}

func (s state) end() state {
	var end state
	end.e = 3
	for i := 0; i < 4; i++ {
		end.floors[3] = append(end.floors[3], s.floors[i]...)
	}
	end.sortfloors()
	return end
}

func (s state) next() []state {
	var states []state
	for _, delta := range []int{+1, -1} {
		if s.e+delta < 0 || s.e+delta > 3 {
			continue
		}
		from, to := s.floors[s.e], s.floors[s.e+delta]
		for i := 0; i < len(from); i++ {
			for j := i; j < len(from); j++ {
				f, t := from.move(to, i, j)
				if f.legal() && t.legal() {
					var next state
					next.e = s.e + delta
					for k := 0; k < 4; k++ {
						if k == s.e {
							next.floors[k] = append(next.floors[k], f...)
						} else if k == s.e+delta {
							next.floors[k] = append(next.floors[k], t...)
						} else {
							next.floors[k] = append(next.floors[k], s.floors[k]...)
						}
					}
					next.sortfloors()
					states = append(states, next)
				}
			}
		}
	}
	return states
}

func main() {
	var start state
	start.e = 0
	objs := make(map[string]int)
	input := strings.Split(util.ReadAll(), "\n")
	for i := range input {
		matches := re.FindAllStringSubmatch(input[i], -1)
		for j := range matches {
			name := matches[j][1]
			if _, ok := objs[name]; !ok {
				objs[name] = len(objs) + 1
			}
			obj := objs[name]
			if matches[j][3] == "microchip" {
				obj = -obj
			}
			start.floors[i] = append(start.floors[i], obj)
		}
	}
	start.sortfloors()
	fmt.Println(findDepth(start))

	start.floors[0] = append(start.floors[0],
		len(objs)+1,
		-(len(objs) + 1),
		len(objs)+2,
		-(len(objs) + 2),
	)
	start.sortfloors()
	fmt.Println(findDepth(start))
}

func findDepth(start state) int {
	end := start.end().hash()
	depth := map[uint64]int{start.hash(): 0}
	remaining := util.PQ{&util.Item{start, 0}}
	heap.Init(&remaining)
	for remaining.Len() > 0 {
		i := heap.Pop(&remaining).(*util.Item)
		if i.Obj.(state).hash() == end {
			break
		}
		currdepth := depth[i.Obj.(state).hash()] + 1
		for _, n := range i.Obj.(state).next() {
			h := n.hash()
			if d, ok := depth[h]; !ok || currdepth < d {
				depth[h] = currdepth
				// heuristic from https://reddit.com/r/adventofcode/comments/5hoia9/2016_day_11_solutions/db1zbu0/
				heap.Push(&remaining, &util.Item{n, -currdepth + 6*len(n.floors[3])})
			}
		}
	}
	return depth[end]
}

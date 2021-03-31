package main

import (
	"container/heap"
	"fmt"
	"image"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

type unit struct {
	kind   byte
	p      image.Point
	ap, hp int
}

func main() {
	input := make(map[image.Point]unit)
	walls := make(map[image.Point]struct{})
	var w, h, y int
	s := util.ScanAll()
	for s.Scan() {
		b := s.Bytes()
		w = len(b)
		for x := range b {
			switch b[x] {
			case 'G', 'E':
				u := unit{b[x], image.Pt(x, y), 3, 200}
				input[image.Pt(x, y)] = u
			case '#':
				walls[image.Pt(x, y)] = struct{}{}
			}
		}
		y++
	}
	h = y

	// Part 1
	units := make(map[image.Point]unit)
	for p, u := range input {
		units[p] = u
	}
	var round int
	for step(units, walls, w, h) {
		round++
	}
	var sum int
	for _, u := range units {
		if u.hp > 0 {
			sum += u.hp
		}
	}
	fmt.Println(round * sum)

	// Part 2
	var nElves, nRemaining int
	for _, u := range input {
		if u.kind == 'E' {
			nElves++
		}
	}
	for ap := 3; nElves > nRemaining; ap++ {
		units = make(map[image.Point]unit)
		for p, u := range input {
			if u.kind == 'E' {
				u.ap = ap
			}
			units[p] = u
		}
		round, nRemaining = 0, 0
		for step(units, walls, w, h) {
			round++
		}
		for _, u := range units {
			if u.kind == 'E' {
				nRemaining++
			}
		}
	}
	sum = 0
	for _, u := range units {
		if u.hp > 0 {
			sum += u.hp
		}
	}
	fmt.Println(round * sum)
}

func step(units map[image.Point]unit, walls map[image.Point]struct{}, w, h int) bool {
	orderedUnits := readingOrder(units)
	for i := range orderedUnits {
		currUnit, ok := units[orderedUnits[i]]
		if !ok || currUnit.hp <= 0 {
			// Dead, nothing to do here.
			continue
		}
		// Find shortest distances.
		currDist := distances(units, walls, currUnit.p, w, h)
		// Candidates in range.
		var cs []image.Point
		for _, u := range units {
			if u.kind != currUnit.kind {
				for _, n := range util.Neighbors4 {
					c := u.p.Add(n)
					if c.In(image.Rect(0, 0, w, h)) {
						cs = append(cs, c)
					}
				}
			}
		}
		// No candidates! We are done.
		if len(cs) == 0 {
			return false
		}
		// Reachable candidates
		for j := 0; j < len(cs); j++ {
			if _, ok := currDist[cs[j]]; !ok {
				cs = append(cs[:j], cs[j+1:]...)
				j--
			}
		}
		// Reading order.
		sort.Slice(cs, func(i, j int) bool {
			if currDist[cs[i]] == currDist[cs[j]] {
				if cs[i].Y == cs[j].Y {
					return cs[i].X < cs[j].X
				}
				return cs[i].Y < cs[j].Y
			}
			return currDist[cs[i]] < currDist[cs[j]]
		})
		// Move if needed.
		if len(cs) > 0 && cs[0] != currUnit.p {
			target := cs[0]
			// Paths to target
			tDist := distances(units, walls, target, w, h)
			var moves []image.Point
			for _, n := range util.Neighbors4 {
				m := currUnit.p.Add(n)
				if d, ok := tDist[m]; ok && d == currDist[target]-1 && m.In(image.Rect(0, 0, w, h)) {
					moves = append(moves, m)
				}
			}
			sort.Slice(moves, func(i, j int) bool {
				if moves[i].Y == moves[j].Y {
					return moves[i].X < moves[j].X
				}
				return moves[i].Y < moves[j].Y
			})
			delete(units, currUnit.p)
			currUnit.p = moves[0]
			units[currUnit.p] = currUnit
		}
		// Attack.
		var enemies []image.Point
		for _, n := range util.Neighbors4 {
			e := currUnit.p.Add(n)
			if u, ok := units[e]; ok && u.kind != currUnit.kind {
				enemies = append(enemies, u.p)
			}
		}
		sort.Slice(enemies, func(i, j int) bool {
			if units[enemies[i]].hp == units[enemies[j]].hp {
				if units[enemies[i]].p.Y == units[enemies[j]].p.Y {
					return units[enemies[i]].p.X < units[enemies[j]].p.X
				}
				return units[enemies[i]].p.Y < units[enemies[j]].p.Y
			}
			return units[enemies[i]].hp < units[enemies[j]].hp
		})
		if len(enemies) > 0 {
			target := units[enemies[0]]
			target.hp -= currUnit.ap
			units[target.p] = target
			if target.hp <= 0 {
				delete(units, target.p)
			}
		}
	}
	return true
}

func readingOrder(units map[image.Point]unit) []image.Point {
	us := make([]image.Point, 0, len(units))
	for _, u := range units {
		us = append(us, u.p)
	}
	sort.Slice(us, func(i, j int) bool {
		if us[i].Y == us[j].Y {
			return us[i].X < us[j].X
		}
		return us[i].Y < us[j].Y
	})
	return us
}

func distances(units map[image.Point]unit, walls map[image.Point]struct{}, p image.Point, w, h int) map[image.Point]int {
	dist := map[image.Point]int{p: 0}
	pq := util.PQ{&util.Item{p, 0}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
		currdist := dist[curr] + 1
		for _, n := range util.Neighbors4 {
			next := curr.Add(n)
			if !next.In(image.Rect(0, 0, w, h)) {
				continue
			}
			if _, ok := units[next]; ok {
				continue
			}
			if _, ok := walls[next]; ok {
				continue
			}
			if d, ok := dist[next]; !ok || currdist < d {
				dist[next] = currdist
				heap.Push(&pq, &util.Item{next, -currdist})
			}
		}
	}
	return dist
}

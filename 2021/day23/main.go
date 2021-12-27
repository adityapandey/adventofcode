package main

import (
	"container/heap"
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

const (
	burrowTop = `#############
#...........#
###.#.#.#.###`
	burrowMid    = `  #.#.#.#.#`
	burrowBottom = `  #########`
)

type state struct {
	pos   string
	depth int
}

func parse(g map[image.Point]byte) state {
	maxY := 0
	for p := range g {
		maxY = util.Max(maxY, p.Y)
	}
	depth := maxY - 2
	b := make([]byte, 11+4*depth)
	for i := 0; i < 11; i++ {
		b[i] = g[image.Pt(i+1, 1)]
	}
	for d := 0; d < depth; d++ {
		for i := 11 + 4*d; i < 15+4*d; i++ {
			b[i] = g[image.Pt(2*(i-(11+4*d))+3, 2+d)]
		}
	}
	return state{string(b), depth}
}

func parseString(input string) state {
	g := map[image.Point]byte{}
	for y, line := range strings.Split(input, "\n") {
		for x := range line {
			g[image.Pt(x, y)] = line[x]
		}
	}
	return parse(g)
}

func (s state) grid() map[image.Point]byte {
	g := map[image.Point]byte{}
	burrow := []string{burrowTop}
	for d := 1; d < s.depth; d++ {
		burrow = append(burrow, burrowMid)
	}
	burrow = append(burrow, burrowBottom)
	for y, line := range strings.Split(strings.Join(burrow, "\n"), "\n") {
		for x := range line {
			g[image.Pt(x, y)] = line[x]
		}
	}
	for i := 0; i < 11; i++ {
		g[image.Pt(i+1, 1)] = s.pos[i]
	}
	for d := 0; d < s.depth; d++ {
		for i := 11 + 4*d; i < 15+4*d; i++ {
			g[image.Pt(2*(i-(11+4*d))+3, 2+d)] = s.pos[i]
		}
	}
	return g
}

func (s state) end() state {
	pos := "..........."
	for d := 0; d < s.depth; d++ {
		pos += "ABCD"
	}
	return state{pos, s.depth}
}

func (s state) String() string {
	g := s.grid()
	var sb strings.Builder
	for y := 0; y < 3+s.depth; y++ {
		for x := 0; x < 13; x++ {
			if b, ok := g[image.Pt(x, y)]; ok {
				fmt.Fprintf(&sb, "%c", b)
			} else {
				fmt.Fprintf(&sb, " ")
			}
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}

func main() {
	s := parseString(util.ReadAll())
	fmt.Println(minEnergy(s))
	s2 := state{s.pos[:15] + "DCBADBAC" + s.pos[15:], 4}
	fmt.Println(minEnergy(s2))
}

func minEnergy(s state) int {
	end := s.end()
	pq := util.PQ{&util.Item{s, 0}}
	energy := map[state]int{s: 0}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(state)
		currEnergy := energy[curr]
		if curr == end {
			return currEnergy
		}
		g := curr.grid()
		// Hallway
		for x := 1; x < 12; x++ {
			p := image.Pt(x, 1)
			if g[p] == '.' {
				continue
			}
			if ok, next, moves := moveToRoom(curr, p, g); ok {
				nextEnergy := currEnergy + expense(g[p])*moves
				if e, ok := energy[next]; !ok || nextEnergy < e {
					energy[next] = nextEnergy
					heap.Push(&pq, &util.Item{next, -nextEnergy})
				}
			}
		}
		// Rooms
		for _, x := range []int{3, 5, 7, 9} {
			p := image.Pt(x, 1)
			for g[p] == '.' {
				p.Y++
			}
			if p.Y < 2 || p.Y > 1+curr.depth {
				continue
			}
			if ok, next, moves := moveToRoom(curr, p, g); ok {
				nextEnergy := currEnergy + expense(g[p])*moves
				if e, ok := energy[next]; !ok || nextEnergy < e {
					energy[next] = nextEnergy
					heap.Push(&pq, &util.Item{next, -nextEnergy})
				}
			}
			for _, hallwayX := range []int{1, 2, 4, 6, 8, 10, 11} {
				if ok, moves := path(p, image.Pt(hallwayX, 1), g); ok {
					nextGrid := curr.grid()
					nextGrid[image.Pt(hallwayX, 1)] = g[p]
					nextGrid[p] = '.'
					next := parse(nextGrid)
					nextEnergy := currEnergy + expense(g[p])*moves
					if e, ok := energy[next]; !ok || nextEnergy < e {
						energy[next] = nextEnergy
						heap.Push(&pq, &util.Item{next, -nextEnergy})
					}
				}
			}
		}
	}
	return -1
}

func moveToRoom(curr state, p image.Point, g map[image.Point]byte) (bool, state, int) {
	roomX := int(g[p]-'A')*2 + 3
	valid := true
	var top int
	for y := 2; y < 2+curr.depth; y++ {
		if g[image.Pt(roomX, y)] != '.' && g[image.Pt(roomX, y)] != g[p] {
			valid = false
			break
		}
		if g[image.Pt(roomX, y)] == '.' {
			top = y
		}
	}
	if !valid {
		return false, state{}, -1
	}

	if ok, moves := path(p, image.Pt(roomX, top), g); ok {
		nextGrid := curr.grid()
		nextGrid[image.Pt(roomX, top)] = g[p]
		nextGrid[p] = '.'
		next := parse(nextGrid)
		return true, next, moves
	}
	return false, state{}, -1
}

func path(from, to image.Point, g map[image.Point]byte) (bool, int) {
	if from == to {
		return false, -1
	}
	if g[to] != '.' {
		return false, -1
	}
	p := from
	var moves int
	for ; p.Y > 1; p.Y-- {
		moves++
		if p != from && g[p] != '.' {
			return false, -1
		}
	}
	step := 1
	if to.X < p.X {
		step = -1
	}
	for ; p.X != to.X; p.X += step {
		moves++
		if p != from && g[p] != '.' {
			return false, -1
		}
	}
	for ; p.Y != to.Y; p.Y++ {
		moves++
		if p != from && g[p] != '.' {
			return false, -1
		}
	}
	return true, moves
}

func expense(b byte) int {
	switch b {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	}
	panic("unknown")
}

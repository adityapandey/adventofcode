package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

const BigNum = 999

type P struct {
	X, Y int
}

type Unit struct {
	P
	B  byte
	AP int
	HP int
}

func MakeUnit(x, y int, b byte) *Unit {
	return &Unit{P{x, y}, b, 3, 200}
}

func (u *Unit) IsAlive() bool {
	return u.HP > 0
}

func makeGrid(input [][]byte) [][]byte {
	grid := make([][]byte, len(input[0]))
	for i := range grid {
		grid[i] = make([]byte, len(input))
	}
	for y := range input {
		for x := range input[y] {
			grid[x][y] = input[y][x]
		}
	}
	return grid
}

type Item struct {
	P
	dist  int
	index int
}

type Q []*Item

func (q Q) Len() int           { return len(q) }
func (q Q) Less(i, j int) bool { return q[i].dist < q[j].dist }
func (q Q) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}
func (q *Q) Push(x interface{}) {
	l := len(*q)
	item := x.(*Item)
	item.index = l
	*q = append(*q, item)
}
func (q *Q) Pop() interface{} {
	old := *q
	l := len(old)
	item := old[l-1]
	item.index = -1
	*q = old[:l-1]
	return item
}

type World struct {
	grid       [][]byte
	lenX, lenY int
	Units      []*Unit
	unit_at    map[P]*Unit
}

func MakeWorld(input [][]byte) *World {
	w := new(World)
	w.grid = makeGrid(input)
	w.lenX = len(w.grid)
	w.lenY = len(w.grid[0])
	w.unit_at = make(map[P]*Unit)

	for x := range w.grid {
		for y := range w.grid[x] {
			if w.grid[x][y] == 'G' || w.grid[x][y] == 'E' {
				w.AddUnit(x, y, w.grid[x][y])
			}
		}
	}

	return w
}

func (w *World) AddUnit(x int, y int, b byte) {
	u := MakeUnit(x, y, b)
	w.Units = append(w.Units, u)
	w.unit_at[P{x, y}] = u
}

func (w *World) Neighbors(p P) []P {
	candidates := []P{
		{p.X - 1, p.Y},
		{p.X, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X, p.Y + 1},
	}
	for i := 0; i < len(candidates); i++ {
		n := candidates[i]
		if n.X < 0 || n.X > w.lenX-1 || n.Y < 0 || n.Y > w.lenY-1 {
			candidates = append(candidates[:i], candidates[i+1:]...)
			i--
		}
	}
	return candidates
}

func (w *World) distanceGrid(p P) [][]int {
	g := make([][]int, w.lenX)
	for i := range g {
		g[i] = make([]int, w.lenY)
		for j := range g[i] {
			g[i][j] = BigNum
		}
	}

	g[p.X][p.Y] = 0
	done := make(map[P]struct{})
	var q Q
	heap.Init(&q)
	heap.Push(&q, &Item{P: p, dist: 0})

	for q.Len() > 0 {
		start := heap.Pop(&q).(*Item)
		dist := start.dist + 1

		for _, n := range w.Neighbors(start.P) {
			if w.grid[n.X][n.Y] != '.' {
				continue
			}
			if _, ok := done[n]; ok {
				continue
			}
			if dist < g[n.X][n.Y] {
				g[n.X][n.Y] = dist
				heap.Push(&q, &Item{P: n, dist: dist})
			}
		}
		done[start.P] = struct{}{}
	}
	return g
}

func (w *World) Advance() bool {
	// Reading order of units
	sort.Slice(w.Units, func(i, j int) bool {
		if w.Units[i].Y == w.Units[j].Y {
			return w.Units[i].X < w.Units[j].X
		}
		return w.Units[i].Y < w.Units[j].Y
	})

	for _, u := range w.Units {
		if !u.IsAlive() {
			// Unit is dead. Nothing to do.
			continue
		}
		// Shortest distances
		distGrid := w.distanceGrid(u.P)
		//Current positions of units
		var goblins []P
		var elves []P
		for x := range w.grid {
			for y := range w.grid[x] {
				switch w.grid[x][y] {
				case 'G':
					goblins = append(goblins, P{x, y})
				case 'E':
					elves = append(elves, P{x, y})
				}
			}
		}
		// Candidates for movement
		var candidates []P
		if u.B == 'G' {
			for _, e := range elves {
				candidates = append(candidates, w.Neighbors(e)...)
			}
		} else {
			for _, g := range goblins {
				candidates = append(candidates, w.Neighbors(g)...)
			}
		}
		// No more candidates. We are done.
		if len(candidates) == 0 {
			return false
		}
		// Closest candidates
		sort.Slice(candidates, func(i, j int) bool {
			return distGrid[candidates[i].X][candidates[i].Y] < distGrid[candidates[j].X][candidates[j].Y]
		})
		minDist := distGrid[candidates[0].X][candidates[0].Y]
		// Reachable candidates
		if minDist > 0 && minDist < BigNum {
			for i, c := range candidates {
				if distGrid[c.X][c.Y] > minDist {
					candidates = candidates[:i]
					break
				}
			}
			// First reading order candidate
			sort.Slice(candidates, func(i, j int) bool {
				if candidates[i].Y == candidates[j].Y {
					return candidates[i].X < candidates[j].X
				}
				return candidates[i].Y < candidates[j].Y
			})
			t := candidates[0] // Target point
			// Paths to target
			distGrid := w.distanceGrid(t)
			var moves []P
			for _, n := range w.Neighbors(u.P) {
				if distGrid[n.X][n.Y] == minDist-1 {
					moves = append(moves, n)
				}
			}

			sort.Slice(moves, func(i, j int) bool {
				if moves[i].Y == moves[j].Y {
					return moves[i].X < moves[j].X
				}
				return moves[i].Y < moves[j].Y
			})
			m := moves[0]
			w.grid[m.X][m.Y], w.grid[u.X][u.Y] = w.grid[u.X][u.Y], w.grid[m.X][m.Y]
			delete(w.unit_at, u.P)
			u.P = m
			w.unit_at[u.P] = u
		}

		// Attack begins
		var enemies []*Unit
		for _, n := range w.Neighbors(u.P) {
			if _, ok := w.unit_at[n]; !ok {
				continue
			}
			if u.B != w.unit_at[n].B {
				enemies = append(enemies, w.unit_at[n])
			}
		}
		sort.Slice(enemies, func(i, j int) bool { return enemies[i].HP < enemies[j].HP })
		if len(enemies) > 0 {
			minHP := enemies[0].HP
			for i := 0; i < len(enemies); i++ {
				if enemies[i].HP > minHP {
					enemies = enemies[:i]
					break
				}
			}
			sort.Slice(enemies, func(i, j int) bool {
				if enemies[i].Y == enemies[j].Y {
					return enemies[i].X < enemies[j].X
				}
				return enemies[i].Y < enemies[j].Y
			})
			target := enemies[0]
			target.HP -= u.AP
			if !target.IsAlive() {
				delete(w.unit_at, target.P)
				w.grid[target.X][target.Y] = '.'
			}
		}
	}
	return true
}

func main() {
	var input [][]byte
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input = append(input, s.Bytes())
	}

	// Part 1
	world := MakeWorld(input)
	var round int
	for world.Advance() {
		round++
	}
	sumHP := 0
	for _, u := range world.Units {
		if u.HP > 0 {
			sumHP += u.HP
		}
	}
	fmt.Println(round, sumHP, round*sumHP)

	// Part 2
	for AP := 3; ; AP++ {
		elfLoss := false
		world = MakeWorld(input)
		for _, u := range world.Units {
			if u.B == 'E' {
				u.AP = AP
			}
		}
		var round int
		for world.Advance() && !elfLoss {
			round++
			for _, u := range world.Units {
				if !u.IsAlive() && u.B == 'E' {
					elfLoss = true
					break
				}
			}
		}
		if !elfLoss {
			sumHP := 0
			for _, u := range world.Units {
				if u.HP > 0 {
					sumHP += u.HP
				}
			}
			fmt.Println(round, sumHP, round*sumHP)
			break
		}
	}
}

func printGrid(grid [][]byte) {
	for y := -1; y < len(grid); y++ {
		for x := -1; x < len(grid[0]); x++ {
			if y == -1 {
				fmt.Printf("%4d", x)
			} else if x == -1 {
				fmt.Printf("%4d", y)
			} else {
				fmt.Printf("%4c", grid[x][y])
			}
		}
		fmt.Println()
	}
}

func printDistGrid(grid [][]int) {
	for y := -1; y < len(grid); y++ {
		for x := -1; x < len(grid[0]); x++ {
			if y == -1 {
				fmt.Printf("%4d", x)
			} else if x == -1 {
				fmt.Printf("%4d", y)
			} else {
				if grid[x][y] == BigNum {
					fmt.Printf("....")
				} else {
					fmt.Printf("%4d", grid[x][y])
				}
			}
		}
		fmt.Println()
	}
}

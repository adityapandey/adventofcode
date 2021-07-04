package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

type point struct {
	pt    image.Point
	level int
}

func (p point) neighbors() []point {
	var ps []point
	for _, n := range util.Neighbors4 {
		next := p.pt.Add(n)
		switch {
		case next.X < 0 || next.Y < 0 || next.X > 4 || next.Y > 4:
			ps = append(ps, point{n.Add(image.Pt(2, 2)), p.level - 1})
		case next == image.Pt(2, 2):
			for i := 0; i < 5; i++ {
				ps = append(ps, point{image.Pt(i*n.Y*n.Y-2*n.X*n.X*(n.X-1), i*n.X*n.X-2*n.Y*n.Y*(n.Y-1)), p.level + 1})
			}
		default:
			ps = append(ps, point{next, p.level})
		}
	}
	return ps
}

func main() {
	var grid [5][5]byte
	grid2 := make(map[point]struct{})
	s := util.ScanAll()
	var y int
	for s.Scan() {
		line := s.Text()
		for x := 0; x < len(line); x++ {
			grid[x][y] = line[x]
			if line[x] == '#' {
				grid2[point{image.Pt(x, y), 0}] = struct{}{}
			}
		}
		y++
	}

	seen := make(map[[5][5]byte]struct{})
	for _, ok := seen[grid]; !ok; _, ok = seen[grid] {
		seen[grid] = struct{}{}
		grid = iterate(grid)
	}
	fmt.Println(biodiversity(grid))

	for i := 0; i < 200; i++ {
		grid2 = iterate2(grid2)
	}
	fmt.Println(len(grid2))
}

func iterate(grid [5][5]byte) [5][5]byte {
	g := grid
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			var alive int
			for _, n := range util.Neighbors4 {
				xx, yy := x+n.X, y+n.Y
				if xx < 0 || xx > 4 || yy < 0 || yy > 4 {
					continue
				}
				if grid[xx][yy] == '#' {
					alive++
				}
			}
			if grid[x][y] == '#' && alive != 1 {
				g[x][y] = '.'
			}
			if grid[x][y] == '.' && (alive == 1 || alive == 2) {
				g[x][y] = '#'
			}
		}
	}
	return g
}

func biodiversity(grid [5][5]byte) int {
	sum, pow := 0, 1
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if grid[x][y] == '#' {
				sum += pow
			}
			pow *= 2
		}
	}
	return sum
}

func iterate2(grid map[point]struct{}) map[point]struct{} {
	g := make(map[point]struct{})
	var ps []point
	for p := range grid {
		ps = append(ps, p)
		ps = append(ps, p.neighbors()...)
		g[p] = grid[p]
	}
	for _, p := range ps {
		var alive int
		for _, n := range p.neighbors() {
			if _, ok := grid[n]; ok {
				alive++
			}
		}
		_, ok := grid[p]

		if ok && alive != 1 {
			delete(g, p)
		}
		if !ok && (alive == 1 || alive == 2) {
			g[p] = struct{}{}
		}
	}
	return g
}

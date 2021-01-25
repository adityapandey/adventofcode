// https://adventofcode.com/2020/day/11
package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

var dirs = []image.Point{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1}}

func main() {
	s := bufio.NewScanner(os.Stdin)
	y := 0
	grid := make(map[image.Point]byte)
	for s.Scan() {
		t := s.Text()
		for x := 0; x < len(t); x++ {
			grid[image.Pt(x, y)] = t[x]
		}
		y++
	}

	// Part 1
	part1(grid)

	// Part 2
	part2(grid)
}

func part1(grid map[image.Point]byte) {
	neighbourFunc := func(p, d image.Point) image.Point {
		return p.Add(d)
	}
	var next map[image.Point]byte
	for {
		next = iterate(grid, 4, neighbourFunc)
		if equal(grid, next) {
			break
		}
		grid = next
	}
	occupied := 0
	for _, v := range grid {
		if v == '#' {
			occupied++
		}
	}
	fmt.Println(occupied)
}

func iterate(grid map[image.Point]byte, maxOccupancy int, neighbourFunc func(p, d image.Point) image.Point) map[image.Point]byte {
	next := make(map[image.Point]byte)
	for p := range grid {
		next[p] = grid[p]
		if grid[p] == '.' {
			continue
		}
		occupied := 0
		for _, d := range dirs {
			n := neighbourFunc(p, d)
			if grid[image.Pt(n.X, n.Y)] == '#' {
				occupied++
			}
		}
		if grid[p] == 'L' && occupied == 0 {
			next[p] = '#'
		} else if grid[p] == '#' && occupied >= maxOccupancy {
			next[p] = 'L'
		}
	}
	return next
}

func equal(a, b map[image.Point]byte) bool {
	for p := range a {
		if a[p] != b[p] {
			return false
		}
	}
	return true
}

func part2(grid map[image.Point]byte) {
	neighbourFunc := func(p, d image.Point) image.Point {
		p = p.Add(d)
		for grid[p] == '.' {
			p = p.Add(d)
		}
		return p
	}
	var next map[image.Point]byte
	for {
		next = iterate(grid, 5, neighbourFunc)
		if equal(grid, next) {
			break
		}
		grid = next
	}
	occupied := 0
	for _, v := range grid {
		if v == '#' {
			occupied++
		}
	}
	fmt.Println(occupied)
}

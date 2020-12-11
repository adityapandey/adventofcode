// https://adventofcode.com/2020/day/11
package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

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
	next := iterate(grid)
	for !equal(next, grid) {
		grid, next = next, iterate(next)
	}
	occupied := 0
	for _, v := range grid {
		if v == '#' {
			occupied++
		}
	}
	fmt.Println(occupied)
}

func iterate(grid map[image.Point]byte) map[image.Point]byte {
	next := make(map[image.Point]byte)
	for p := range grid {
		next[p] = grid[p]
		if grid[p] == '.' {
			continue
		}
		occupied := 0
		for _, n := range neighbours(p) {
			if grid[image.Pt(n.X, n.Y)] == '#' {
				occupied++
			}
		}
		if grid[p] == 'L' && occupied == 0 {
			next[p] = '#'
		} else if grid[p] == '#' && occupied >= 4 {
			next[p] = 'L'
		}
	}
	return next
}

func neighbours(p image.Point) []image.Point {
	dirs := []image.Point{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}
	var n []image.Point
	for _, d := range dirs {
		n = append(n, p.Add(d))
	}
	return n
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
	next := iterate2(grid)
	for !equal(next, grid) {
		grid, next = next, iterate2(next)
	}
	occupied := 0
	for _, v := range grid {
		if v == '#' {
			occupied++
		}
	}
	fmt.Println(occupied)
}

func iterate2(grid map[image.Point]byte) map[image.Point]byte {
	next := make(map[image.Point]byte)
	for p := range grid {
		next[p] = grid[p]
		if grid[p] == '.' {
			continue
		}
		occupied := 0
		for _, n := range neighbours2(p, grid) {
			if grid[image.Pt(n.X, n.Y)] == '#' {
				occupied++
			}
		}
		if grid[p] == 'L' && occupied == 0 {
			next[p] = '#'
		} else if grid[p] == '#' && occupied >= 5 {
			next[p] = 'L'
		}
	}
	return next
}

func neighbours2(p image.Point, grid map[image.Point]byte) []image.Point {
	dirs := []image.Point{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}
	var n []image.Point
	for _, d := range dirs {
		curr := p.Add(d)
		for grid[curr] == '.' {
			curr = curr.Add(d)
		}
		n = append(n, curr)

	}
	return n
}

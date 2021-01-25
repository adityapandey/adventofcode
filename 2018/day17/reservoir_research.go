package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type P struct{ x, y int }

type World struct {
	grid       [][]byte
	source     P
	minY, maxY int
}

func NewWorld(xranges []XRange, yranges []YRange) *World {
	minY, maxY := math.MaxUint32, 0
	minX, maxX := math.MaxUint32, 0

	for _, r := range xranges {
		if r.y < minY {
			minY = r.y
		}
		if r.y > maxY {
			maxY = r.y
		}
		if r.xmin < minX {
			minX = r.xmin
		}
		if r.xmax > maxX {
			maxX = r.xmax
		}
	}
	for _, r := range yranges {
		if r.x < minX {
			minX = r.x
		}
		if r.x > maxX {
			maxX = r.x
		}
		if r.ymin < minY {
			minY = r.ymin
		}
		if r.ymax > maxY {
			maxY = r.ymax
		}
	}
	var w World
	// Leave 1 col on both sides.
	// Origin is now (minX-1)
	w.grid = make([][]byte, maxX+2-(minX-1))
	for i := range w.grid {
		w.grid[i] = make([]byte, maxY+1)
		for j := range w.grid[i] {
			w.grid[i][j] = '.'
		}
	}

	for _, r := range xranges {
		for x := r.xmin; x <= r.xmax; x++ {
			w.grid[x-(minX-1)][r.y] = '#'
		}
	}

	for _, r := range yranges {
		for y := r.ymin; y <= r.ymax; y++ {
			w.grid[r.x-(minX-1)][y] = '#'
		}
	}

	w.source = P{500 - (minX - 1), 0}
	w.grid[w.source.x][w.source.y] = '+'
	w.minY, w.maxY = minY, maxY

	return &w
}

func (w *World) Print() {
	for y := 0; y < len(w.grid[0]); y++ {
		for x := 0; x < len(w.grid); x++ {
			fmt.Printf("%c", w.grid[x][y])
		}
		fmt.Println()
	}
}

func (w *World) startFlow(p P) {
	for p.y <= w.maxY && w.grid[p.x][p.y] != '#' && w.grid[p.x][p.y] != '~' {
		w.grid[p.x][p.y] = '|'
		p.y++
	}
	if p.y > w.maxY {
		return
	}
	p.y--
	leftEdge, leftWall := w.goLeft(p)
	rightEdge, rightWall := w.goRight(p)

	if leftWall && rightWall {
		for x := leftEdge.x + 1; x < rightEdge.x; x++ {
			w.grid[x][leftEdge.y] = '~'
		}
	} else {
		if !leftWall {
			w.startFlow(leftEdge)
		}
		if !rightWall {
			w.startFlow(rightEdge)
		}
		return
	}
	p.y--
	w.startFlow(p)
}

func (w *World) goLeft(p P) (P, bool) {
	for (w.grid[p.x][p.y+1] == '#' || w.grid[p.x][p.y+1] == '~') && (w.grid[p.x][p.y] != '#' && w.grid[p.x][p.y] != '~') {
		w.grid[p.x][p.y] = '|'
		p.x--
	}
	return p, w.grid[p.x][p.y] == '#' || w.grid[p.x][p.y] == '~'
}

func (w *World) goRight(p P) (P, bool) {
	for (w.grid[p.x][p.y+1] == '#' || w.grid[p.x][p.y+1] == '~') && (w.grid[p.x][p.y] != '#' && w.grid[p.x][p.y] != '~') {
		w.grid[p.x][p.y] = '|'
		p.x++
	}
	return p, w.grid[p.x][p.y] == '#' || w.grid[p.x][p.y] == '~'
}

func (w *World) Fill() {
	p := w.source
	w.startFlow(p)
}

type XRange struct {
	y, xmin, xmax int
}

type YRange struct {
	x, ymin, ymax int
}

func main() {
	var xranges []XRange
	var yranges []YRange
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text()[0] == 'x' {
			var r YRange
			fmt.Sscanf(s.Text(), "x=%d, y=%d..%d", &r.x, &r.ymin, &r.ymax)
			yranges = append(yranges, r)
		} else {
			var r XRange
			fmt.Sscanf(s.Text(), "y=%d, x=%d..%d", &r.y, &r.xmin, &r.xmax)
			xranges = append(xranges, r)
		}
	}
	w := NewWorld(xranges, yranges)
	w.Fill()
	// w.Print()

	var sumFlow int
	var sumRest int
	for x := 0; x < len(w.grid); x++ {
		for y := w.minY; y <= w.maxY; y++ {
			switch w.grid[x][y] {
			case '|':
				sumFlow++
			case '~':
				sumRest++
			}
		}
	}

	// Part 1
	fmt.Println(sumRest + sumFlow)

	// Part 2
	fmt.Println(sumRest)
}

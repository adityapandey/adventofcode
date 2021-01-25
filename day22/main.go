package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

type dir int

const (
	N dir = iota
	E
	S
	W
	NUM_DIR
)

var dirs = map[dir]image.Point{
	N: {0, -1},
	E: {1, 0},
	S: {0, 1},
	W: {-1, 0},
}

type state int

const (
	CLEAN state = iota
	WEAKENED
	INFECTED
	FLAGGED
	NUM_STATE
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	input := make(map[image.Point]state)
	var y, midx int
	for s.Scan() {
		midx = len(s.Bytes()) / 2
		for x, b := range s.Bytes() {
			if b == '#' {
				input[image.Pt(x, y)] = INFECTED
			}
		}
		y++
	}
	curr := image.Pt(midx, y/2)
	var d dir
	var infections int
	g := copymap(input)

	for i := 0; i < 10000; i++ {
		s := g[curr]
		switch s {
		case CLEAN:
			g[curr] = INFECTED
			infections++
			d = (d + 3) % NUM_DIR
		case INFECTED:
			delete(g, curr)
			d = (d + 1) % NUM_DIR
		}
		curr = curr.Add(dirs[d])
	}
	fmt.Println(infections)

	curr = image.Pt(midx, y/2)
	d = N
	infections = 0
	g = copymap(input)
	for i := 0; i < 10000000; i++ {
		s := g[curr]
		switch s {
		case CLEAN:
			d = (d + 3) % NUM_DIR
		case WEAKENED:
			infections++
		case INFECTED:
			d = (d + 1) % NUM_DIR
		case FLAGGED:
			d = (d + 2) % NUM_DIR
		}
		g[curr] = (g[curr] + 1) % NUM_STATE
		curr = curr.Add(dirs[d])
	}
	fmt.Println(infections)
}

func copymap(input map[image.Point]state) map[image.Point]state {
	g := make(map[image.Point]state)
	for k := range input {
		g[k] = input[k]
	}
	return g
}

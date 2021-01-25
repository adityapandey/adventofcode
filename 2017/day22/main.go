package main

import (
	"bufio"
	"fmt"
	"image"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

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
	var d util.Dir
	var infections int
	g := copymap(input)

	for i := 0; i < 10000; i++ {
		s := g[curr]
		switch s {
		case CLEAN:
			g[curr] = INFECTED
			infections++
			d = d.Prev()
		case INFECTED:
			delete(g, curr)
			d = d.Next()
		}
		curr = curr.Add(d.PointR())
	}
	fmt.Println(infections)

	curr = image.Pt(midx, y/2)
	d = util.N
	infections = 0
	g = copymap(input)
	for i := 0; i < 10000000; i++ {
		s := g[curr]
		switch s {
		case CLEAN:
			d = d.Prev()
		case WEAKENED:
			infections++
		case INFECTED:
			d = d.Next()
		case FLAGGED:
			d = d.Reverse()
		}
		g[curr] = (g[curr] + 1) % NUM_STATE
		curr = curr.Add(d.PointR())
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

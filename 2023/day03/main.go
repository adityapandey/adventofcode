package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type part struct {
	xmin, xmax int
	y          int
	n          int
}

func (p part) valid(grid map[image.Point]byte) bool {
	for x := p.xmin; x <= p.xmax; x++ {
		for _, n := range util.Neighbors8 {
			c, ok := grid[n.Add(image.Pt(x, p.y))]
			if ok && c != '.' && (c < '0' || c > '9') {
				return true
			}
		}
	}
	return false
}

func main() {
	grid := map[image.Point]byte{}
	var parts []part
	var curr *part
	for y, line := range strings.Split(util.ReadAll(), "\n") {
		if curr != nil {
			parts = append(parts, *curr)
			curr = nil
		}
		for x := 0; x < len(line); x++ {
			c := line[x]
			grid[image.Pt(x, y)] = c
			if c >= '0' && c <= '9' {
				if curr == nil {
					curr = new(part)
					curr.y = y
					curr.xmin = x
					curr.xmax = x
					curr.n = int(c - '0')
				} else {
					curr.n *= 10
					curr.n += int(c - '0')
					curr.xmax = x
				}
			} else if curr != nil {
				parts = append(parts, *curr)
				curr = nil
			}
		}
	}

	partsGrid := map[image.Point]int{}
	sum := 0
	for i, p := range parts {
		for x := p.xmin; x <= p.xmax; x++ {
			partsGrid[image.Pt(x, p.y)] = i
		}
		if p.valid(grid) {
			sum += p.n
		}
	}
	fmt.Println(sum)

	sum2 := 0
	for p, c := range grid {
		if c == '*' {
			neighborParts := map[int]struct{}{}
			for _, n := range util.Neighbors8 {
				if i, ok := partsGrid[n.Add(p)]; ok {
					neighborParts[i] = struct{}{}
				}
			}
			if len(neighborParts) == 2 {
				prod := 1
				for i := range neighborParts {
					prod *= parts[i].n
				}
				sum2 += prod
			}
		}
	}
	fmt.Println(sum2)
}

package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type state struct {
	pos  image.Point
	step int
}

func main() {
	grid := map[image.Point]byte{}
	lines := strings.Split(util.ReadAll(), "\n")
	h, w := len(lines), len(lines[0])
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] != '.' {
				grid[image.Pt(x, y)] = line[x]
			}
		}
	}

	entrance, exit := image.Pt(1, 0), image.Pt(w-2, h-1)
	firstCrossing := steps(grid, w, h, entrance, exit, 0)
	secondCrossing := steps(grid, w, h, exit, entrance, firstCrossing)
	thirdCrossing := steps(grid, w, h, entrance, exit, secondCrossing)

	fmt.Println(firstCrossing)
	fmt.Println(thirdCrossing)
}

func steps(grid map[image.Point]byte, w int, h int, start image.Point, end image.Point, initialStep int) int {
	q := []state{{start, initialStep}}
	seen := map[state]struct{}{}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.pos == end {
			return curr.step
		}
	loop:
		for _, n := range append(util.Neighbors4, image.Pt(0, 0)) {
			nextstate := state{curr.pos.Add(n), curr.step + 1}
			if _, ok := seen[nextstate]; ok {
				continue
			}
			if !nextstate.pos.In(image.Rect(0, 0, w, h)) {
				continue
			}
			if grid[nextstate.pos] == '#' {
				continue
			}
			if nextstate.pos.Y > 0 && nextstate.pos.Y < h-1 {
				for _, bliz := range []byte{'^', '>', 'v', '<'} {
					prev := nextstate.pos.Sub(util.DirFromByte(bliz).PointR().Mul(nextstate.step)).Mod(image.Rect(1, 1, w-1, h-1))
					if grid[prev] == bliz {
						continue loop
					}
				}
			}
			q = append(q, nextstate)
			seen[nextstate] = struct{}{}
		}
	}
	return -1
}

package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

func main() {
	grid := map[image.Point]struct{}{}
	y := 0
	for _, line := range strings.Split(util.ReadAll(), "\n") {
		for x := range line {
			if line[x] == '#' {
				grid[image.Pt(x, y)] = struct{}{}
			}
		}
		y++
	}
	for i := 0; ; i++ {
		newgrid := round(i, grid)
		if i == 9 {
			b := bounds(maps.Keys(newgrid))
			fmt.Println((b.Dx()+1)*(b.Dy()+1) - len(newgrid))
		}
		stop := true
		for elf := range grid {
			if _, ok := newgrid[elf]; !ok {
				stop = false
				break
			}
		}
		if stop {
			fmt.Println(i + 1)
			break
		}
		grid = newgrid
	}
}

func round(roundNum int, grid map[image.Point]struct{}) map[image.Point]struct{} {
	next := map[image.Point]image.Point{}
	for elf := range grid {
		var mask byte
		for i, n := range []image.Point{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}} {
			if _, ok := grid[elf.Add(n)]; ok {
				mask |= 1 << i
			}
		}
		if mask == 0 {
			next[elf] = elf
		} else {
			conds := []struct {
				cond bool
				dir  util.Dir
			}{
				{mask&(1<<7+1<<1+1<<0) == 0, util.N},
				{mask&(1<<3+1<<4+1<<5) == 0, util.S},
				{mask&(1<<5+1<<6+1<<7) == 0, util.W},
				{mask&(1<<1+1<<2+1<<3) == 0, util.E},
			}
			conds = append(conds[roundNum%4:4], conds[:roundNum%4]...)
			for _, c := range conds {
				if c.cond {
					next[elf] = elf.Add(c.dir.PointR())
					break
				}
			}
		}
		if _, ok := next[elf]; !ok {
			next[elf] = elf
		}
	}

	m := map[image.Point]int{}
	for _, n := range next {
		m[n]++
	}

	newgrid := map[image.Point]struct{}{}
	for elf := range grid {
		if m[next[elf]] == 1 {
			newgrid[next[elf]] = struct{}{}
		} else {
			newgrid[elf] = struct{}{}
		}
	}

	return newgrid
}

func bounds(p []image.Point) image.Rectangle {
	r := image.Rectangle{p[0], p[0]}
	for i := 1; i < len(p); i++ {
		r = r.Union(image.Rect(p[i-1].X, p[i-1].Y, p[i].X, p[i].Y))
	}
	return r.Bounds()
}

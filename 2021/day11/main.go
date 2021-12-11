package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type octopus struct {
	energy int
	flash  bool
}

var bounds = image.Rect(0, 0, 10, 10)

type grid map[image.Point]*octopus

func (g grid) step() int {
	var flashes int
	for p := range g {
		flashes += g.inc(p)
	}
	for p := range g {
		if g[p].flash {
			g[p].flash = false
			g[p].energy = 0
		}
	}
	return flashes
}

func (g grid) inc(p image.Point) int {
	if g[p].flash {
		return 0
	}
	var flashes int
	g[p].energy++
	if g[p].energy > 9 {
		flashes++
		g[p].flash = true
		for _, n := range util.Neighbors8 {
			nn := p.Add(n)
			if nn.In(bounds) {
				flashes += g.inc(nn)
			}
		}
	}
	return flashes
}

func (g grid) isSync() bool {
	for p := range g {
		if g[p].energy != 0 {
			return false
		}
	}
	return true
}

func main() {
	g := grid{}
	for y, line := range strings.Split(util.ReadAll(), "\n") {
		for x, n := range strings.Split(line, "") {
			g[image.Pt(x, y)] = &octopus{energy: util.Atoi(n)}
		}
	}

	var flashes, step, score100, scoreSync int
	var synced bool
	for step = 1; step <= 100 || !synced; step++ {
		flashes += g.step()
		if step == 100 {
			score100 = flashes
		}
		if g.isSync() {
			scoreSync = step
			synced = true
		}
	}
	fmt.Println(score100)
	fmt.Println(scoreSync)
}

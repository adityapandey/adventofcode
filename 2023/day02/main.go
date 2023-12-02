package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type game struct {
	id   int
	sets []set
}

func parseGame(s string) game {
	var g game
	data := strings.Split(s, ": ")
	g.id = util.Atoi(strings.Split(data[0], " ")[1])
	sets := strings.Split(data[1], "; ")
	for _, set := range sets {
		g.sets = append(g.sets, parseSet(set))
	}
	return g
}

func (g game) possible() bool {
	for _, s := range g.sets {
		for _, c := range s.cubes {
			if (c.color == "red" && c.n > 12) || (c.color == "green" && c.n > 13) || (c.color == "blue" && c.n > 14) {
				return false
			}
		}
	}
	return true
}

func (g game) power() int {
	var reds, greens, blues []int
	for _, s := range g.sets {
		for _, c := range s.cubes {
			switch c.color {
			case "red":
				reds = append(reds, c.n)
			case "green":
				greens = append(greens, c.n)
			case "blue":
				blues = append(blues, c.n)
			}
		}
	}
	return util.Max(reds...) * util.Max(greens...) * util.Max(blues...)
}

type set struct {
	cubes []cube
}

func parseSet(s string) set {
	var ss set
	cubes := strings.Split(s, ", ")
	for _, cube := range cubes {
		ss.cubes = append(ss.cubes, parseCube(cube))
	}
	return ss
}

type cube struct {
	n     int
	color string
}

func parseCube(s string) cube {
	var c cube
	sp := strings.Split(s, " ")
	c.n = util.Atoi(sp[0])
	c.color = sp[1]
	return c
}

func main() {
	var games []game
	for _, g := range strings.Split(util.ReadAll(), "\n") {
		games = append(games, parseGame(g))
	}

	var sumIds, sumPowers int
	for _, g := range games {
		if g.possible() {
			sumIds += g.id
		}
		sumPowers += g.power()
	}
	fmt.Println(sumIds)
	fmt.Println(sumPowers)
}

package main

import (
	"fmt"
	"image"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

type Turn int

const (
	LeftTurn = iota
	Straight
	RightTurn
)

type cart struct {
	p        image.Point
	d        util.Dir
	nextTurn Turn
	id       byte
}

func (c *cart) TurnLeft() {
	c.d = c.d.Prev()
}

func (c *cart) TurnRight() {
	c.d = c.d.Next()
}

func (c *cart) GetNextTurn() Turn {
	t := c.nextTurn
	c.nextTurn++
	c.nextTurn %= 3
	return t
}

func (c *cart) Step(grid map[image.Point]byte) {
	c.p = c.p.Add(c.d.PointR())
	switch grid[c.p] {
	case '/':
		if c.d == util.N || c.d == util.S {
			c.TurnRight()
		} else {
			c.TurnLeft()
		}
	case '\\':
		if c.d == util.N || c.d == util.S {
			c.TurnLeft()
		} else {
			c.TurnRight()
		}
	case '+':
		switch c.GetNextTurn() {
		case LeftTurn:
			c.TurnLeft()
		case RightTurn:
			c.TurnRight()
		}
	}
}

func main() {
	var carts []cart
	grid := make(map[image.Point]byte)
	var id byte = 'A'
	s := util.ScanAll()
	for y := 0; s.Scan(); y++ {
		b := s.Bytes()
		for x := range b {
			grid[image.Pt(x, y)] = b[x]
			switch b[x] {
			case '>', '<', '^', 'v':
				carts = append(carts, cart{p: image.Pt(x, y), d: util.DirFromByte(b[x]), id: id})
				id++
			}
		}
	}

	cartsCopy := make([]cart, len(carts))
	copy(cartsCopy, carts)
	fmt.Println(firstCrash(grid, cartsCopy))

	fmt.Println(lastCart(grid, carts))
}

func firstCrash(grid map[image.Point]byte, carts []cart) image.Point {
	for {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].p.Y == carts[j].p.Y {
				return carts[i].p.X < carts[j].p.X
			}
			return carts[i].p.Y < carts[j].p.Y
		})
		for i := range carts {
			c := &carts[i]
			c.Step(grid)
			for j := range carts {
				o := &carts[j]
				if i != j && c.p == o.p {
					return c.p
				}
			}
		}
	}
}

func lastCart(grid map[image.Point]byte, carts []cart) image.Point {
	for len(carts) > 1 {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].p.Y == carts[j].p.Y {
				return carts[i].p.X < carts[j].p.X
			}
			return carts[i].p.Y < carts[j].p.Y
		})
		for i := 0; i < len(carts); i++ {
			c := &carts[i]
			c.Step(grid)
			for j := 0; j < len(carts); j++ {
				o := &carts[j]
				if i != j && c.p == o.p {
					if i < j {
						carts = append(carts[:i], carts[i+1:]...)
						carts = append(carts[:j-1], carts[j:]...)
						i--
					} else {
						carts = append(carts[:j], carts[j+1:]...)
						carts = append(carts[:i-1], carts[i:]...)
						i -= 2
					}
					break
				}
			}
		}
	}
	return carts[0].p
}

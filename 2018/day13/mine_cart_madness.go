package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Direction int

const (
	Up Direction = iota
	Left
	Down
	Right
)

type Turn int

const (
	LeftTurn = iota
	Straight
	RightTurn
)

type Cart struct {
	x, y     int
	d        Direction
	nextTurn Turn
	id       byte
}

func (c *Cart) TurnLeft() {
	c.d++
	c.d %= 4
}

func (c *Cart) TurnRight() {
	c.d += 3
	c.d %= 4
}

func (c *Cart) Step() {
	switch c.d {
	case Up:
		c.y--
	case Down:
		c.y++
	case Left:
		c.x--
	case Right:
		c.x++
	}
}

func (c *Cart) GetNextTurn() Turn {
	t := c.nextTurn
	c.nextTurn++
	c.nextTurn %= 3
	return t
}

func (c *Cart) Byte() byte {
	switch c.d {
	case Up:
		return '^'
	case Down:
		return 'v'
	case Left:
		return '<'
	case Right:
		return '>'
	}
	log.Fatal(*c)
	return '?'
}

func main() {
	var carts []Cart
	var grid [][]byte
	var id byte = 'A'
	s := bufio.NewScanner(os.Stdin)
	for y := 0; s.Scan(); y++ {
		b := s.Bytes()
		for x := range b {
			switch b[x] {
			case '>':
				carts = append(carts, Cart{x: x, y: y, d: Right, id: id})
				id++
				b[x] = '-'
			case '<':
				carts = append(carts, Cart{x: x, y: y, d: Left, id: id})
				id++
				b[x] = '-'
			case '^':
				carts = append(carts, Cart{x: x, y: y, d: Up, id: id})
				id++
				b[x] = '|'
			case 'v':
				carts = append(carts, Cart{x: x, y: y, d: Down, id: id})
				id++
				b[x] = '|'
			}
		}
		grid = append(grid, make([]byte, len(b)))
		copy(grid[y], b)
	}

	// Part 1
	cartsCopy := make([]Cart, len(carts))
	copy(cartsCopy, carts)
	x, y := firstCrash(grid, cartsCopy)
	fmt.Printf("%d,%d\n", x, y)

	// Part 2
	x, y = lastCart(grid, carts)
	fmt.Printf("%d,%d\n", x, y)
}

func firstCrash(grid [][]byte, carts []Cart) (int, int) {
	for {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})
		for i := range carts {
			c := &carts[i]
			c.Step()
			switch grid[c.y][c.x] {
			case '/':
				if c.d == Up || c.d == Down {
					c.TurnRight()
				} else {
					c.TurnLeft()
				}
			case '\\':
				if c.d == Up || c.d == Down {
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
			for j := range carts {
				o := &carts[j]
				if o != c && c.x == o.x && c.y == o.y {
					return c.x, c.y
				}
			}
		}
	}
}

func lastCart(grid [][]byte, carts []Cart) (int, int) {
	for len(carts) > 1 {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y == carts[j].y {
				return carts[i].x < carts[j].x
			}
			return carts[i].y < carts[j].y
		})
		for i := 0; i < len(carts); i++ {
			c := &carts[i]
			c.Step()
			switch grid[c.y][c.x] {
			case '/':
				if c.d == Up || c.d == Down {
					c.TurnRight()
				} else {
					c.TurnLeft()
				}
			case '\\':
				if c.d == Up || c.d == Down {
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
			for j := 0; j < len(carts); j++ {
				o := &carts[j]
				if i != j && c.x == o.x && c.y == o.y {
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
	return carts[0].x, carts[0].y
}

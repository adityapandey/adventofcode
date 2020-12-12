// https://adventofcode.com/2020/day/12

package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

type instr struct {
	action byte
	val    int
}

var dirs = map[byte]image.Point{
	'E': image.Pt(1, 0),
	'S': image.Pt(0, 1),
	'W': image.Pt(-1, 0),
	'N': image.Pt(0, -1),
}

func main() {
	var instructions []instr
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var action byte
		var val int
		fmt.Sscanf(s.Text(), "%c%d", &action, &val)
		instructions = append(instructions, instr{action, val})
	}

	// Part 1
	dir := dirs['E']
	ship := image.Pt(0, 0)
	for _, in := range instructions {
		switch in.action {
		case 'R':
			// always a multiple of 90
			dir = clockwise(dir, in.val/90)
		case 'L':
			// always a multiple of 90
			dir = clockwise(dir, 3*in.val/90)
		case 'N', 'S', 'E', 'W':
			ship = ship.Add(dirs[in.action].Mul(in.val))
		case 'F':
			ship = ship.Add(dir.Mul(in.val))
		}
	}
	fmt.Println(abs(ship.X) + abs(ship.Y))

	// Part 2
	waypoint := image.Pt(10, -1)
	ship = image.Pt(0, 0)
	for _, in := range instructions {
		switch in.action {
		case 'R':
			// always a multiple of 90
			waypoint = clockwise(waypoint, in.val/90)
		case 'L':
			// always a multiple of 90
			waypoint = clockwise(waypoint, 3*in.val/90)
		case 'N', 'S', 'E', 'W':
			waypoint = waypoint.Add(dirs[in.action].Mul(in.val))
		case 'F':
			ship = ship.Add(waypoint.Mul(in.val))
		}
	}
	fmt.Println(abs(ship.X) + abs(ship.Y))
}

func clockwise(p image.Point, turns int) image.Point {
	for i := 0; i < turns%4; i++ {
		p = image.Pt(-p.Y, p.X)
	}
	return p
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

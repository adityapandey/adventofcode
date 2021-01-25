// https://adventofcode.com/2020/day/12
package main

import (
	"bufio"
	"fmt"
	"image"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

type instr struct {
	action byte
	val    int
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
	dir := util.E
	ship := image.Pt(0, 0)
	for _, in := range instructions {
		switch in.action {
		case 'R':
			// always a multiple of 90
			dir = dir.Add(in.val / 90)
		case 'L':
			// always a multiple of 90
			dir = dir.Add(3 * in.val / 90)
		case 'N', 'S', 'E', 'W':
			ship = ship.Add(util.DirFromByte(in.action).Point().Mul(in.val))
		case 'F':
			ship = ship.Add(dir.Point().Mul(in.val))
		}
	}
	fmt.Println(util.Manhattan(ship, image.Pt(0, 0)))

	// Part 2
	waypoint := image.Pt(10, 1)
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
			waypoint = waypoint.Add(util.DirFromByte(in.action).Point().Mul(in.val))
		case 'F':
			ship = ship.Add(waypoint.Mul(in.val))
		}
	}
	fmt.Println(util.Manhattan(ship, image.Pt(0, 0)))
}

func clockwise(p image.Point, turns int) image.Point {
	for i := 0; i < turns%4; i++ {
		p = image.Pt(p.Y, -p.X)
	}
	return p
}

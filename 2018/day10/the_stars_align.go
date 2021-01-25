package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y, vx, vy int
}

func MaxMinY(points []Point) (int, int) {
	maxY, minY := 0, math.MaxInt32
	for _, p := range points {
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}
	return maxY, minY
}

func MaxMinX(points []Point) (int, int) {
	maxX, minX := 0, math.MaxInt32
	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
	}
	return maxX, minX
}

func empty() [126][90]byte {
	var b [126][90]byte
	for x := 0; x < 126; x++ {
		for y := 0; y < 90; y++ {
			b[x][y] = '.'
		}
	}
	return b
}

func main() {
	var points []Point
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var p Point
		fmt.Sscanf(s.Text(), "position=<%d,%d> velocity=<%d,%d>", &p.x, &p.y, &p.vx, &p.vy)
		points = append(points, p)
	}

	var spreads []int

	for t := 0; ; t++ {
		max, min := MaxMinY(points)
		spreads = append(spreads, max-min)
		if t > 0 && spreads[t] > spreads[t-1] {
			for i := range points {
				points[i].x -= points[i].vx
				points[i].y -= points[i].vy
			}
			X, x := MaxMinX(points)
			Y, y := MaxMinY(points)
			screen := make([]byte, (Y-y+1)*(X-x+1))
			for i := range screen {
				screen[i] = '.'
			}
			for _, p := range points {
				screen[(p.y-y)*(X-x+1)+(p.x-x)] = '#'
			}
			// Part 1
			for i := y; i <= Y; i++ {
				for j := x; j <= X; j++ {
					fmt.Printf("%c", screen[(i-y)*(X-x+1)+(j-x)])
				}
				fmt.Println()
			}
			// Part 2
			fmt.Println(t - 1)
			break
		}
		for i := range points {
			points[i].x += points[i].vx
			points[i].y += points[i].vy
		}
	}
}

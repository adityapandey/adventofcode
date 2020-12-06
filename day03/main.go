// https://adventofcode.com/2020/day/3
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adityapandey/adventofcode2020-go/util"
)

func main() {
	m := make(map[util.Pt]bool)
	var width, height int
	s := bufio.NewScanner(os.Stdin)
	y := 0
	for s.Scan() {
		t := s.Text()
		width = len(t)
		for x := 0; x < width; x++ {
			if t[x] == '.' {
				m[util.Pt{x, y}] = false
			} else {
				m[util.Pt{x, y}] = true
			}
		}
		y++
	}
	height = y

	// Part 1
	fmt.Println(numTrees(3, 1, m, width, height))

	// Part 2
	product := 1
	for _, p := range []util.Pt{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		product *= numTrees(p.X, p.Y, m, width, height)
	}
	fmt.Println(product)
}

func numTrees(deltaX, deltaY int, m map[util.Pt]bool, width, height int) int {
	trees := 0
	x, y := 0, 0
	for y < height {
		if m[util.Pt{x, y}] {
			trees++
		}
		y += deltaY
		x = (x + deltaX) % width
	}
	return trees
}

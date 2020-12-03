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
	trees := 0
	x, y := 0, 0
	for y < height {
		if m[util.Pt{x, y}] {
			trees++
		}
		y++
		x = (x + 3) % width
	}
	fmt.Println(trees)
}

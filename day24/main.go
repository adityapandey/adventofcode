// https://adventofcode.com/2020/day/24
package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
)

var (
	e  = image.Pt(2, 0)
	se = image.Pt(1, 1)
	sw = image.Pt(-1, 1)
	w  = image.Pt(-2, 0)
	nw = image.Pt(-1, -1)
	ne = image.Pt(1, -1)
)

func main() {
	var input [][]image.Point
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input = append(input, parseDirections(s.Text()))
	}

	floor := make(map[image.Point]bool) // false = white tile
	for i := 0; i < len(input); i++ {
		tile := image.Pt(0, 0)
		for j := 0; j < len(input[i]); j++ {
			tile = tile.Add(input[i][j])
		}
		floor[tile] = !floor[tile]
	}

	// Part 1
	fmt.Println(countTrue(floor))

	// Part 2
	for i := 0; i < 100; i++ {
		floor = iterate(floor)
	}
	fmt.Println(countTrue(floor))

}

func parseDirections(s string) []image.Point {
	var ds []image.Point
	l := len(s)
	for i := 0; i < l; i++ {
		var d image.Point
		switch s[i] {
		case 'n', 's':
			i++
			switch s[i-1 : i+1] {
			case "ne":
				d = ne
			case "nw":
				d = nw
			case "se":
				d = se
			case "sw":
				d = sw
			default:
				log.Fatal("Unknown direction", s[i-1:i+1])
			}
		case 'e':
			d = e
		case 'w':
			d = w
		default:
			log.Fatal("Unknown direction", s[i])
		}
		ds = append(ds, d)
	}
	return ds
}

func countTrue(floor map[image.Point]bool) int {
	sum := 0
	for tile := range floor {
		if floor[tile] {
			sum++
		}
	}
	return sum
}

func iterate(floor map[image.Point]bool) map[image.Point]bool {
	floorNew := make(map[image.Point]bool)
	for tile := range floor {
		// Black tiles
		if floor[tile] {
			adjTrue := 0
			for _, n := range neighbours(tile) {
				if floor[n] {
					adjTrue++
				}
			}
			if adjTrue == 0 || adjTrue > 2 {
				floorNew[tile] = false // not strictly needed
			} else {
				floorNew[tile] = true
			}

			// White tiles: only need to consider those around black tiles
			for _, n := range neighbours(tile) {
				if !floor[n] {
					adjTrue = 0
					for _, nn := range neighbours(n) {
						if floor[nn] {
							adjTrue++
						}
					}
					if adjTrue == 2 {
						floorNew[n] = true
					}
				}
			}
		}
	}
	return floorNew
}

func neighbours(pt image.Point) []image.Point {
	var n []image.Point
	for _, i := range []image.Point{e, se, sw, w, nw, ne} {
		n = append(n, pt.Add(i))
	}
	return n
}

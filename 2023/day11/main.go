package main

import (
	"fmt"
	"image"
	"slices"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	galaxies := map[image.Point]struct{}{}
	s := strings.Split(util.ReadAll(), "\n")
	h, w := len(s), len(s[0])
	for y := range s {
		for x := range s[y] {
			if s[y][x] == '#' {
				galaxies[image.Pt(x, y)] = struct{}{}
			}
		}
	}
	var emptyRows, emptyCols []int
	for x := 0; x < w; x++ {
		empty := true
		for y := 0; y < h; y++ {
			if _, ok := galaxies[image.Pt(x, y)]; ok {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, x)
		}
	}
	for y := 0; y < h; y++ {
		empty := true
		for x := 0; x < w; x++ {
			if _, ok := galaxies[image.Pt(x, y)]; ok {
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, y)
		}
	}
	var expanded1, expanded2 []image.Point
	for g := range galaxies {
		deltaX, _ := slices.BinarySearch(emptyCols, g.X)
		deltaY, _ := slices.BinarySearch(emptyRows, g.Y)
		expanded1 = append(expanded1, g.Add(image.Pt(deltaX, deltaY)))
		expanded2 = append(expanded2, g.Add(image.Pt(deltaX, deltaY).Mul(1e6-1)))
	}

	sum1, sum2 := 0, 0
	for i := 0; i < len(expanded1)-1; i++ {
		for j := i + 1; j < len(expanded1); j++ {
			sum1 += util.Manhattan(expanded1[i], expanded1[j])
			sum2 += util.Manhattan(expanded2[i], expanded2[j])
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

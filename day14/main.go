package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"

	"github.com/adityapandey/adventofcode2017-go/util"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	grid := make(map[image.Point]struct{})
	for i := 0; i < 128; i++ {
		line := append(input, []byte(fmt.Sprintf("-%d", i))...)
		knot := util.NewKnot(256)
		hashString := knot.DenseHash(line)
		for j := 0; j < 32; j++ {
			d := dec(hashString[j])
			for k := 0; k < 4; k++ {
				if d%2 == 1 {
					grid[image.Pt(j*4+(3-k), i)] = struct{}{}
				}
				d /= 2
			}
		}
	}
	fmt.Println(len(grid))
	fmt.Println(regions(grid))
}

func dec(c byte) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c) - '0'
	case c >= 'a' && c <= 'f':
		return int(c) - 'a' + 10
	}
	return 0
}

func regions(grid map[image.Point]struct{}) int {
	var c int
	for p := range grid {
		c++
		visit(p, grid)
	}
	return c
}

func visit(p image.Point, grid map[image.Point]struct{}) {
	delete(grid, p)
	for _, n := range []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		if _, ok := grid[p.Add(n)]; ok {
			visit(p.Add(n), grid)
		}
	}
}

package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

type drone []int

func (d drone) hasBeam(p image.Point) bool {
	in := make(chan int, 2)
	in <- p.X
	in <- p.Y
	close(in)
	out := machine.Run(d, in)
	return <-out != 0
}

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}
	d := drone(program)

	var sum int
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if d.hasBeam(image.Pt(x, y)) {
				sum++
			}
		}
	}
	fmt.Println(sum)

	x, y := 0, 99
	for {
		for !d.hasBeam(image.Pt(x, y)) {
			x++
		}
		if d.hasBeam(image.Pt(x+99, y-99)) {
			break
		}
		y++
	}
	fmt.Println(10_000*x + (y - 99))
}

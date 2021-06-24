package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}

	grid := make(map[image.Point]int)
	paint(program, grid)
	fmt.Println(len(grid))

	grid = map[image.Point]int{image.Pt(0, 0): 1}
	paint(program, grid)
	var points []image.Point
	for p := range grid {
		points = append(points, p)
	}
	b := util.Bounds(points)
	for y := b.Min.Y; y <= b.Max.Y; y++ {
		for x := b.Min.X; x <= b.Max.X; x++ {
			if grid[image.Pt(x, y)] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func paint(program []int, grid map[image.Point]int) {
	in, out := make(chan int, 1), make(chan int)
	robot := machine.New(program, in, out)
	go robot.Run()
	var p image.Point
	dir := util.DirFromByte('^')
	for {
		in <- grid[p]
		grid[p] = <-out
		o, ok := <-out
		if o == 1 {
			dir = dir.Next()
		} else {
			dir = dir.Prev()
		}
		p = p.Add(dir.PointR())
		if !ok {
			break
		}
	}
	close(in)
}

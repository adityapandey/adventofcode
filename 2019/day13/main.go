package main

import (
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

type tile int

const (
	empty tile = iota
	wall
	block
	paddle
	ball
)

type move int

const (
	left move = iota - 1
	none
	right
)

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}

	fmt.Println(countBlocks(program))

	refreshRate := 5 * time.Millisecond
	score := play(program, refreshRate)
	fmt.Println(score)
}

func countBlocks(program []int) int {
	grid := make(map[image.Point]tile)
	in := make(chan int)
	close(in)
	out := machine.Run(program, in)
	for {
		x, ok := <-out
		if !ok {
			break
		}
		y := <-out
		grid[image.Pt(x, y)] = tile(<-out)
	}
	var n int
	for _, t := range grid {
		if t == block {
			n++
		}
	}
	return n
}

func play(program []int, refreshRate time.Duration) int {
	grid := make(map[image.Point]tile)
	program[0] = 2
	in, out := make(chan int), make(chan int)
	m := machine.New(program, in, out)
	go m.Run()
	refresh := time.Tick(refreshRate)
	var done bool
	var xBall, xPaddle, score int
	for !done {
		select {
		case x, ok := <-out:
			if !ok {
				done = true
				break
			}
			y := <-out
			p := image.Pt(x, y)
			tileId := <-out
			if x != -1 {
				grid[p] = tile(tileId)
			} else {
				score = tileId
			}
			switch tile(tileId) {
			case ball:
				xBall = p.X
			case paddle:
				xPaddle = p.X
			}
		case <-refresh:
			delta := xPaddle - xBall
			switch {
			case delta < 0:
				in <- int(right)
			case delta > 0:
				in <- int(left)
			default:
				in <- int(none)
			}
		}
	}
	close(in)
	return score
}

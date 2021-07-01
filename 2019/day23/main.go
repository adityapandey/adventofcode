package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

type nic struct {
	in, out chan int
}

type nat struct {
	x, y int
}

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}
	var nics [50]nic
	for i := 0; i < 50; i++ {
		in := make(chan int)
		out := machine.Run(program, in)
		in <- i
		in <- -1
		nics[i] = nic{in, out}
	}
	var natSignal bool
	var curr, prev nat
	idle := make(map[int]struct{})
loop:
	for {
		for i := 0; i < 50; i++ {
			select {
			case dest := <-nics[i].out:
				x := <-nics[i].out
				y := <-nics[i].out
				if dest == 255 {
					if !natSignal {
						fmt.Println(y)
						natSignal = true
					}
					curr = nat{x, y}
				} else {
					nics[dest].in <- x
					nics[dest].in <- y
				}
				delete(idle, i)
			case nics[i].in <- -1:
				idle[i] = struct{}{}
			}
			if len(idle) == 50 {
				if prev.y == curr.y {
					fmt.Println(curr.y)
					break loop
				}
				nics[0].in <- curr.x
				nics[0].in <- curr.y
				prev = curr
				idle = make(map[int]struct{})
			}
		}
	}
}

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}

	in := make(chan int, 1)
	in <- 1
	var out []int
	for o := range machine.Run(program, in) {
		out = append(out, o)
	}
	for _, z := range out[:len(out)-1] {
		if z != 0 {
			log.Fatal("Expected zero")
		}
	}
	fmt.Println(out[len(out)-1])

	in <- 5
	for o := range machine.Run(program, in) {
		fmt.Println(o)
	}
}

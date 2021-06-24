package main

import (
	"fmt"
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
	fmt.Println(<-machine.Run(program, in))
	in <- 2
	close(in)
	fmt.Println(<-machine.Run(program, in))
}

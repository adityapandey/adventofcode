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

	fmt.Println(damage(program, `NOT C J
AND D J
NOT A T
OR T J
WALK
`))

	fmt.Println(damage(program, `NOT H J
OR C J
AND B J
AND A J
NOT J J
AND D J
RUN
`))
}

func damage(program []int, script string) int {
	var sb strings.Builder
	in := make(chan int, len(script))
	for i := 0; i < len(script); i++ {
		in <- int(script[i])
	}
	close(in)
	out := machine.Run(program, in)
	var o int
	for o = range out {
		fmt.Fprintf(&sb, "%c", o)
	}
	if o == '\n' {
		fmt.Println(sb.String())
		panic("Did not finish")
	}
	return o
}

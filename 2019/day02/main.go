package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := util.ReadAll()
	var program []int
	for _, n := range strings.Split(input, ",") {
		program = append(program, util.Atoi(n))
	}

	p := make([]int, len(program))
	copy(p, program)
	p[1], p[2] = 12, 2
	execute(p)
	fmt.Println(p[0])

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(p, program)
			p[1], p[2] = noun, verb
			execute(p)
			if p[0] == 19690720 {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}

func execute(program []int) {
	var ip int
	for {
		switch program[ip] {
		case 1:
			program[program[ip+3]] = program[program[ip+1]] + program[program[ip+2]]
		case 2:
			program[program[ip+3]] = program[program[ip+1]] * program[program[ip+2]]
		case 99:
			return
		default:
			log.Fatal("unknown opcode: ", program[ip])
		}
		ip += 4
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/adityapandey/adventofcode/2018/machine"
)

type line struct {
	opcode  string
	a, b, c int
}

func main() {
	var program []line
	var ipRegister int
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Sscanf(s.Text(), "#ip %d", &ipRegister)
	for s.Scan() {
		var l line
		fmt.Sscanf(s.Text(), "%s %d %d %d", &l.opcode, &l.a, &l.b, &l.c)
		program = append(program, l)
	}

	// Part 1
	// The only part of the program that is influenced by state[0] is:
	//   29 eqrr 5 0 4
	//   30 addr 4 1 1
	//   31 seti 5 5 1
	//   [end of program]
	// So if at inst 29, state[5] == state[0], ip (1) will become 32, causing it to halt
	m := machine.New(6)
	var ip int
	var state0 int
	eqrr0 := program[28].a
	for {
		m.R[ipRegister] = ip
		if ip == 29 {
			state0 = m.R[eqrr0]
			fmt.Println(state0)
			break
		}
		line := program[ip]
		m.Execute(line.opcode, line.a, line.b, line.c)
		ip = m.R[ipRegister]
		ip++
		if ip >= len(program) {
			log.Fatal("program exit before executing inst 29")
		}
	}

	// Part 2
	// Assuming state[5] repeats, break at the last seen value when the first one reoccurs.
	for i := 0; i < 6; i++ {
		m.R[i] = 0
	}
	ip = 0
	var lastseen int
	seen := make(map[int]int)
	for {
		m.R[ipRegister] = ip
		line := program[ip]
		m.Execute(line.opcode, line.a, line.b, line.c)
		ip = m.R[ipRegister]
		ip++
		if ip == 29 {
			seen[m.R[eqrr0]]++
			if seen[m.R[eqrr0]] == 2 {
				fmt.Println(lastseen)
				break
			}
			lastseen = m.R[eqrr0]
		}
		if ip >= len(program) {
			log.Fatal("program exit before executing last inst 29")
		}
	}
}

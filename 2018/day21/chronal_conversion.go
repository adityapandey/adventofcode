package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type State [6]int

func (s State) String() string {
	return fmt.Sprintf("%15d,%15d,%15d,%15d,%15d,%15d", s[0], s[1], s[2], s[3], s[4], s[5])
}

type Instruction func(s State, i [2]int, o int) State

func addr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] + s[i[1]]
	return
}

func addi(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] + i[1]
	return
}

func mulr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] * s[i[1]]
	return
}

func muli(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] * i[1]
	return
}

func banr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] & s[i[1]]
	return
}

func bani(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] & i[1]
	return
}

func borr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] | s[i[1]]
	return
}

func bori(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] | i[1]
	return
}

func setr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]]
	return
}

func seti(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = i[0]
	return
}

func gtir(s State, i [2]int, o int) (r State) {
	r = s
	if i[0] > s[i[1]] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func gtri(s State, i [2]int, o int) (r State) {
	r = s
	if s[i[0]] > i[1] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func gtrr(s State, i [2]int, o int) (r State) {
	r = s
	if s[i[0]] > s[i[1]] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func eqir(s State, i [2]int, o int) (r State) {
	r = s
	if i[0] == s[i[1]] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func eqri(s State, i [2]int, o int) (r State) {
	r = s
	if s[i[0]] == i[1] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func eqrr(s State, i [2]int, o int) (r State) {
	r = s
	if s[i[0]] == s[i[1]] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

var inst = map[string]Instruction{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}

type Line struct {
	opcode string
	i      [2]int
	o      int
}

func main() {
	var program []Line
	var ipReg int
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Sscanf(s.Text(), "#ip %d", &ipReg)
	for s.Scan() {
		var l Line
		fmt.Sscanf(s.Text(), "%s %d %d %d", &l.opcode, &l.i[0], &l.i[1], &l.o)
		program = append(program, l)
	}

	// Part 1
	// The only part of the program that is influenced by state[0] is:
	//   29 eqrr 5 0 4
	//   30 addr 4 1 1
	//   31 seti 5 5 1
	//   [end of program]
	// So if at inst 29, state[5] == state[0], ip (1) will become 32, causing it to halt
	var state State
	var ip int
	var state0 int
	eqrr0 := program[28].i[0]
	for {
		state[ipReg] = ip
		if ip == 29 {
			state0 = state[eqrr0]
			fmt.Println(state0)
			break
		}
		line := program[ip]
		state = inst[line.opcode](state, line.i, line.o)
		ip = state[ipReg]
		ip++
		if ip >= len(program) {
			log.Fatal("program exit before executing inst 29")
		}
	}

	// Part 2
	// Assuming state[5] repeats, break at the last seen value when the first one reoccurs.
	state, ip = State{0, 0, 0, 0, 0, 0}, 0
	var lastseen int
	seen := make(map[int]int)
	for {
		state[ipReg] = ip
		line := program[ip]
		state = inst[line.opcode](state, line.i, line.o)
		ip = state[ipReg]
		ip++
		if ip == 29 {
			seen[state[eqrr0]]++
			if seen[state[eqrr0]] == 2 {
				fmt.Println(lastseen)
				break
			}
			lastseen = state[eqrr0]
		}
		if ip >= len(program) {
			log.Fatal("program exit before executing last inst 29")
		}
	}
}

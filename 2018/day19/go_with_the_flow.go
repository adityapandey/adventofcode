package main

import (
	"bufio"
	"fmt"
	"math"
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
	var state State
	var ip int
	for {
		state[ipReg] = ip
		line := program[ip]
		state = inst[line.opcode](state, line.i, line.o)
		ip = state[ipReg]
		ip++
		if ip >= len(program) {
			break
		}
	}
	fmt.Println(state[0])

	// Part 2
	// Intermediate states:
	// state = State{0, 10551360, 0, 1, 4, 10551364}
	// state = State{1, 5275678, 0, 2, 4, 10551364}
	// state = State{3, 10551360, 0, 2, 4, 10551364}
	// state = State{3, 3517118, 0, 3, 4, 10551364}
	// state = State{3, 10551360, 0, 3, 4, 10551364}
	// state = State{3, 2637838, 0, 4, 4, 10551364}
	// state = State{7, 10551360, 0, 4, 4, 10551364}
	// state = State{7, 10551360, 0, 5, 4, 10551364}
	// state = State{7, 10551360, 0, 36, 4, 10551364}
	// state = State{7, 285170, 0, 37, 4, 10551364}
	// state = State{44, 10551360, 0, 37, 4, 10551364}
	// state = State{44, 10551360, 0, 38, 4, 10551364}
	// state = State{44, 10551360, 0, 73, 4, 10551364}
	// state = State{44, 142585, 0, 74, 4, 10551364}

	state = State{1, 0, 0, 0, 0, 0}
	ip = 0
	for i := 0; i < 20; i++ {
		state[ipReg] = ip
		line := program[ip]
		state = inst[line.opcode](state, line.i, line.o)
		ip = state[ipReg]
		ip++
	}

	// After several state changes,
	//   - register 5 contains 10551364
	//   - the following program is executed:
	// 11 seti 2 8 4
	// 3  mulr 3 1 2
	// 4  eqrr 2 5 2
	// 5  addr 2 4 4
	// 6  addi 4 1 4
	// 8  addi 1 1 1
	// 9  gtrr 1 5 2
	// 10 addr 4 2 4
	//
	// 7  addr 3 0 0

	// f = 10551364
	// c = d*b
	// if (c == f) { a += d}
	// b++
	// if (b > f) { d++; b = 0; }

	// This is a very slow program that sums all factors of 'f' in a.

	f, sum := state[5], 0
	for i := 1; i < int(math.Floor(math.Sqrt(float64(f)))); i++ {
		if f%i == 0 {
			sum += i
			sum += f / i
		}
	}
	fmt.Println(sum)
}

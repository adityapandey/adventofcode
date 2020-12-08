// https://adventofcode.com/2020/day/8
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type opcode int

const (
	acc opcode = iota
	jmp
	nop
)

func fromCode(s string) opcode {
	switch s {
	case "acc":
		return acc
	case "jmp":
		return jmp
	case "nop":
		return nop
	}
	log.Fatal("Unknown opcode: ", s)
	return -1
}

type instruction struct {
	op  opcode
	arg int
}

type runner struct {
	accumulator int
	program     map[int]instruction
	pc          int
}

func newRunner(program map[int]instruction) *runner {
	return &runner{0, program, 0}
}

func (r *runner) step() {
	//fmt.Printf("%+v\n", r)
	curr := r.program[r.pc]
	switch curr.op {
	case acc:
		r.accumulator += curr.arg
		r.pc++
	case jmp:
		r.pc += curr.arg
	case nop:
		r.pc++
	}
}

func main() {
	program := make(map[int]instruction)
	s := bufio.NewScanner(os.Stdin)
	c := 0
	for s.Scan() {
		var opcode string
		var arg int
		fmt.Sscanf(s.Text(), "%s %d", &opcode, &arg)
		program[c] = instruction{fromCode(opcode), arg}
		c++
	}

	// Part 1
	r := newRunner(program)
	visited := make(map[int]struct{})
	for _, ok := visited[r.pc]; !ok; _, ok = visited[r.pc] {
		visited[r.pc] = struct{}{}
		r.step()
	}
	fmt.Println(r.accumulator)
}

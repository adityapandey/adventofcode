package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type cpu struct {
	regs map[string]int
	ip   int
	prog []string
}

func newCPU(prog []string) *cpu {
	return &cpu{
		regs: map[string]int{},
		ip:   0,
		prog: prog,
	}
}

func (c *cpu) step() bool {
	if c.ip < 0 || c.ip >= len(c.prog) {
		return false
	}
	isJmp := false
	f := strings.Fields(strings.ReplaceAll(c.prog[c.ip], ",", ""))
	switch f[0] {
	case "hlf":
		c.regs[f[1]] /= 2
	case "tpl":
		c.regs[f[1]] *= 3
	case "inc":
		c.regs[f[1]]++
	case "jmp":
		offset := util.Atoi(f[1])
		isJmp = true
		c.ip += offset
	case "jie":
		offset := util.Atoi(f[2])
		if c.regs[f[1]]%2 == 0 {
			isJmp = true
			c.ip += offset
		}
	case "jio":
		offset := util.Atoi(f[2])
		if c.regs[f[1]] == 1 {
			isJmp = true
			c.ip += offset
		}
	}
	if !isJmp {
		c.ip++
	}
	return true
}

func main() {
	input := strings.Split(util.ReadAll(), "\n")
	c := newCPU(input)
	for c.step() {
	}
	fmt.Println(c.regs["b"])
	c = newCPU(input)
	c.regs["a"] = 1
	for c.step() {
	}
	fmt.Println(c.regs["b"])
}

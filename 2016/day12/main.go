package main

import (
	"fmt"
	"strconv"
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
	line := c.prog[c.ip]
	f := strings.Fields(line)
	isJump := false
	switch f[0] {
	case "cpy":
		c.regs[f[2]] = c.getValue(f[1])
	case "inc":
		c.regs[f[1]]++
	case "dec":
		c.regs[f[1]]--
	case "jnz":
		if c.getValue(f[1]) != 0 {
			isJump = true
			c.ip += c.getValue(f[2])
		}
	}
	if !isJump {
		c.ip++
	}
	return true
}

func (c *cpu) getValue(s string) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return c.regs[s]
}

func main() {
	prog := strings.Split(util.ReadAll(), "\n")
	c := newCPU(prog)
	for c.step() {
	}
	fmt.Println(c.regs["a"])
	c = newCPU(prog)
	c.regs["c"] = 1
	for c.step() {
	}
	fmt.Println(c.regs["a"])
}

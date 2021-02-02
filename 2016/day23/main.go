package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type instruction struct {
	op   string
	args []string
}

type cpu struct {
	regs  map[string]int
	ip    int
	prog  []instruction
	debug bool
}

func newCPU(prog []string) *cpu {
	c := &cpu{
		regs: map[string]int{},
		ip:   0,
	}
	for _, line := range prog {
		f := strings.Fields(line)
		var instr instruction
		instr.op = f[0]
		for j := 1; j < len(f); j++ {
			instr.args = append(instr.args, f[j])
		}
		c.prog = append(c.prog, instr)
	}
	return c
}

func (c *cpu) step() bool {
	if c.ip < 0 || c.ip >= len(c.prog) {
		return false
	}
	instr := c.prog[c.ip]
	if c.debug {
		fmt.Printf("%-3d %-30v %v %v\n", c.ip, fmt.Sprint(c.regs), instr.op, instr.args)
	}
	isJump := false
	switch instr.op {
	case "cpy":
		if _, err := strconv.Atoi(instr.args[1]); err != nil {
			c.regs[instr.args[1]] = c.getValue(instr.args[0])
		}
	case "inc":
		c.regs[instr.args[0]]++
	case "dec":
		c.regs[instr.args[0]]--
	case "jnz":
		if c.getValue(instr.args[0]) != 0 {
			isJump = true
			c.ip += c.getValue(instr.args[1])
		}
	case "tgl":
		tgl := c.ip + c.getValue(instr.args[0])
		if tgl < 0 || tgl >= len(c.prog) {
			break
		}
		switch c.prog[tgl].op {
		case "inc":
			c.prog[tgl].op = "dec"
		case "jnz":
			c.prog[tgl].op = "cpy"
		default:
			switch len(c.prog[tgl].args) {
			case 1:
				c.prog[tgl].op = "inc"
			case 2:
				c.prog[tgl].op = "jnz"
			}
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
	c := newCPU(strings.Split(util.ReadAll(), "\n"))
	c.regs["a"] = 7
	for c.step() {
	}
	fmt.Println(c.regs["a"])
	// On inspection:
	// output => a! + c*d (when a! is reached)
	cTimesD := c.regs["a"] - factorial(7)
	fmt.Println(factorial(12) + cTimesD)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	f := 1
	for i := n; i > 0; i-- {
		f *= i
	}
	return f
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type cpu struct {
	regs  map[string]int
	ip    int
	prog  []string
	out   []int
	debug bool
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
	if c.debug {
		fmt.Printf("%-4d %-30s %s\n", c.ip, fmt.Sprint(c.regs), line)
	}
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
	case "out":
		c.out = append(c.out, c.getValue(f[1]))
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

// Program:
// a = a + b*c
// output: a%2, a/2 % 2, (a/2)/2 % 2 etc, then repeat when a goes to 0.
func main() {
	cpu := newCPU(strings.Split(util.ReadAll(), "\n"))
	for i := 0; i < 3; i++ { // 3 steps to populate b and c.
		cpu.step()
	}
	b, c := cpu.regs["b"], cpu.regs["c"]
	a := (b * c) % 2 // sum must be even for output to start with 0.
	for !alternates(a + b*c) {
		a += 2
	}
	fmt.Println(a)
}

func alternates(n int) bool {
	prev := n % 2
	for n > 0 {
		n /= 2
		if n%2 == prev {
			return false
		}
		prev = n % 2
	}
	return true
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/adityapandey/adventofcode/2018/machine"
)

type line struct {
	opcode  string
	a, b, c int
}

type device struct {
	m          *machine.Machine
	ip         int
	ipRegister int
	program    []line
}

func (d *device) step() bool {
	d.m.R[d.ipRegister] = d.ip
	line := d.program[d.ip]
	d.m.Execute(line.opcode, line.a, line.b, line.c)
	d.ip = d.m.R[d.ipRegister]
	d.ip++
	if d.ip >= len(d.program) {
		return false
	}
	return true
}

func main() {
	var ipRegister int
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Sscanf(s.Text(), "#ip %d", &ipRegister)
	d := device{
		m:          machine.New(6),
		ipRegister: ipRegister}
	for s.Scan() {
		var opcode string
		var a, b, c int
		fmt.Sscanf(s.Text(), "%s %d %d %d", &opcode, &a, &b, &c)
		d.program = append(d.program, line{opcode, a, b, c})
	}

	// Part 1
	for d.step() {
	}
	fmt.Println(d.m.R[0])

	// Part 2
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

	d.m.R[0] = 1
	for i := 1; i < 6; i++ {
		d.m.R[i] = 0
	}
	d.ip = 0
	for i := 0; i < 20; i++ {
		d.step()
	}
	f, sum := d.m.R[5], 0
	for i := 1; i < int(math.Floor(math.Sqrt(float64(f)))); i++ {
		if f%i == 0 {
			sum += i
			sum += f / i
		}
	}
	fmt.Println(sum)
}

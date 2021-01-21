package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type runner struct {
	prog []string
	ip   int
	id   int
	regs map[string]int
	freq int
	sndq *[]int
	rcvq *[]int
	sent int
}

func newRunner(prog []string, id int, sndq *[]int, rcvq *[]int) *runner {
	return &runner{
		prog: prog,
		id:   id,
		regs: map[string]int{"p": id},
		sndq: sndq,
		rcvq: rcvq,
	}
}

func (r *runner) step(part2 bool) bool {
	line := r.prog[r.ip]
	f := strings.Fields(line)
	var isJmp bool
	switch f[0] {
	case "snd":
		if !part2 {
			r.freq = r.regs[f[1]]
		} else {
			*r.sndq = append(*r.sndq, r.regs[f[1]])
			r.sent++
		}
	case "set":
		r.regs[f[1]] = r.getVal(f[2])
	case "add":
		r.regs[f[1]] += r.getVal(f[2])
	case "mul":
		r.regs[f[1]] *= r.getVal(f[2])
	case "mod":
		r.regs[f[1]] %= r.getVal(f[2])
	case "rcv":
		if !part2 {
			return false
		}
		if len(*r.rcvq) == 0 {
			return false
		}
		r.regs[f[1]] = (*r.rcvq)[0]
		*r.rcvq = (*r.rcvq)[1:]
	case "jgz":
		if r.getVal(f[1]) > 0 {
			isJmp = true
			r.ip += r.getVal(f[2])
		}
	}
	if !isJmp {
		r.ip++
	}
	return true
}

func (r *runner) getVal(f string) int {
	if v, err := strconv.Atoi(f); err == nil {
		return v
	}
	return r.regs[f]
}

func main() {
	var prog []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		prog = append(prog, s.Text())
	}
	var q0, q1 []int
	r := newRunner(prog, 0, &q0, &q1)
	for {
		if !r.step(false) {
			break
		}
	}
	fmt.Println(r.freq)

	q0, q1 = []int{}, []int{}
	r0 := newRunner(prog, 0, &q0, &q1)
	r1 := newRunner(prog, 1, &q1, &q0)
	for {
		if !r0.step(true) && !r1.step(true) {
			fmt.Println(r1.sent)
			break
		}
	}
}

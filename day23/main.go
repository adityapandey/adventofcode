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
	regs map[string]int
	mul  int
}

func newRunner(prog []string) *runner {
	return &runner{
		prog: prog,
		regs: map[string]int{},
	}
}

func (r *runner) step() bool {
	if r.ip < 0 || r.ip >= len(r.prog) {
		return false
	}
	line := r.prog[r.ip]
	f := strings.Fields(line)
	var isJmp bool
	switch f[0] {
	case "set":
		r.regs[f[1]] = r.getVal(f[2])
	case "sub":
		r.regs[f[1]] -= r.getVal(f[2])
	case "mul":
		r.regs[f[1]] *= r.getVal(f[2])
		r.mul++
	case "jnz":
		if r.getVal(f[1]) != 0 {
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

	r := newRunner(prog)
	for r.step() {
	}
	fmt.Println(r.mul)

	// 	b = 81
	// 	c = b
	// 	if a != 0 {
	// 		b *= 100
	// 		b += 100000
	// 		c = b
	// 		c += 17000
	// 	}
	// loop1:
	// 	f = 1
	// 	d = 2
	// loop2:
	// 	e = 2
	// loop3:
	// 	g = d
	// 	g *= e
	// 	g -= b
	// 	if g == 0 {
	// 		f = 0
	// 	}
	// 	e++
	// 	g = e
	// 	g -= b
	// 	if g == 0 {
	// 		d++ // 21
	// 		g = d
	// 		g -= b
	// 		if g == 0 {
	// 			if f == 0 {
	// 				h++
	// 			}
	// 			g = b
	// 			g -= c
	// 			if g == 0 {
	// 				end
	// 			} else {
	// 				b += 17
	// 				goto loop1
	// 			}
	// 		} else {
	// 			goto loop2
	// 		}
	// 	} else {
	// 		goto loop3
	// 	}

	// b = 81
	// c = 81
	// if a != 0 {
	// 	b = 108100
	// 	c = 125100
	// }
	// for {
	// 	f = 1
	// 	d = 2
	// 	for {
	// 		e = 2
	// 		for {
	// 			if d*e == b {
	// 				f = 0
	// 			}
	// 			e++
	// 			if e == b {
	// 				break
	// 			}
	// 		}
	// 		d++
	// 		if d == b {
	// 			break
	// 		}
	// 	}
	// 	if f == 0 {
	// 		h++
	// 	}
	// 	if b == c {
	// 		break
	// 	}
	// 	b += 17
	// }

	var b, c, d, f, h int

	r = newRunner(prog)
	r.regs["a"] = 1
	for i := 0; i < 10; i++ {
		r.step()
	}
	b = r.regs["b"]
	c = r.regs["c"]
	for {
		f = 1
		d = 2
		for {
			if b%d == 0 {
				f = 0
			}
			d++
			if d == b {
				break
			}
		}
		if f == 0 {
			h++
		}
		if b == c {
			break
		}
		b += 17
	}
	fmt.Println(h)
}

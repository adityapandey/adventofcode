package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	wires := make(map[string]uint16)
	instr := make(map[string]string)
	s := util.ScanAll()
	for s.Scan() {
		f := strings.Fields(s.Text())
		instr[f[len(f)-1]] = strings.Join(f[:len(f)-2], " ")
	}

	a := eval("a", instr, wires)
	fmt.Println(a)

	wires = make(map[string]uint16)
	wires["b"] = a
	a = eval("a", instr, wires)
	fmt.Println(a)
}

func eval(w string, instr map[string]string, wires map[string]uint16) uint16 {
	if v, ok := wires[w]; ok {
		return v
	}
	if v, err := strconv.Atoi(w); err == nil {
		return uint16(v)
	}
	var v uint16
	f := strings.Fields(instr[w])
	switch len(f) {
	case 1:
		v = eval(f[0], instr, wires)
	case 2:
		v = ^eval(f[1], instr, wires)
	case 3:
		l, r := eval(f[0], instr, wires), eval(f[2], instr, wires)
		switch f[1] {
		case "AND":
			v = l & r
		case "OR":
			v = l | r
		case "LSHIFT":
			v = l << r
		case "RSHIFT":
			v = l >> r
		}
	}
	wires[w] = v
	return v
}

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var monkeys = map[string]*monkey{}

type monkey struct {
	val    int
	m1, m2 string
	op     string
}

func (m *monkey) yell() int {
	if m.m1 == "" && m.m2 == "" {
		return m.val
	}
	m1y, m2y := monkeys[m.m1].yell(), monkeys[m.m2].yell()
	switch m.op {
	case "+":
		return m1y + m2y
	case "-":
		return m1y - m2y
	case "*":
		return m1y * m2y
	case "/":
		return m1y / m2y
	}
	return -1
}

func main() {
	s := util.ScanAll()
	for s.Scan() {
		sp := strings.Split(s.Text(), ": ")
		name := sp[0]
		f := strings.Fields(sp[1])
		if len(f) == 1 {
			monkeys[name] = &monkey{val: util.Atoi(f[0])}
		} else {
			monkeys[name] = &monkey{m1: f[0], op: f[1], m2: f[2]}
		}
	}

	fmt.Println(monkeys["root"].yell())

	monkeys["humn"] = &monkey{val: 0}
	m1, m2 := monkeys["root"].m1, monkeys["root"].m2
	if monkeys[m1].yell() < monkeys[m2].yell() {
		m1, m2 = m2, m1
	}

	var solns []int
	for soln, ok := math.MaxInt, true; ok; soln, ok = sort.Find(soln,
		func(i int) int {
			monkeys["humn"].val = i
			return monkeys[m1].yell() - monkeys[m2].yell()
		}) {
		solns = append(solns, soln)
	}
	sort.Ints(solns)
	fmt.Println(solns[0])
}

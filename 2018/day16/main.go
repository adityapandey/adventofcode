package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/2018/machine"
	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := strings.Split(util.ReadAll(), "\n\n\n\n")
	samples := strings.Split(input[0], "\n\n")
	var sum3orMoreFits int
	codeCanBeOp := make(map[int]map[string]struct{})
	codeCannotBeOp := make(map[int]map[string]struct{})
	for _, s := range samples {
		sp := strings.Split(s, "\n")
		before, after := machine.New(4), machine.New(4)
		fmt.Sscanf(sp[0], "Before: [%d, %d, %d, %d]", &before.R[0], &before.R[1], &before.R[2], &before.R[3])
		fmt.Sscanf(sp[2], "After:  [%d, %d, %d, %d]", &after.R[0], &after.R[1], &after.R[2], &after.R[3])
		var code, a, b, c int
		fmt.Sscanf(sp[1], "%d %d %d %d", &code, &a, &b, &c)

		var fits int
		for o, i := range machine.Instructions {
			m := copyOf(before)
			i(m, a, b, c)
			if m.R[0] == after.R[0] && m.R[1] == after.R[1] && m.R[2] == after.R[2] && m.R[3] == after.R[3] {
				fits++
				if _, ok := codeCanBeOp[code]; !ok {
					codeCanBeOp[code] = make(map[string]struct{})
				}
				codeCanBeOp[code][o] = struct{}{}
			} else {
				if _, ok := codeCannotBeOp[code]; !ok {
					codeCannotBeOp[code] = make(map[string]struct{})
				}
				codeCannotBeOp[code][o] = struct{}{}
			}
		}
		if fits >= 3 {
			sum3orMoreFits++
		}
	}
	fmt.Println(sum3orMoreFits)

	codeToOp := make(map[int]string)
	for code := range codeCanBeOp {
		for o := range codeCannotBeOp[code] {
			delete(codeCanBeOp[code], o)
		}
	}
	for len(codeToOp) < 16 {
		for code := range codeCanBeOp {
			if len(codeCanBeOp[code]) == 1 {
				codeToOp[code] = getOnlyEntry(codeCanBeOp[code])
			}
			for c, o := range codeToOp {
				if c != code {
					delete(codeCanBeOp[code], o)
				}
			}
		}
	}

	prog := input[1]
	m := machine.New(4)
	for _, line := range strings.Split(prog, "\n") {
		var code, a, b, c int
		fmt.Sscanf(line, "%d %d %d %d", &code, &a, &b, &c)
		m.Execute(codeToOp[code], a, b, c)
	}
	fmt.Println(m.R[0])
}

func copyOf(m *machine.Machine) *machine.Machine {
	n := len(m.R)
	c := machine.New(n)
	for i := 0; i < n; i++ {
		c.R[i] = m.R[i]
	}
	return c
}

func getOnlyEntry(m map[string]struct{}) string {
	for op := range m {
		return op
	}
	panic(len(m))
}

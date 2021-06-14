package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type device struct {
	r [4]int
}

type operation int

const (
	addr operation = iota
	addi
	mulr
	muli
	banr
	bani
	borr
	bori
	setr
	seti
	gtir
	gtri
	gtrr
	eqir
	eqri
	eqrr
)

type instruction func(d *device, a, b, c int)

var instructions = map[operation]instruction{
	addr: func(d *device, a, b, c int) {
		d.r[c] = d.r[a] + d.r[b]
	},
	addi: func(d *device, a, b, c int) {
		d.r[c] = d.r[a] + b
	},
	mulr: func(d *device, a, b, c int) {
		d.r[c] = d.r[a] * d.r[b]
	},
	muli: func(d *device, a, b, c int) {
		d.r[c] = d.r[a] * b
	},
	banr: func(d *device, a, b, c int) {
		d.r[c] = d.r[a] & d.r[b]
	},
	bani: func(d *device, a, b, c int) {
		d.r[c] = d.r[a] & b
	},
	borr: func(d *device, a, b, c int) {
		d.r[c] = d.r[a] | d.r[b]
	},
	bori: func(d *device, a, b, c int) {
		d.r[c] = d.r[a] | b
	},
	setr: func(d *device, a, b, c int) {
		d.r[c] = d.r[a]
	},
	seti: func(d *device, a, b, c int) {
		d.r[c] = a
	},
	gtir: func(d *device, a, b, c int) {
		if a > d.r[b] {
			d.r[c] = 1
		} else {
			d.r[c] = 0
		}
	},
	gtri: func(d *device, a, b, c int) {
		if d.r[a] > b {
			d.r[c] = 1
		} else {
			d.r[c] = 0
		}
	},
	gtrr: func(d *device, a, b, c int) {
		if d.r[a] > d.r[b] {
			d.r[c] = 1
		} else {
			d.r[c] = 0
		}
	},
	eqir: func(d *device, a, b, c int) {
		if a == d.r[b] {
			d.r[c] = 1
		} else {
			d.r[c] = 0
		}
	},
	eqri: func(d *device, a, b, c int) {
		if d.r[a] == b {
			d.r[c] = 1
		} else {
			d.r[c] = 0
		}
	},
	eqrr: func(d *device, a, b, c int) {
		if d.r[a] == d.r[b] {
			d.r[c] = 1
		} else {
			d.r[c] = 0
		}
	},
}

func main() {
	input := strings.Split(util.ReadAll(), "\n\n\n\n")
	samples := strings.Split(input[0], "\n\n")
	var sum3orMoreFits int
	codeCanBeOp := make(map[int]map[operation]struct{})
	codeCannotBeOp := make(map[int]map[operation]struct{})
	for _, s := range samples {
		sp := strings.Split(s, "\n")
		var before, after device
		fmt.Sscanf(sp[0], "Before: [%d, %d, %d, %d]", &before.r[0], &before.r[1], &before.r[2], &before.r[3])
		fmt.Sscanf(sp[2], "After:  [%d, %d, %d, %d]", &after.r[0], &after.r[1], &after.r[2], &after.r[3])
		var code, a, b, c int
		fmt.Sscanf(sp[1], "%d %d %d %d", &code, &a, &b, &c)

		var fits int
		for o, i := range instructions {
			d := before
			i(&d, a, b, c)
			if d == after {
				fits++
				if _, ok := codeCanBeOp[code]; !ok {
					codeCanBeOp[code] = make(map[operation]struct{})
				}
				codeCanBeOp[code][o] = struct{}{}
			} else {
				if _, ok := codeCannotBeOp[code]; !ok {
					codeCannotBeOp[code] = make(map[operation]struct{})
				}
				codeCannotBeOp[code][o] = struct{}{}
			}
		}
		if fits >= 3 {
			sum3orMoreFits++
		}
	}
	fmt.Println(sum3orMoreFits)

	codeToOp := make(map[int]operation)
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
	var d device
	for _, line := range strings.Split(prog, "\n") {
		var code, a, b, c int
		fmt.Sscanf(line, "%d %d %d %d", &code, &a, &b, &c)
		i := instructions[codeToOp[code]]
		i(&d, a, b, c)
	}
	fmt.Println(d.r[0])
}

func getOnlyEntry(m map[operation]struct{}) operation {
	for op := range m {
		return op
	}
	panic(len(m))
}

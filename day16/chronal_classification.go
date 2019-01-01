package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

type State [4]int

type Instruction func(s State, i [2]int, o int) State

func addr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] + s[i[1]]
	return
}

func addi(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] + i[1]
	return
}

func mulr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] * s[i[1]]
	return
}

func muli(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] * i[1]
	return
}

func banr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] & s[i[1]]
	return
}

func bani(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] & i[1]
	return
}

func borr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] | s[i[1]]
	return
}

func bori(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]] | i[1]
	return
}

func setr(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = s[i[0]]
	return
}

func seti(s State, i [2]int, o int) (r State) {
	r = s
	r[o] = i[0]
	return
}

func gtir(s State, i [2]int, o int) (r State) {
	r = s
	if i[0] > s[i[1]] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func gtri(s State, i [2]int, o int) (r State) {
	r = s
	if s[i[0]] > i[1] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func gtrr(s State, i [2]int, o int) (r State) {
	r = s
	if s[i[0]] > s[i[1]] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func eqir(s State, i [2]int, o int) (r State) {
	r = s
	if i[0] == s[i[1]] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func eqri(s State, i [2]int, o int) (r State) {
	r = s
	if s[i[0]] == i[1] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func eqrr(s State, i [2]int, o int) (r State) {
	r = s
	if s[i[0]] == s[i[1]] {
		r[o] = 1
	} else {
		r[o] = 0
	}
	return
}

func main() {
	instructions := []Instruction{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	candidates := make(map[int]map[int]struct{})
	var sum3Matches int
	s := bufio.NewScanner(os.Stdin)

	// Part 1
	for {
		var before, after State
		var opcode int
		var i [2]int
		var o int
		s.Scan()
		if _, err := fmt.Sscanf(s.Text(), "Before: [%d, %d, %d, %d]", &before[0], &before[1], &before[2], &before[3]); err != nil {
			break
		}
		s.Scan()
		fmt.Sscanf(s.Text(), "%d %d %d %d", &opcode, &i[0], &i[1], &o)
		s.Scan()
		fmt.Sscanf(s.Text(), "After: [%d, %d, %d, %d]", &after[0], &after[1], &after[2], &after[3])
		s.Scan()

		var matches int
		for x, inst := range instructions {
			if inst(before, i, o) == after {
				matches++
				if len(candidates[opcode]) == 0 {
					candidates[opcode] = make(map[int]struct{})
				}
				candidates[opcode][x] = struct{}{}
			}
		}

		if matches >= 3 {
			sum3Matches++
		}
	}
	fmt.Println(sum3Matches)

	// Generate map of opcode -> instruction
	instructionFor, err := match(candidates)
	if err != nil {
		log.Fatal(err)
	}

	// Part 2
	var state State
	s.Scan()
	for s.Scan() {
		var opcode int
		var i [2]int
		var o int
		fmt.Sscanf(s.Text(), "%d %d %d %d", &opcode, &i[0], &i[1], &o)
		state = instructions[instructionFor[opcode]](state, i, o)
	}
	fmt.Println(state[0])
}

func match(m map[int]map[int]struct{}) (map[int]int, error) {
	if len(m) == 0 {
		return make(map[int]int), nil
	}
	var order []int
	for k := range m {
		order = append(order, k)
	}
	sort.Slice(order, func(i, j int) bool { return len(m[order[i]]) < len(m[order[j]]) })
	for _, k := range order {
		if len(m[k]) == 0 {
			return nil, errors.New("no more values")
		}
		var v []int
		for val := range m[k] {
			v = append(v, val)
		}
		if len(m[k]) == 1 {
			mReduced := m
			delete(mReduced, k)
			for kReduced := range mReduced {
				delete(mReduced[kReduced], v[0])
			}
			if o, err := match(mReduced); err == nil {
				o[k] = v[0]
				return o, nil
			} else {
				return nil, err
			}
		}
		for _, val := range v {
			mReduced := m
			delete(mReduced, k)
			for kReduced := range mReduced {
				delete(mReduced[kReduced], val)
			}
			if o, err := match(mReduced); err == nil {
				o[k] = val
				return o, nil
			} else {
				continue
			}
		}
	}
	return nil, errors.New("no more keys")
}

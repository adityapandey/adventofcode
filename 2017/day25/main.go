package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type instr struct {
	write bool
	move  int
	next  byte
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	paras := strings.Split(string(input), "\n\n")
	preamble := strings.Split(paras[0], "\n")
	var init byte
	fmt.Sscanf(preamble[0], "Begin in state %c.", &init)
	var steps int
	fmt.Sscanf(preamble[1], "Perform a diagnostic checksum after %d steps.", &steps)
	m := make(map[byte]map[bool]instr)
	for _, para := range paras[1:] {
		parseState(strings.Split(para, "\n"), m)
	}

	tape := make(map[int]bool)
	state := init
	curr := 0
	for i := 0; i < steps; i++ {
		in := m[state][tape[curr]]
		state = in.next
		tape[curr] = in.write
		curr += in.move
	}
	var sum int
	for _, v := range tape {
		if v {
			sum++
		}
	}
	fmt.Println(sum)
}

func parseState(lines []string, m map[byte]map[bool]instr) {
	var state byte
	fmt.Sscanf(lines[0], "In state %c:", &state)
	m[state] = make(map[bool]instr)
	for val := 0; val <= 1; val++ {
		var i instr
		var w int
		fmt.Sscanf(lines[2+4*val], "    - Write the value %d.", &w)
		i.write = w == 1
		var dir string
		fmt.Sscanf(lines[3+4*val], "    - Move one slot to the %s", &dir)
		if dir == "right." {
			i.move = 1
		} else {
			i.move = -1
		}
		fmt.Sscanf(lines[4+4*val], "    - Continue with state %c.", &i.next)
		m[state][val == 1] = i
	}
}

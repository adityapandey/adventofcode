package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	s := strings.Split(util.ReadAll(), "\n\n")
	input := strings.Split(s[0], "\n")
	stacks := make([][]byte, (len(input[0])+1)/4)
	for _, line := range input {
		for i, b := range []byte(line) {
			if b >= 'A' && b <= 'Z' {
				stacks[(i-1)/4] = append(stacks[(i-1)/4], b)
			}
		}
	}

	steps := strings.Split(s[1], "\n")
	fmt.Println(move(stacks, steps, false))
	fmt.Println(move(stacks, steps, true))
}

func move(st [][]byte, steps []string, together bool) string {
	stacks := make([][]byte, len(st))
	// Reverse stacks, make them bottom to top.
	for i := range st {
		stacks[i] = make([]byte, len(st[i]))
		for j := range st[i] {
			stacks[i][j] = st[i][len(st[i])-j-1]
		}
	}

	for _, step := range steps {
		var n, from, to int
		fmt.Sscanf(step, "move %d from %d to %d", &n, &from, &to)
		from--
		to--
		if together {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-n:]...)
			stacks[from] = stacks[from][:len(stacks[from])-n]
		} else {
			for i := 0; i < n; i++ {
				stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
				stacks[from] = stacks[from][:len(stacks[from])-1]
			}
		}
	}

	b := make([]byte, len(stacks))
	for i := range stacks {
		b[i] = stacks[i][len(stacks[i])-1]
	}
	return string(b)
}

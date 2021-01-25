package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	prev string
	next byte
}

func main() {
	notes := make(map[string]byte)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	var initialState string
	fmt.Sscanf(s.Text(), "initial state: %s", &initialState)
	s.Scan()
	for s.Scan() {
		var prev string
		var next byte
		fmt.Sscanf(s.Text(), "%s => %c", &prev, &next)
		notes[prev] = next
	}

	// Part 1
	origin, start := 0, initialState
	for i := 0; i < 20; i++ {
		start, origin = generate(start, origin, &notes)
	}
	var sum int
	for i := 0; i < len(start); i++ {
		if start[i] == '#' {
			sum += i - origin
		}
	}
	fmt.Println(sum)

	// Part 2
	origin, start = 0, initialState
	var prevStart string
	var generation, prevOrigin int
	for prevStart != start {
		generation++
		prevStart, prevOrigin = start, origin
		start, origin = generate(start, origin, &notes)
	}

	finalGeneration := 50000000000
	finalOrigin := (finalGeneration-generation)*(origin-prevOrigin) + origin

	sum = 0
	for i := 0; i < len(start); i++ {
		if start[i] == '#' {
			sum += i - finalOrigin
		}
	}
	fmt.Println(sum)
}

func generate(start string, origin int, notes *map[string]byte) (string, int) {
	delta, start := padString(start)
	origin += delta
	next := make([]byte, len(start)-4)
	for i := 0; i < len(start)-4; i++ {
		next[i] = (*notes)[start[i:i+5]]
	}
	start = string(next)
	origin -= 2
	return start, origin
}

func padString(s string) (int, string) {
	i := strings.LastIndexByte(s, '#')
	if i != -1 {
		s = s[:i+1] + "...."
	}
	i = strings.IndexByte(s, '#')
	if i != -1 {
		s = "...." + s[i:]
	}
	return 4 - i, s
}

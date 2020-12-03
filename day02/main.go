package main

import (
	"bufio"
	"fmt"
	"os"
)

type entry struct {
	min, max int
	c        byte
	password string
}

func main() {
	var entries []entry
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var e entry
		fmt.Sscanf(s.Text(), "%d-%d %c: %s", &e.min, &e.max, &e.c, &e.password)
		entries = append(entries, e)
	}

	// Part 1
	var valid int
	for _, e := range entries {
		var c int
		for i := 0; i < len(e.password); i++ {
			if e.password[i] == e.c {
				c++
			}
		}
		if c >= e.min && c <= e.max {
			valid++
		}
	}
	fmt.Println(valid)

	// Part 2
	valid = 0
	for _, e := range entries {
		posMin := e.password[e.min-1] == e.c
		posMax := e.password[e.max-1] == e.c
		if posMin != posMax {
			valid++
		}
	}
	fmt.Println(valid)
}

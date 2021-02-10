package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	b := []byte(util.ReadAll())
	for {
		b = next(b)
		if valid(b) {
			break
		}
	}
	fmt.Println(string(b))
	for {
		b = next(b)
		if valid(b) {
			break
		}
	}
	fmt.Println(string(b))
}

func next(b []byte) []byte {
	for i, add := len(b)-1, true; add; i-- {
		b[i] = b[i] + 1
		if b[i] > 'z' {
			b[i] = 'a'
		} else {
			add = false
		}
	}
	return b
}

func valid(s []byte) bool {
	pairs := make(map[byte]struct{})
	var straight bool
	for i := range s {
		switch s[i] {
		case 'i', 'o', 'l':
			return false
		}
		if !straight && i < len(s)-2 {
			if s[i+1] == s[i]+1 && s[i+2] == s[i]+2 {
				straight = true
			}
		}
		if i < len(s)-1 && s[i] == s[i+1] {
			pairs[s[i]] = struct{}{}
		}
	}
	return len(pairs) >= 2 && straight
}

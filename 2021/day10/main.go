package main

import (
	"fmt"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

var points = map[byte][2]int{
	')': {3, 1},
	']': {57, 2},
	'}': {1197, 3},
	'>': {25137, 4},
}

var closing = map[byte]byte{
	'{': '}',
	'(': ')',
	'[': ']',
	'<': '>',
}

func main() {
	syntaxScore := 0
	var autocompleteScores []int
	s := util.ScanAll()
	for s.Scan() {
		if valid, bad, stack := parse(s.Text()); !valid {
			syntaxScore += points[bad][0]
			continue
		} else {
			sum := 0
			for i := len(stack) - 1; i >= 0; i-- {
				sum *= 5
				sum += points[stack[i]][1]
			}
			autocompleteScores = append(autocompleteScores, sum)
		}
	}
	fmt.Println(syntaxScore)
	sort.Ints(autocompleteScores)
	fmt.Println(autocompleteScores[len(autocompleteScores)/2])
}

func parse(s string) (valid bool, bad byte, stack []byte) {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[', '<':
			stack = append(stack, closing[s[i]])
		default:
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last != s[i] {
				return false, s[i], stack
			}
		}
	}
	return true, 0, stack
}

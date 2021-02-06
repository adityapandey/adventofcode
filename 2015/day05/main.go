package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var sum1, sum2 int
	s := util.ScanAll()
	for s.Scan() {
		if nice1(s.Text()) {
			sum1++
		}
		if nice2(s.Text()) {
			sum2++
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func nice1(s string) bool {
	var vowels int
	var double bool
	for i := range s {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		if i < len(s)-1 && s[i] == s[i+1] {
			double = true
		}
	}
	var disallowed bool
	for _, ss := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, ss) {
			disallowed = true
			break
		}
	}
	return vowels >= 3 && double && !disallowed
}

func nice2(s string) bool {
	pairs := make(map[string][]int)
	var palindrome bool
	for i := 0; i < len(s)-1; i++ {
		if i > 0 && s[i-1] == s[i+1] {
			palindrome = true
		}
		pairs[s[i:i+2]] = append(pairs[s[i:i+2]], i)
	}
	var twicePair bool
	for _, indices := range pairs {
		if len(indices) > 2 {
			twicePair = true
			break
		}
		if len(indices) == 1 {
			continue
		}
		if util.Abs(indices[0]-indices[1]) > 1 {
			twicePair = true
			break
		}
	}
	return twicePair && palindrome
}

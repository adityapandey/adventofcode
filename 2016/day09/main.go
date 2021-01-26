package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := util.ReadAll()
	b := []byte(input)
	fmt.Println(expandLen(b, false))
	fmt.Println(expandLen(b, true))
}

func expandLen(b []byte, recurse bool) int {
	var l int
	for i := 0; i < len(b); i++ {
		if b[i] == '(' {
			start := i
			for b[i] != ')' {
				i++
			}
			var offset, rep int
			fmt.Sscanf(string(b[start+1:i]), "%dx%d", &offset, &rep)
			if recurse {
				l += rep * expandLen(b[i+1:i+1+offset], true)
			} else {
				l += rep * offset
			}
			i += offset
		} else {
			l++
		}
	}
	return l
}

package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type triangle [3]int

func (t triangle) possible() bool {
	return t[0]+t[1] > t[2] &&
		t[1]+t[2] > t[0] &&
		t[2]+t[0] > t[1]
}

func main() {
	input := strings.Split(util.ReadAll(), "\n")

	var sum int
	for _, row := range input {
		var t triangle
		f := strings.Fields(row)
		for i := range f {
			t[i] = util.Atoi(f[i])
		}
		if t.possible() {
			sum++
		}
	}
	fmt.Println(sum)

	sum = 0
	for row := 0; row < len(input); row += 3 {
		var ts [3]triangle
		for j := 0; j < 3; j++ {
			f := strings.Fields(input[row+j])
			for i := range f {
				ts[i][j] = util.Atoi(f[i])
			}
		}
		for j := 0; j < 3; j++ {
			if ts[j].possible() {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	sum, sum2 := 0, 0
	for _, input := range strings.Split(util.ReadAll(), "\n\n") {
		rows := strings.Split(input, "\n")
		cols := make([]string, len(rows[0]))
		for i := 0; i < len(rows[0]); i++ {
			for j := 0; j < len(rows); j++ {
				cols[i] += string(rows[j][i])
			}
		}
		sum += score(cols) + 100*score(rows)
		sum2 += smudge(cols) + 100*smudge(rows)
	}
	fmt.Println(sum)
	fmt.Println(sum2)
}

func score(terrain []string) int {
	for i := 0; i < len(terrain)-1; i++ {
		mirror := false
		if terrain[i] == terrain[i+1] {
			mirror = true
			for j := 1; i-j >= 0 && i+1+j < len(terrain); j++ {
				if terrain[i-j] != terrain[i+1+j] {
					mirror = false
					break
				}
			}
		}
		if mirror {
			return i + 1
		}
	}
	return 0
}

func smudge(terrain []string) int {
	for i := range terrain {
		d := 0
		for j := 0; i-j >= 0 && i+1+j < len(terrain); j++ {
			d += diff(terrain[i-j], terrain[i+1+j])
		}
		if d == 1 {
			return i + 1
		}
	}
	return 0
}

func diff(a, b string) int {
	c := 0
	for i := range a {
		if a[i] != b[i] {
			c++
		}
	}
	return c
}

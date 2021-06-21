package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	s := util.ScanAll()
	var f, ff int
	for s.Scan() {
		mass := util.Atoi(s.Text())
		f += fuel(mass)
		ff += fuelfuel(mass)
	}
	fmt.Println(f)
	fmt.Println(ff)
}

func fuel(mass int) int {
	return mass/3 - 2
}

func fuelfuel(mass int) int {
	var ff int
	for {
		f := fuel(mass)
		if f <= 0 {
			break
		}
		ff += f
		mass = f
	}
	return ff
}

// https://adventofcode.com/2020/day/17
package main

import (
	"bufio"
	"fmt"
	"os"
)

type pt3 struct {
	X, Y, Z int
}

func (a pt3) add(b pt3) pt3 {
	return pt3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

type pt4 struct {
	X, Y, Z, W int
}

func (a pt4) add(b pt4) pt4 {
	return pt4{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

func main() {
	cube := make(map[pt3]byte)
	hypercube := make(map[pt4]byte)
	s := bufio.NewScanner(os.Stdin)
	y := 0
	for s.Scan() {
		line := s.Text()
		for x := 0; x < len(line); x++ {
			cube[pt3{x, y, 0}] = line[x]
			hypercube[pt4{x, y, 0, 0}] = line[x]
		}
		y++
	}

	// Part 1
	for i := 0; i < 6; i++ {
		cube = iterate3(cube)
	}
	numActive := 0
	for _, s := range cube {
		if s == '#' {
			numActive++
		}
	}
	fmt.Println(numActive)

	// Part 2
	for i := 0; i < 6; i++ {
		hypercube = iterate4(hypercube)
	}
	numActive = 0
	for _, s := range hypercube {
		if s == '#' {
			numActive++
		}
	}
	fmt.Println(numActive)

}

func iterate3(cube map[pt3]byte) map[pt3]byte {
	next := make(map[pt3]byte)
	for p := range cube {
		for _, n := range neighbours3(p) {
			next[n] = '.'
		}
	}

	for p := range next {
		numActive := 0
		for _, n := range neighbours3(p) {
			if cube[n] == '#' {
				numActive++
			}
		}
		if cube[p] == '#' && numActive == 2 || numActive == 3 {
			next[p] = '#'
		}
	}
	return next
}

func neighbours3(p pt3) []pt3 {
	var n []pt3
	near := []int{-1, 0, 1}
	for _, z := range near {
		for _, y := range near {
			for _, x := range near {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				n = append(n, p.add(pt3{x, y, z}))
			}
		}
	}
	return n
}

func iterate4(hypercube map[pt4]byte) map[pt4]byte {
	next := make(map[pt4]byte)
	for p := range hypercube {
		for _, n := range neighbours4(p) {
			next[n] = '.'
		}
	}

	for p := range next {
		numActive := 0
		for _, n := range neighbours4(p) {
			if hypercube[n] == '#' {
				numActive++
			}
		}
		if hypercube[p] == '#' && numActive == 2 || numActive == 3 {
			next[p] = '#'
		}
	}
	return next
}

func neighbours4(p pt4) []pt4 {
	var n []pt4
	near := []int{-1, 0, 1}
	for _, w := range near {
		for _, z := range near {
			for _, y := range near {
				for _, x := range near {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					n = append(n, p.add(pt4{x, y, z, w}))
				}
			}
		}
	}
	return n
}

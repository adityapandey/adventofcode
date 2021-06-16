package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

type grid map[image.Point]byte

func (g grid) copy() grid {
	r := make(grid)
	for p, b := range g {
		r[p] = b
	}
	return r
}

func (g grid) next() grid {
	r := g.copy()
	for p, b := range g {
		r[p] = b
		m := make(map[byte]int)
		for _, n := range util.Neighbors8 {
			m[g[p.Add(n)]]++
		}
		if b == '.' && m['|'] >= 3 {
			r[p] = '|'
		}
		if b == '|' && m['#'] >= 3 {
			r[p] = '#'
		}
		if b == '#' {
			if m['#'] >= 1 && m['|'] >= 1 {
				r[p] = '#'
			} else {
				r[p] = '.'
			}
		}
	}
	return r
}

func main() {
	g := make(grid)
	s := util.ScanAll()
	var y int
	for s.Scan() {
		t := s.Text()
		for x := 0; x < len(t); x++ {
			g[image.Pt(x, y)] = t[x]
		}
		y++
	}

	// Part 1
	gc := g.copy()
	for i := 0; i < 10; i++ {
		gc = gc.next()
	}
	var nTrees, nLumberyards int
	for _, b := range gc {
		switch b {
		case '|':
			nTrees++
		case '#':
			nLumberyards++

		}
	}
	fmt.Println(nLumberyards * nTrees)

	// Part 2
	// Pattern repeats
	gc = g.copy()
	iterMap := make(map[int]int)
	revIterMap := make(map[int]int)
	var repeatStart, repeatPeriod int
	for i := 0; ; i++ {
		gc = gc.next()
		var nTrees, nLumberyards int
		for _, b := range gc {
			switch b {
			case '|':
				nTrees++
			case '#':
				nLumberyards++
			}
		}
		iterMap[i] = nLumberyards * nTrees
		if prev, ok := revIterMap[iterMap[i]]; ok && prev < i-1 {
			repeat := true
			for j := i - 1; j > prev; j-- {
				if prev-(i-j) > 0 && iterMap[j] != iterMap[prev-(i-j)] {
					repeat = false
					break
				}
			}
			if repeat {
				repeatStart = prev
				repeatPeriod = i - prev
				break
			}
		}
		revIterMap[iterMap[i]] = i
	}
	fmt.Println(iterMap[repeatStart+((999999999-repeatStart)%repeatPeriod)])
}

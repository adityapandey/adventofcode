package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type World struct {
	grid [][]byte
}

func (w *World) Init(b [][]byte) {
	w.grid = make([][]byte, len(b[0]))
	for x := 0; x < len(b[0]); x++ {
		w.grid[x] = make([]byte, len(b))
		for y := 0; y < len(b); y++ {
			w.grid[x][y] = b[y][x]
		}
	}
}

func (w World) String() string {
	var s strings.Builder
	for i := 0; i < len(w.grid); i++ {
		for j := 0; j < len(w.grid[i]); j++ {
			fmt.Fprintf(&s, "%c", w.grid[i][j])
		}
		fmt.Fprintln(&s)
	}
	return s.String()
}

func (w *World) neighbors(x, y int) map[byte]int {
	m := make(map[byte]int)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i > 49 || y+j < 0 || y+j > 49 {
				continue
			}
			m[w.grid[x+i][y+j]]++
		}
	}
	//fmt.Println(x, y, m)
	return m
}

func (w *World) Evolve() {
	next := make([][]byte, len(w.grid))
	for i := range w.grid {
		next[i] = make([]byte, len(w.grid[i]))
		copy(next[i], w.grid[i])
	}
	for i := 0; i < len(w.grid); i++ {
		for j := 0; j < len(w.grid[i]); j++ {
			m := w.neighbors(i, j)
			if w.grid[i][j] == '.' && m['|'] >= 3 {
				next[i][j] = '|'
			}
			if w.grid[i][j] == '|' && m['#'] >= 3 {
				next[i][j] = '#'
			}
			if w.grid[i][j] == '#' {
				if m['#'] >= 1 && m['|'] >= 1 {
					next[i][j] = '#'
				} else {
					next[i][j] = '.'
				}
			}
		}
	}
	for i := range w.grid {
		copy(w.grid[i], next[i])
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var input [][]byte
	for s.Scan() {
		input = append(input, s.Bytes())
	}
	var w World
	w.Init(input)
	// fmt.Print("\033[H\033[2J")
	// fmt.Println(w)
	for i := 0; i < 10; i++ {
		w.Evolve()
		// time.Sleep(100 * time.Millisecond)
		// fmt.Print("\033[H\033[2J")
		// fmt.Println(w)
	}

	// Part 1
	tree, lumberyard := 0, 0
	for i := 0; i < len(w.grid); i++ {
		for j := 0; j < len(w.grid[i]); j++ {
			switch w.grid[i][j] {
			case '|':
				tree++
			case '#':
				lumberyard++
			}
		}
	}
	fmt.Println(lumberyard * tree)

	// Part 2
	// Pattern repeats
	w.Init(input)
	iterMap := make(map[int]int)
	revIterMap := make(map[int]int)
	var repeatStart, repeatPeriod int
	for i := 0; ; i++ {
		w.Evolve()
		tree, lumberyard := 0, 0
		for i := 0; i < len(w.grid); i++ {
			for j := 0; j < len(w.grid[i]); j++ {
				switch w.grid[i][j] {
				case '|':
					tree++
				case '#':
					lumberyard++
				}
			}
		}
		iterMap[i] = lumberyard * tree
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

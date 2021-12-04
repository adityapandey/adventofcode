package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type board [5][5]struct {
	val    int
	marked bool
}

func (b *board) mark(val int) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if (*b)[row][col].val == val {
				(*b)[row][col].marked = true
				return
			}
		}
	}
}

func (b board) won() bool {
	wonRow := false
	for row := 0; row < 5; row++ {
		won := true
		for col := 0; col < 5; col++ {
			won = won && b[row][col].marked
		}
		wonRow = wonRow || won
	}
	wonCol := false
	for col := 0; col < 5; col++ {
		won := true
		for row := 0; row < 5; row++ {
			won = won && b[row][col].marked
		}
		wonCol = wonCol || won
	}
	return wonRow || wonCol
}

func main() {
	s := strings.Split(util.ReadAll(), "\n\n")
	var ns []int
	for _, n := range strings.Split(s[0], ",") {
		ns = append(ns, util.Atoi(n))
	}
	var boards []board
	for i := 1; i < len(s); i++ {
		var b board
		for row, line := range strings.Split(s[i], "\n") {
			for col, n := range strings.Fields(line) {
				b[row][col].val = util.Atoi(n)
			}
		}
		boards = append(boards, b)
	}

	winners := map[int]struct{}{}
	for i := 0; len(winners) < len(boards) && i < len(ns); i++ {
		for j := range boards {
			boards[j].mark(ns[i])
			if _, ok := winners[j]; !ok && boards[j].won() {
				winners[j] = struct{}{}
				if len(winners) == 1 || len(winners) == len(boards) {
					var sum int
					for row := 0; row < 5; row++ {
						for col := 0; col < 5; col++ {
							if !boards[j][row][col].marked {
								sum += boards[j][row][col].val
							}
						}
					}
					fmt.Println(sum * ns[i])
				}
			}
		}
	}
}

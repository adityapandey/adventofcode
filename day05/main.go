// https://adventofcode.com/2020/day/5
package main

import (
	"bufio"
	"fmt"
	"os"
)

type seat struct {
	row, col byte
}

func main() {
	var seats []seat
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		row := parse(s.Text()[:7])
		col := parse(s.Text()[7:])
		seats = append(seats, seat{row, col})
	}

	// Part 1
	var max uint
	for _, se := range seats {
		seatID := 8*uint(se.row) + uint(se.col)
		if seatID > max {
			max = seatID
		}
	}
	fmt.Println(max)

	// Part 2
	m := make(map[uint]struct{})
	for _, se := range seats {
		seatID := 8*uint(se.row) + uint(se.col)
		m[seatID] = struct{}{}
	}

	for seatID := uint(1); seatID < uint(1<<10)-1; seatID++ {
		_, prevOk := m[seatID-1]
		_, ok := m[seatID]
		_, nextOk := m[seatID+1]

		if !ok && prevOk && nextOk {
			fmt.Println(seatID)
			return
		}
	}

}

func parse(s string) byte {
	var sum, pow byte = 0, 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'B' || s[i] == 'R' {
			sum += pow
		}
		pow *= 2
	}
	return sum
}

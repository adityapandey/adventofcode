package main

import (
	"bufio"
	"fmt"
	"os"
)

type Claim struct {
	id, x, y, w, h int
}

func (c *Claim) hasOverlap(grid *[1000][1000]int) bool {
	for i := 0; i < c.w; i++ {
		for j := 0; j < c.h; j++ {
			if (*grid)[c.x+i][c.y+j] != 1 {
				return true
			}
		}
	}
	return false
}

func main() {
	var claims []Claim
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var claim Claim
		fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &claim.id, &claim.x, &claim.y, &claim.w, &claim.h)
		claims = append(claims, claim)
	}

	var grid [1000][1000]int

	for _, c := range claims {
		for i := 0; i < c.w; i++ {
			for j := 0; j < c.h; j++ {
				grid[c.x+i][c.y+j]++
			}
		}
	}

	//Part 1
	overlap := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] > 1 {
				overlap++
			}
		}
	}

	fmt.Println(overlap)

	// Part 2
	for _, c := range claims {
		if !c.hasOverlap(&grid) {
			fmt.Println(c.id)
			return
		}
	}
}

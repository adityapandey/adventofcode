package main

import (
	"fmt"
	"slices"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

type brick struct {
	start, end util.Pt3
}

func main() {
	var bricks []brick
	s := util.ScanAll()
	for s.Scan() {
		var b brick
		fmt.Sscanf(s.Text(), "%d,%d,%d~%d,%d,%d", &b.start.X, &b.start.Y, &b.start.Z, &b.end.X, &b.end.Y, &b.end.Z)
		bricks = append(bricks, b)
	}
	sort.Slice(bricks, func(i, j int) bool { return bricks[i].start.Z < bricks[j].start.Z })

	moves := -1
	for moves != 0 {
		bricks, moves = fall(bricks)
	}

	var part1, part2 int
	for i := range bricks {
		bricksCopy := slices.Clone(bricks)
		bricksCopy = append(bricksCopy[:i], bricksCopy[i+1:]...)
		_, moves := fall(bricksCopy)
		if moves == 0 {
			part1++
		}
		part2 += moves
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func fall(bricks []brick) ([]brick, int) {
	var moves int
	bricksCopy := slices.Clone(bricks)
loop:
	for i, b := range bricksCopy {
		if b.start.Z == 1 {
			continue
		}
		b.start.Z--
		b.end.Z--
		for j := 0; j < i; j++ {
			if collides(b, bricksCopy[j]) {
				continue loop
			}
		}
		bricksCopy[i] = b
		moves++
	}
	return bricksCopy, moves
}

func collides(a, b brick) bool {
	return a.start.X <= b.end.X && a.end.X >= b.start.X &&
		a.start.Y <= b.end.Y && a.end.Y >= b.start.Y &&
		a.start.Z <= b.end.Z && a.end.Z >= b.start.Z
}

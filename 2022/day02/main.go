package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

type choice int

const (
	ROCK choice = iota
	PAPER
	SCISSORS
)

type outcome int

const (
	LOSS outcome = iota
	DRAW
	WIN
)

func main() {
	score, score2 := 0, 0
	s := util.ScanAll()
	for s.Scan() {
		var opp, you byte
		fmt.Sscanf(s.Text(), "%c %c", &opp, &you)
		score += round(opp, you)
		score2 += round2(opp, you)
	}
	fmt.Println(score)
	fmt.Println(score2)
}

func round(opp, you byte) int {
	oppchoice, yourchoice := choice(opp-'A'), choice(you-'X')
	var o outcome
	switch {
	case yourchoice == oppchoice:
		o = DRAW
	case yourchoice == (oppchoice+1)%3:
		o = WIN
	default:
		o = LOSS
	}
	return int(yourchoice) + 1 + 3*int(o)
}

func round2(opp, you byte) int {
	oppchoice, o := choice(opp-'A'), outcome(you-'X')
	var yourchoice choice
	switch o {
	case DRAW:
		yourchoice = oppchoice
	case WIN:
		yourchoice = (oppchoice + 1) % 3
	default:
		yourchoice = (oppchoice + 2) % 3
	}
	return int(yourchoice) + 1 + 3*int(o)
}

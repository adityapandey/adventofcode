package main

import (
	"fmt"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var pos1, pos2 int
	fmt.Fscanf(os.Stdin, "Player 1 starting position: %d", &pos1)
	fmt.Fscanf(os.Stdin, "Player 2 starting position: %d", &pos2)

	fmt.Println(trials(pos1, pos2))

	fmt.Println(diracDiceWins(state{pos1, 0, pos2, 0}))
}

type trialDice struct {
	n     int
	rolls int
}

func (d *trialDice) roll() int {
	d.rolls++
	d.n++
	if d.n > 100 {
		d.n = 1
	}
	return d.n
}

func trials(pos1, pos2 int) int {
	var score1, score2, loser int
	var d trialDice
	for {
		pos1 = ((pos1 - 1 + d.roll() + d.roll() + d.roll()) % 10) + 1
		score1 += pos1
		if score1 >= 1000 {
			loser = score2
			break
		}
		pos2 = ((pos2 - 1 + d.roll() + d.roll() + d.roll()) % 10) + 1
		score2 += pos2
		if score2 >= 1000 {
			loser = score1
			break
		}
	}
	return loser * d.rolls
}

type state struct {
	pos1, score1 int
	pos2, score2 int
}

func diracDiceWins(s state) int {
	var wins1, wins2 int
	states := map[state]int{s: 1}
	for len(states) > 0 {
		var w1, w2 int
		states, w1, w2 = next(states)
		wins1 += w1
		wins2 += w2
	}
	return util.Max(wins1, wins2)
}

func next(m map[state]int) (map[state]int, int, int) {
	nextState := map[state]int{}
	var wins1, wins2 int
	for s, v := range m {
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				for k := 1; k <= 3; k++ {
					var ss state
					ss.pos1 = ((s.pos1 - 1 + i + j + k) % 10) + 1
					ss.score1 = s.score1 + ss.pos1
					if ss.score1 >= 21 {
						wins1 += v
						continue
					}
					for a := 1; a <= 3; a++ {
						for b := 1; b <= 3; b++ {
							for c := 1; c <= 3; c++ {
								ss.pos2 = ((s.pos2 - 1 + a + b + c) % 10) + 1
								ss.score2 = s.score2 + ss.pos2
								if ss.score2 >= 21 {
									wins2 += v
									continue
								}
								nextState[ss] += v
							}
						}
					}
				}
			}
		}
	}
	return nextState, wins1, wins2
}

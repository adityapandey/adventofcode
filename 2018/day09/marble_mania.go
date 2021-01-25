package main

import (
	"container/ring"
	"fmt"
	"os"
)

func maxScore(numPlayers, iterations int) int {
	players := ring.New(numPlayers)
	scores := make([]int, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players.Value = i
		players = players.Next()
	}

	game := ring.New(1)
	game.Value = 0

	for i := 1; i <= iterations; i++ {
		player := players.Value.(int)
		if i%23 != 0 {
			newValue := ring.New(1)
			newValue.Value = i
			game = game.Prev()
			game.Link(newValue)
		} else {
			scores[player] += i
			game = game.Move(7)
			scores[player] += game.Unlink(1).Value.(int)
			game = game.Prev()
		}
		players = players.Next()
	}

	var maxScore int
	for _, s := range scores {
		if s > maxScore {
			maxScore = s
		}
	}
	return maxScore
}

func main() {
	var numPlayers, iterations int
	fmt.Fscanf(os.Stdin, "%d players; last marble is worth %d points", &numPlayers, &iterations)
	// Part 1
	fmt.Println(maxScore(numPlayers, iterations))
	// Part 2
	fmt.Println(maxScore(numPlayers, iterations*100))
}

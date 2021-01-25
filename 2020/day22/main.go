// https://adventofcode.com/2020/day/22
package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	players := strings.Split(string(input), "\n\n")
	var decks [2][]byte
	for i := 0; i < 2; i++ {
		for _, s := range strings.Split(players[i], "\n")[1:] {
			decks[i] = append(decks[i], byte(util.Atoi(s)))
		}
	}

	// Part 1
	_, deck := combat(decks, false)
	fmt.Println(score(deck))

	// Part 2
	_, deck = combat(decks, true)
	fmt.Println(score(deck))
}

func combat(decks [2][]byte, recurse bool) (winner int, deck []byte) {
	seen := make(map[uint32]struct{})
	for len(decks[0]) > 0 && len(decks[1]) > 0 {
		h := hash(decks)
		if _, ok := seen[h]; ok {
			return 0, decks[0]
		}
		seen[h] = struct{}{}
		card0, card1 := decks[0][0], decks[1][0]
		var winner int
		if int(card0) < len(decks[0]) && int(card1) < len(decks[1]) && recurse {
			deck0 := make([]byte, card0)
			deck1 := make([]byte, card1)
			copy(deck0, decks[0][1:])
			copy(deck1, decks[1][1:])
			winner, _ = combat([2][]byte{deck0, deck1}, true)
		} else {
			if card0 > card1 {
				winner = 0
			} else {
				winner = 1
			}
		}
		if winner == 0 {
			decks[0] = append(decks[0][1:], card0, card1)
			decks[1] = decks[1][1:]
		} else {
			decks[0] = decks[0][1:]
			decks[1] = append(decks[1][1:], card1, card0)
		}
	}
	if len(decks[0]) == 0 {
		return 1, decks[1]
	}
	return 0, decks[0]
}

var crc = crc32.NewIEEE()

func hash(decks [2][]byte) uint32 {
	crc.Reset()
	crc.Write(decks[0])
	crc.Write([]byte{100}) // 100 does not occur in input
	crc.Write(decks[1])
	return crc.Sum32()
}

func score(deck []byte) int {
	score := 0
	l := len(deck)
	for i, c := range deck {
		score += (l - i) * int(c)
	}
	return score
}

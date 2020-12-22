// https://adventofcode.com/2020/day/22
package main

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode2020-go/util"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	players := strings.Split(string(input), "\n\n")
	var decks [2][]int
	for i := 0; i < 2; i++ {
		for _, s := range strings.Split(players[i], "\n")[1:] {
			decks[i] = append(decks[i], util.Atoi(s))
		}
	}

	// Part 1
	_, deck := combat(decks, false)
	fmt.Println(score(deck))

	// Part 2
	_, deck = combat(decks, true)
	fmt.Println(score(deck))
}

var crc = crc32.NewIEEE()

func combat(decks [2][]int, recurse bool) (int, []int) {
	seen := make(map[uint32]struct{})
	for len(decks[0]) > 0 && len(decks[1]) > 0 {
		h := hash(decks)
		if _, ok := seen[h]; ok {
			return 0, decks[0]
		}
		seen[h] = struct{}{}
		card0, card1 := decks[0][0], decks[1][0]
		var winner int
		if card0 < len(decks[0]) && card1 < len(decks[1]) && recurse {
			deck0 := make([]int, card0)
			deck1 := make([]int, card1)
			copy(deck0, decks[0][1:])
			copy(deck1, decks[1][1:])
			winner, _ = combat([2][]int{deck0, deck1}, recurse)
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

func hash(decks [2][]int) uint32 {
	var b bytes.Buffer
	for i := 0; i < 2; i++ {
		b.WriteByte(byte(100))
		for _, c := range decks[i] {
			b.WriteByte(byte(c))
		}
	}
	crc.Reset()
	crc.Write(b.Bytes())
	return crc.Sum32()
}

func score(a []int) int {
	score := 0
	l := len(a)
	for i, n := range a {
		score += (l - i) * n
	}
	return score
}

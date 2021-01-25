// https://adventofcode.com/2020/day/25
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adityapandey/adventofcode2020-go/util"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	cardKey := util.Atoi(s.Text())
	s.Scan()
	doorKey := util.Atoi(s.Text())
	cardLoopSize := findLoopSize(cardKey)
	fmt.Println(transform(doorKey, cardLoopSize))
}

func findLoopSize(n int) int {
	i, curr, subject := 0, 1, 7
	for ; curr != n; i++ {
		curr *= subject
		curr %= 20201227
	}
	return i
}

func transform(subject, loopSize int) int {
	curr := 1
	for i := 0; i < loopSize; i++ {
		curr *= subject
		curr %= 20201227
	}
	return curr
}

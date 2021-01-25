package main

import (
	"bufio"
	"fmt"
	"os"
)

func TwosAndThrees(s string) (int, int) {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	var twos, threes int
	for _, v := range m {
		if v == 2 {
			twos = 1
		}
		if v == 3 {
			threes = 1
		}
	}
	return twos, threes
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var ids []string
	for s.Scan() {
		ids = append(ids, s.Text())
	}

	// Part 1
	var sumTwos, sumThrees int
	for _, id := range ids {
		twos, threes := TwosAndThrees(id)
		sumTwos += twos
		sumThrees += threes
	}
	fmt.Println(sumTwos * sumThrees)

	// Part 2

	type Pair struct{ a, b string }
	type E struct{}
	m := make(map[Pair]E)

	l := len(ids[0])
	for _, id := range ids {
		for i := 1; i < l-1; i++ {
			if _, ok := m[Pair{id[:i], id[i+1:]}]; ok {
				fmt.Printf("%s%s\n", id[:i], id[i+1:])
				return
			} else {
				m[Pair{id[:i], id[i+1:]}] = E{}
			}
		}
	}
}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	moves := strings.Split(string(input), ",")
	l := 16
	start := make([]byte, l)
	for i := 0; i < l; i++ {
		start[i] = byte('a' + i)
	}
	next := dance(string(start), l, moves)
	fmt.Println(string(next))

	// Assuming there's a cycle.
	seen := map[string]struct{}{string(start): {}}
	seenOrdered := []string{string(start)}
	for {
		if _, ok := seen[next]; ok {
			break
		} else {
			seen[next] = struct{}{}
			seenOrdered = append(seenOrdered, next)
		}
		next = dance(next, l, moves)
	}
	lastSeen := 0
	for seenOrdered[lastSeen] != string(next) {
		lastSeen++
	}
	fmt.Println(seenOrdered[lastSeen+(1000000000-lastSeen)%(len(seenOrdered)-lastSeen)])
}

func dance(s string, l int, moves []string) string {
	list := []byte(s)
	for _, m := range moves {
		switch m[0] {
		case 's':
			var size int
			fmt.Sscanf(m[1:], "%d", &size)
			list = append(list[l-size:], list[:l-size]...)
		case 'x':
			var i, j int
			fmt.Sscanf(m[1:], "%d/%d", &i, &j)
			list[i], list[j] = list[j], list[i]
		case 'p':
			var a, b byte
			fmt.Sscanf(m[1:], "%c/%c", &a, &b)
			i, j := bytes.IndexByte(list, a), bytes.IndexByte(list, b)
			list[i], list[j] = list[j], list[i]
		}
	}
	return string(list)
}

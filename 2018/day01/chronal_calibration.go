package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	var freqChanges []int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		freqChanges = append(freqChanges, atoi(s.Text()))
	}

	// Part 1
	freq := 0
	for _, c := range freqChanges {
		freq += c
	}
	fmt.Println(freq)

	// Part 2
	freq = 0
	freqSet := make(map[int]struct{})
	freqSet[freq] = struct{}{}
	found := false
	for !found {
		for _, change := range freqChanges {
			freq += change
			if _, found = freqSet[freq]; found {
				fmt.Println(freq)
				break
			}
			freqSet[freq] = struct{}{}
		}
	}
}

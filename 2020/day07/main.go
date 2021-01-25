// https://adventofcode.com/2020/day/7
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var contentsRe = regexp.MustCompile(`(\d+) (\w+ \w+) bags?(, )?`)

type bag struct {
	color    string
	contents map[string]int
}

func main() {
	bags := make(map[string]bag)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		b := parseBag(s.Text())
		bags[b.color] = b
	}

	// Part 1
	sum := 0
	for _, b := range bags {
		if containsGold(b.color, bags) {
			sum++
		}
	}
	fmt.Println(sum)

	// Part 2
	fmt.Println(expandBags("shiny gold", bags))
}

func parseBag(s string) bag {
	var b bag
	split := strings.Split(s, " bags contain ")
	b.color = split[0]
	b.contents = make(map[string]int)
	for _, match := range contentsRe.FindAllStringSubmatch(split[1], -1) {
		b.contents[match[2]] = util.Atoi(match[1])
	}
	return b
}

var memoizeContainsGold = make(map[string]bool)

func containsGold(color string, bags map[string]bag) bool {
	if contains, ok := memoizeContainsGold[color]; ok {
		return contains
	}
	for c := range bags[color].contents {
		if c == "shiny gold" || containsGold(c, bags) {
			memoizeContainsGold[color] = true
			return true
		}
	}
	memoizeContainsGold[color] = false
	return false
}

var memoizeExpandBags = make(map[string]int)

func expandBags(color string, bags map[string]bag) int {
	if contains, ok := memoizeExpandBags[color]; ok {
		return contains
	}
	sum := 0
	for c, v := range bags[color].contents {
		sum += v * (1 + expandBags(c, bags))
	}
	memoizeExpandBags[color] = sum
	return sum
}

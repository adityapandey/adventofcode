package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

const ticker = `children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`

func main() {
	sues := make(map[int]map[string]int)
	s := util.ScanAll()
	for s.Scan() {
		line := s.Text()
		colon := strings.IndexByte(s.Text(), ':')
		var i int
		fmt.Sscanf(line[:colon], "Sue %d", &i)
		sues[i] = parseSue(strings.Split(line[colon+2:], ", "))
	}
	giftSue := parseSue(strings.Split(ticker, "\n"))

	suesCopy := copymap(sues)
	for i := range suesCopy {
		if !matches1(giftSue, suesCopy[i]) {
			delete(suesCopy, i)
		}
	}
	fmt.Println(only(suesCopy))

	for i := range sues {
		if !matches2(giftSue, sues[i]) {
			delete(sues, i)
		}
	}
	fmt.Println(only(sues))

}

func parseSue(a []string) map[string]int {
	sue := make(map[string]int)
	for _, s := range a {
		matches := strings.Split(s, ": ")
		sue[matches[0]] = util.Atoi(matches[1])
	}
	return sue
}

func matches1(a, b map[string]int) bool {
	for k, v := range a {
		if vv, ok := b[k]; ok && vv != v {
			return false
		}
	}
	return true
}

func matches2(a, b map[string]int) bool {
	for k, v := range a {
		vv, ok := b[k]
		if !ok {
			continue
		}
		switch k {
		case "cats", "trees":
			if vv <= v {
				return false
			}
		case "pomeranians", "goldfish":
			if vv >= v {
				return false
			}
		default:
			if vv != v {
				return false
			}
		}
	}
	return true
}

func copymap(m map[int]map[string]int) map[int]map[string]int {
	r := make(map[int]map[string]int)
	for k, v := range m {
		r[k] = v
	}
	return r
}

func only(m map[int]map[string]int) int {
	if len(m) != 1 {
		log.Fatal(len(m))
	}
	var ks []int
	for k := range m {
		ks = append(ks, k)
	}
	return ks[0]
}

package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

type node struct {
	id   string
	l, r string
}

var re = regexp.MustCompile(`(...) = \((...), (...+)\)`)

func main() {
	sp := strings.Split(util.ReadAll(), "\n\n")
	path := sp[0]
	nodes := map[string]node{}
	for _, line := range strings.Split(sp[1], "\n") {
		m := re.FindStringSubmatch(line)
		nodes[m[1]] = node{m[1], m[2], m[3]}
	}

	fmt.Println(findZ("AAA", nodes, path, func(curr string) bool { return curr == "ZZZ" }))

	var ghosts []string
	for _, n := range maps.Keys(nodes) {
		if n[2] == 'A' {
			ghosts = append(ghosts, n)
		}
	}

	lcm := findZ(ghosts[0], nodes, path, func(curr string) bool { return curr[2] == 'Z' })
	for _, ghost := range ghosts[1:] {
		lcm = util.Lcm(lcm, findZ(ghost, nodes, path, func(curr string) bool { return curr[2] == 'Z' }))
	}
	fmt.Println(lcm)
}

func findZ(curr string, nodes map[string]node, path string, cond func(curr string) bool) int {
	c := 0
	for i := 0; !cond(curr); i, c = (i+1)%len(path), c+1 {
		if path[i] == 'L' {
			curr = nodes[curr].l
		} else {
			curr = nodes[curr].r
		}
	}
	return c
}

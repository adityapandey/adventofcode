package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type part [4]int // x, m, a, s

type workflow struct {
	id    string
	rules []rule
	next  string
}

func (w workflow) run(p part) string {
	for _, r := range w.rules {
		if next, ok := r.run(p); ok {
			return next
		}
	}
	return w.next
}

type rule struct {
	xmas int
	gt   bool
	n    int
	next string
}

func (r rule) run(p part) (string, bool) {
	if (r.gt && p[r.xmas] > r.n) || (!r.gt && p[r.xmas] < r.n) {
		return r.next, true
	}
	return "", false
}

func main() {
	workflows := map[string]workflow{}
	// input := util.ReadAll()
	in, _ := os.ReadFile("input")
	input := string(in)
	sp := strings.Split(input, "\n\n")
	for _, s := range strings.Split(sp[0], "\n") {
		var w workflow
		w.id = s[:strings.IndexByte(s, '{')]
		rr := strings.Split(s[strings.IndexByte(s, '{')+1:], ",")
		next := rr[len(rr)-1]
		w.next = next[:strings.IndexByte(next, '}')]
		for i := 0; i < len(rr)-1; i++ {
			var r rule
			r.xmas = strings.IndexByte("xmas", rr[i][0])
			if rr[i][1] == '>' {
				r.gt = true
			}
			r.n = util.Atoi(rr[i][2:strings.IndexByte(rr[i], ':')])
			r.next = rr[i][strings.IndexByte(rr[i], ':')+1:]
			w.rules = append(w.rules, r)
		}
		workflows[w.id] = w
	}
	var parts []part
	for _, s := range strings.Split(sp[1], "\n") {
		s = s[:len(s)-1]
		var p part
		for i, ss := range strings.Split(s, ",") {
			p[i] = util.Atoi(strings.Split(ss, "=")[1])
		}
		parts = append(parts, p)
	}

	sum := 0
	for _, p := range parts {
		w := "in"
		for ; w != "A" && w != "R"; w = workflows[w].run(p) {
		}
		if w == "A" {
			sum += p[0] + p[1] + p[2] + p[3]
		}
	}
	fmt.Println(sum)
	fmt.Println(combinations("in", [4]rng{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}, workflows))
}

type rng struct {
	// min inclusive, max exclusive
	min, max int
}

func combinations(start string, rngs [4]rng, workflows map[string]workflow) int {
	if start == "R" {
		return 0
	}
	if start == "A" {
		prod := 1
		for _, r := range rngs {
			prod *= r.max + 1 - r.min
		}
		return prod
	}

	sum := 0
	for _, r := range workflows[start].rules {
		if r.gt {
			uBound := util.Min(rngs[r.xmas].max, r.n)
			lBound := util.Max(rngs[r.xmas].min, r.n+1)
			nextRngs := rngs
			nextRngs[r.xmas].min = lBound
			rngs[r.xmas].max = uBound
			sum += combinations(r.next, nextRngs, workflows)
		} else {
			lBound := util.Min(rngs[r.xmas].max, r.n-1)
			uBound := util.Max(rngs[r.xmas].min, r.n)
			nextRngs := rngs
			nextRngs[r.xmas].max = lBound
			rngs[r.xmas].min = uBound
			sum += combinations(r.next, nextRngs, workflows)
		}
	}
	sum += combinations(workflows[start].next, rngs, workflows)
	return sum
}

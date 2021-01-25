package main

import (
	"bufio"
	"fmt"
	"os"
)

type P [4]int

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distance(a, b P) int {
	s := 0
	for i := 0; i < 4; i++ {
		s += abs(a[i] - b[i])
	}
	return s
}

type Star struct {
	P
	c *Constellation
}

type Constellation map[*Star]struct{}

func (c *Constellation) Add(s *Star) {
	if s.c == c {
		return
	}
	delete(*s.c, s)
	for star := range *s.c {
		star.c = c
		(*c)[star] = struct{}{}
		delete(*s.c, star)
	}
	s.c = c
	(*c)[s] = struct{}{}
}

func main() {
	var stars []*Star
	var constellations []Constellation
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var p P
		fmt.Sscanf(s.Text(), "%d,%d,%d,%d", &p[0], &p[1], &p[2], &p[3])
		var star Star
		star.P = p
		m := make(Constellation)
		m[&star] = struct{}{}
		star.c = &m
		stars = append(stars, &star)
		constellations = append(constellations, m)
	}

	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {
			if distance(stars[i].P, stars[j].P) <= 3 {
				stars[i].c.Add(stars[j])
			}
		}
	}

	uniq := 0
	for _, c := range constellations {
		if len(c) > 0 {
			uniq++
		}
	}
	fmt.Println(uniq)
}

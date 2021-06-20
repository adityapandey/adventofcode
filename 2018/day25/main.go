package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

type Pt4 [4]int

func manhattan4(a, b Pt4) int {
	s := 0
	for i := 0; i < 4; i++ {
		s += util.Abs(a[i] - b[i])
	}
	return s
}

type star struct {
	p Pt4
	c *constellation
}

type constellation map[*star]struct{}

func (c *constellation) add(s *star) {
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
	var stars []*star
	var constellations []constellation
	s := util.ScanAll()
	for s.Scan() {
		var p Pt4
		fmt.Sscanf(s.Text(), "%d,%d,%d,%d", &p[0], &p[1], &p[2], &p[3])
		st := star{p, nil}
		m := make(constellation)
		m[&st] = struct{}{}
		st.c = &m
		stars = append(stars, &st)
		constellations = append(constellations, m)
	}

	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {
			if manhattan4(stars[i].p, stars[j].p) <= 3 {
				stars[i].c.add(stars[j])
			}
		}
	}

	var uniq int
	for _, c := range constellations {
		if len(c) > 0 {
			uniq++
		}
	}
	fmt.Println(uniq)
}

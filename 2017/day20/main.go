package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

type particle struct {
	id      int
	p, v, a util.Pt3
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var particles []particle
	c := 0
	for s.Scan() {
		p := particle{id: c}
		fmt.Sscanf(s.Text(), "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &p.p.X, &p.p.Y, &p.p.Z, &p.v.X, &p.v.Y, &p.v.Z, &p.a.X, &p.a.Y, &p.a.Z)
		particles = append(particles, p)
		c++
	}
	sort.Slice(particles,
		func(i, j int) bool {
			ai := util.Manhattan3(particles[i].a, util.Pt3{0, 0, 0})
			aj := util.Manhattan3(particles[j].a, util.Pt3{0, 0, 0})
			vi := util.Manhattan3(particles[i].v, util.Pt3{0, 0, 0})
			vj := util.Manhattan3(particles[j].v, util.Pt3{0, 0, 0})
			pi := util.Manhattan3(particles[i].p, util.Pt3{0, 0, 0})
			pj := util.Manhattan3(particles[j].p, util.Pt3{0, 0, 0})
			if ai == aj {
				if vi == vj {
					return pi < pj
				}
				return vi < vj
			}
			return ai < aj
		})
	fmt.Println(particles[0].id)

	// Randomly assumed that 1000 steps should be enough.
	for t := 0; t < 1000; t++ {
		m := make(map[util.Pt3]int)
		collisions := make(map[int]struct{})
		for i := range particles {
			particles[i].v = particles[i].v.Add(particles[i].a)
			particles[i].p = particles[i].p.Add(particles[i].v)
			if j, ok := m[particles[i].p]; ok {
				collisions[particles[i].id] = struct{}{}
				collisions[particles[j].id] = struct{}{}
			} else {
				m[particles[i].p] = i
			}
		}
		for i := 0; i < len(particles); i++ {
			if _, ok := collisions[particles[i].id]; ok {
				particles = append(particles[:i], particles[i+1:]...)
				i--
			}
		}
	}
	fmt.Println(len(particles))
}

func ids(p []particle, n int) []int {
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, p[i].id)
	}
	return a
}

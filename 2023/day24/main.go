package main

import (
	"fmt"
	"slices"

	"github.com/adityapandey/adventofcode/util"
)

type hail struct {
	pos util.Pt3
	vel util.Pt3
}

func main() {
	var hails []hail
	s := util.ScanAll()
	for s.Scan() {
		var h hail
		fmt.Sscanf(s.Text(), "%d, %d, %d @ %d, %d, %d", &h.pos.X, &h.pos.Y, &h.pos.Z, &h.vel.X, &h.vel.Y, &h.vel.Z)
		hails = append(hails, h)
	}
	var c int
	for i := 0; i < len(hails)-1; i++ {
		for j := i + 1; j < len(hails); j++ {
			x, y, parallel := intersection(hails[i], hails[j])
			if !parallel && x >= 2e14 && y >= 2e14 && x <= 4e14 && y <= 4e14 && isFuture(hails[i], hails[j], x) {
				c++
			}
		}
	}
	fmt.Println(c)

	rockVel := getRockVelocity(hails)
	x, y, _ := intersection(
		hail{
			pos: hails[0].pos,
			vel: hails[0].vel.Add(rockVel.Mul(-1)),
		}, hail{
			pos: hails[1].pos,
			vel: hails[1].vel.Add(rockVel.Mul(-1)),
		})
	t := float64(x-float64(hails[0].pos.X)) / float64(hails[0].vel.X-rockVel.X)
	z := float64(hails[0].pos.Z) + t*float64(hails[0].vel.Z-rockVel.Z)
	fmt.Println(int(x + y + z))
}

func isFuture(h1, h2 hail, x float64) bool {
	t1 := (x - float64(h1.pos.X)) / float64(h1.vel.X)
	t2 := (x - float64(h2.pos.X)) / float64(h2.vel.X)
	return t1 > 0 && t2 > 0
}

func intersection(h1, h2 hail) (x float64, y float64, parallel bool) {
	m1, C1 := makeLine(h1)
	m2, C2 := makeLine(h2)
	if m1 == m2 {
		parallel = true
	} else {
		x = (C2 - C1) / (m1 - m2)
		y = m1*x + C1
	}
	return
}

func makeLine(h hail) (float64, float64) {
	slope := float64(h.vel.Y) / float64(h.vel.X)
	intercept := float64(h.pos.Y) - slope*float64(h.pos.X)
	return slope, intercept
}

func getRockVelocity(hails []hail) util.Pt3 {
	fitRockVelocity := func(h *[]int, h0, h1 hail, perAxisPosVel func(h hail) (int, int)) {
		pos0, vel0 := perAxisPosVel(h0)
		pos1, vel1 := perAxisPosVel(h1)
		if vel0 == vel1 {
			candidates := match(pos1-pos0, vel0)
			if len(*h) == 0 {
				*h = candidates
			} else {
				*h = setIntersection(*h, candidates)
			}
		}
	}

	var x, y, z []int
loopi:
	for i := 0; i < len(hails)-1; i++ {
		for j := i + 1; j < len(hails); j++ {
			fitRockVelocity(&x, hails[i], hails[j], func(h hail) (int, int) { return h.pos.X, h.vel.X })
			fitRockVelocity(&y, hails[i], hails[j], func(h hail) (int, int) { return h.pos.Y, h.vel.Y })
			fitRockVelocity(&z, hails[i], hails[j], func(h hail) (int, int) { return h.pos.Z, h.vel.Z })
			if len(x) == 1 && len(y) == 1 && len(z) == 1 {
				break loopi
			}
		}
	}
	return util.Pt3{x[0], y[0], z[0]}
}

func match(delta, v int) []int {
	var res []int
	// Brute force through velocities [-1000, 1000]
	// https://old.reddit.com/r/adventofcode/comments/18pnycy/2023_day_24_solutions/keqf8uq/#:~:text=every%20velocity%20from-,%2D1000%20to%201000,-that%20satisfies%20this
	for vv := -1000; vv <= 1000; vv++ {
		if vv != v && delta%(vv-v) == 0 {
			res = append(res, vv)
		}
	}
	return res
}

func setIntersection(a, b []int) []int {
	var res []int
	for _, v := range a {
		if slices.Contains(b, v) {
			res = append(res, v)
		}
	}
	return res
}

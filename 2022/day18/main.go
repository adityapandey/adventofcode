package main

import (
	"fmt"
	"math"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	cubes := map[util.Pt3]struct{}{}
	neighbors := []util.Pt3{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}
	min := util.Pt3{math.MaxInt, math.MaxInt, math.MaxInt}
	max := util.Pt3{math.MinInt, math.MinInt, math.MinInt}
	s := util.ScanAll()
	for s.Scan() {
		var cube util.Pt3
		fmt.Sscanf(s.Text(), "%d,%d,%d", &cube.X, &cube.Y, &cube.Z)
		cubes[cube] = struct{}{}
		min = util.Pt3{util.Min(min.X, cube.X), util.Min(min.Y, cube.Y), util.Min(min.Z, cube.Z)}
		max = util.Pt3{util.Max(max.X, cube.X), util.Max(max.Y, cube.Y), util.Max(max.Z, cube.Z)}
	}
	min = min.Add(util.Pt3{-1, -1, -1})
	max = max.Add(util.Pt3{1, 1, 1})

	faces := 0
	for cube := range cubes {
		for _, delta := range neighbors {
			if _, ok := cubes[cube.Add(delta)]; !ok {
				faces++
			}
		}
	}
	fmt.Println(faces)

	faces = 0
	q := []util.Pt3{min}
	seen := map[util.Pt3]struct{}{min: {}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, delta := range neighbors {
			next := curr.Add(delta)
			if next.X < min.X ||
				next.Y < min.Y ||
				next.Z < min.Z ||
				next.X > max.X ||
				next.Y > max.Y ||
				next.Z > max.Z {
				continue
			}
			if _, ok := cubes[next]; ok {
				faces++
			} else if _, ok := seen[next]; !ok {
				seen[next] = struct{}{}
				q = append(q, next)
			}
		}
	}
	fmt.Println(faces)
}

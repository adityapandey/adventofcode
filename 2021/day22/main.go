package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

type cuboid struct {
	min util.Pt3 // inclusive
	max util.Pt3 // exclusive
	on  bool
}

func newCuboid(x0, x1, y0, y1, z0, z1 int, on bool) cuboid {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	if z0 > z1 {
		z0, z1 = z1, z0
	}
	return cuboid{util.Pt3{x0, y0, z0}, util.Pt3{x1 + 1, y1 + 1, z1 + 1}, on}
}

func (c cuboid) empty() bool {
	return c.min.X >= c.max.X || c.min.Y >= c.max.Y || c.min.Z >= c.max.Z

}

func (c cuboid) intersect(d cuboid) cuboid {
	if c.min.X < d.min.X {
		c.min.X = d.min.X
	}
	if c.min.Y < d.min.Y {
		c.min.Y = d.min.Y
	}
	if c.min.Z < d.min.Z {
		c.min.Z = d.min.Z
	}
	if c.max.X > d.max.X {
		c.max.X = d.max.X
	}
	if c.max.Y > d.max.Y {
		c.max.Y = d.max.Y
	}
	if c.max.Z > d.max.Z {
		c.max.Z = d.max.Z
	}
	if c.empty() {
		return cuboid{}
	}
	return c
}

func (c cuboid) volume() int {
	return (c.max.X - c.min.X) * (c.max.Y - c.min.Y) * (c.max.Z - c.min.Z)
}

func main() {
	var init, cuboids []cuboid
	s := util.ScanAll()
	for s.Scan() {
		var onOff string
		var x0, x1, y0, y1, z0, z1 int
		fmt.Sscanf(s.Text(), "%s x=%d..%d,y=%d..%d,z=%d..%d", &onOff, &x0, &x1, &y0, &y1, &z0, &z1)
		c := newCuboid(x0, x1, y0, y1, z0, z1, onOff == "on")
		cuboids = append(cuboids, c)
		if c.min.X >= -50 && c.max.X <= 51 &&
			c.min.Y >= -50 && c.max.Y <= 51 &&
			c.min.Z >= -50 && c.max.Z <= 51 {
			init = append(init, c)
		}
	}

	fmt.Println(lit(init))
	fmt.Println(lit(cuboids))
}

func lit(cuboids []cuboid) int {
	var sum int
	for i := range cuboids {
		if cuboids[i].on {
			sum += partialVolume(cuboids[i], cuboids[i+1:])
		}
	}
	return sum
}

func partialVolume(c cuboid, cuboids []cuboid) int {
	var doubleCounts []cuboid
	for _, d := range cuboids {
		overlap := c.intersect(d)
		if overlap.empty() {
			continue
		}
		doubleCounts = append(doubleCounts, overlap)
	}
	var sum int
	for i := range doubleCounts {
		sum += partialVolume(doubleCounts[i], doubleCounts[i+1:])
	}
	return c.volume() - sum
}

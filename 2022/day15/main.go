package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type sensor struct {
	pos    image.Point
	beacon image.Point
	dist   int
}

func main() {
	var sensors []sensor
	for _, line := range strings.Split(util.ReadAll(), "\n") {
		var s sensor
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.pos.X, &s.pos.Y, &s.beacon.X, &s.beacon.Y)
		s.dist = util.Manhattan(s.pos, s.beacon)
		sensors = append(sensors, s)
	}
	fmt.Println(impossible(sensors, 2000000))
	fmt.Println(distress(sensors, 4000000))
}

func impossible(sensors []sensor, y int) int {
	pts := util.SetOf([]int{})
	for _, s := range sensors {
		dist := s.dist - util.Abs(s.pos.Y-y)
		for x := 0; x <= dist; x++ {
			pts[s.pos.X+x] = struct{}{}
			pts[s.pos.X-x] = struct{}{}
		}
	}
	for _, s := range sensors {
		if s.beacon.Y == y {
			delete(pts, s.beacon.X)
		}
	}
	return len(pts)
}

func distress(sensors []sensor, maxcoord int) int {
	for x := 0; x <= maxcoord; x++ {
		for y := 0; y <= maxcoord; y++ {
			p := image.Pt(x, y)
			detected := false
			skip := 0
			for _, s := range sensors {
				if util.Manhattan(s.pos, p) <= s.dist {
					detected = true
					dist := s.dist - util.Abs(s.pos.X-x)
					skip = util.Max(skip, dist+s.pos.Y-y)
				}
			}
			if !detected {
				return x*4000000 + y
			}
			y += skip
		}
	}
	return -1
}

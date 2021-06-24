package main

import (
	"container/ring"
	"fmt"
	"image"
	"math"

	"github.com/adityapandey/adventofcode/util"
	"go4.org/sort"
)

type slope struct {
	x, y int
}

func slopeFrom(p image.Point) slope {
	g := gcd(p.X, p.Y)
	return slope{p.X / g, p.Y / g}
}

func (s slope) quadrant() int {
	// y-axis increases down
	switch {
	case s.x >= 0 && s.y <= 0:
		return 0
	case s.x >= 0 && s.y >= 0:
		return 1
	case s.x <= 0 && s.y >= 0:
		return 2
	default:
		return 3
	}
}

func (s slope) less(other slope) bool {
	if s == other {
		return false
	}
	if s.quadrant() == other.quadrant() {
		switch s.quadrant() {
		case 0, 2:
			if s.y == 0 {
				return false
			}
			return math.Abs(float64(s.x)/float64(s.y)) < math.Abs(float64(other.x)/float64(other.y))
		default:
			if s.x == 0 {
				return false
			}
			return math.Abs(float64(s.y)/float64(s.x)) < math.Abs(float64(other.y)/float64(other.x))
		}
	}
	return s.quadrant() < other.quadrant()
}

func main() {
	var asteroids []image.Point
	s := util.ScanAll()
	var y int
	for s.Scan() {
		line := s.Text()
		for x, b := range []byte(line) {
			if b == '#' {
				asteroids = append(asteroids, image.Pt(x, y))
			}
		}
		y++
	}

	var maxDetected int
	var station image.Point
	for _, a := range asteroids {
		m := make(map[slope]struct{})
		for i := range asteroids {
			if asteroids[i] == a {
				continue
			}
			m[slopeFrom(asteroids[i].Sub(a))] = struct{}{}
		}
		if len(m) > maxDetected {
			maxDetected, station = len(m), a
		}
	}
	fmt.Println(maxDetected)

	asteroidsBySlope := make(map[slope][]image.Point)
	for _, a := range asteroids {
		if a == station {
			continue
		}
		asteroidsBySlope[slopeFrom(a.Sub(station))] = append(asteroidsBySlope[slopeFrom(a.Sub(station))], a)
	}
	for s := range asteroidsBySlope {
		sort.Slice(asteroidsBySlope[s], func(i, j int) bool {
			return util.Manhattan(asteroidsBySlope[s][i], station) < util.Manhattan(asteroidsBySlope[s][j], station)
		})
	}
	var slopes []slope
	for s := range asteroidsBySlope {
		slopes = append(slopes, s)
	}
	sort.Slice(slopes, func(i, j int) bool { return slopes[i].less(slopes[j]) })
	r := ring.New(len(slopes))
	for i := 0; i < len(slopes); i++ {
		r.Value = slopes[i]
		r = r.Next()
	}
	for i := 0; i < 199; i++ {
		for len(asteroidsBySlope[r.Value.(slope)]) == 0 {
			r = r.Next()
		}
		asteroidsBySlope[r.Value.(slope)] = asteroidsBySlope[r.Value.(slope)][1:]
		r = r.Next()
	}
	vaporized := asteroidsBySlope[r.Value.(slope)][0]
	fmt.Println(vaporized.X*100 + vaporized.Y)
}

func gcd(a, b int) int {
	a, b = util.Abs(a), util.Abs(b)
	if a == 0 || b == 0 {
		return a + b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

type reindeer struct {
	name                  string
	speed, duration, rest int
	points                int
}

func (r reindeer) distance(t int) int {
	d := r.speed * r.duration * (t / (r.duration + r.rest))
	d += r.speed * min(r.duration, t%(r.duration+r.rest))
	return d
}

func main() {
	var reindeers []reindeer
	s := util.ScanAll()
	for s.Scan() {
		var r reindeer
		fmt.Sscanf(s.Text(), "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &r.name, &r.speed, &r.duration, &r.rest)
		reindeers = append(reindeers, r)
	}

	const endTime = 2503
	var leader int
	for _, r := range reindeers {
		leader = max(leader, r.distance(endTime))
	}
	fmt.Println(leader)

	for t := 1; t <= endTime; t++ {
		var leader int
		for i := range reindeers {
			leader = max(leader, reindeers[i].distance(t))
		}
		for i := range reindeers {
			if reindeers[i].distance(t) == leader {
				reindeers[i].points++
			}
		}
	}

	var maxPoints int
	for _, r := range reindeers {
		maxPoints = max(maxPoints, r.points)
	}
	fmt.Println(maxPoints)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type orientation struct {
	x, y, z int
	xyz     [3]int
}

func (o orientation) transform(p [3]int) [3]int {
	var r [3]int
	for i := 0; i < 3; i++ {
		r[i] = p[o.xyz[i]]
	}
	r[0] *= o.x
	r[1] *= o.y
	r[2] *= o.z
	return r
}

type scanner struct {
	pos     [3]int
	beacons [][3]int
}

func (s scanner) relativeTo(to, from [3]int, o orientation) scanner {
	from = o.transform(from)
	var r scanner
	for i := 0; i < 3; i++ {
		r.pos[i] = to[i] - from[i]
	}
	for _, b := range s.beacons {
		bb := o.transform(b)
		for i := 0; i < 3; i++ {
			bb[i] += to[i] - from[i]
		}
		r.beacons = append(r.beacons, bb)
	}
	return r
}

func (s scanner) common(other scanner) int {
	var sum int
	for _, i := range s.beacons {
		for _, j := range other.beacons {
			if i == j {
				sum++
			}
		}
	}
	return sum
}

func main() {
	scanners := map[int]scanner{}
	unresolved := map[int]struct{}{}
	for i, sp := range strings.Split(util.ReadAll(), "\n\n") {
		unresolved[i] = struct{}{}
		var s scanner
		lines := strings.Split(sp, "\n")
		for j := 1; j < len(lines); j++ {
			var b [3]int
			fmt.Sscanf(lines[j], "%d,%d,%d", &b[0], &b[1], &b[2])
			s.beacons = append(s.beacons, b)
		}
		scanners[i] = s
	}

	delete(unresolved, 0)
	resolved := []int{0}
	for len(unresolved) > 0 {
		r := resolved[0]
		resolved = resolved[1:]
		for u := range unresolved {
		outer:
			for _, xyz := range [][3]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}} {
				for _, x := range []int{-1, 1} {
					for _, y := range []int{-1, 1} {
						for _, z := range []int{-1, 1} {
							o := orientation{x, y, z, xyz}
							for _, unresolvedBeacon := range scanners[u].beacons {
								for _, resolvedBeacon := range scanners[r].beacons {
									relativePos := scanners[u].relativeTo(resolvedBeacon, unresolvedBeacon, o)
									if relativePos.common(scanners[r]) >= 12 {
										scanners[u] = relativePos
										delete(unresolved, u)
										resolved = append(resolved, u)
										break outer
									}
								}
							}
						}
					}
				}
			}
		}
	}

	beacons := map[[3]int]struct{}{}
	for _, s := range scanners {
		for _, b := range s.beacons {
			beacons[b] = struct{}{}
		}
	}
	fmt.Println(len(beacons))

	max := 0
	for i := 0; i < len(scanners)-1; i++ {
		for j := i + 1; j < len(scanners); j++ {
			max = util.Max(max, manhattan(scanners[i].pos, scanners[j].pos))
		}
	}
	fmt.Println(max)
}

func manhattan(p, q [3]int) int {
	var s int
	for i := 0; i < 3; i++ {
		s += util.Abs(p[i] - q[i])
	}
	return s
}

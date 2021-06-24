package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var pos, vel [4][3]int
	input := strings.Split(util.ReadAll(), "\n")
	for i := range input {
		fmt.Sscanf(input[i], "<x=%d, y=%d, z=%d>", &pos[i][0], &pos[i][1], &pos[i][2])
	}

	fmt.Println(energy(pos, vel))

	var repeats [3]int
	for dim := 0; dim < 3; dim++ {
		repeats[dim] = repeat(pos, vel, dim)
	}
	fmt.Println(util.Lcm(repeats[0], util.Lcm(repeats[1], repeats[2])))
}

func energy(pos, vel [4][3]int) int {
	for step := 0; step < 1000; step++ {
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				for k := 0; k < 3; k++ {
					if pos[i][k] > pos[j][k] {
						vel[i][k]--
						vel[j][k]++
					} else if pos[i][k] < pos[j][k] {
						vel[i][k]++
						vel[j][k]--
					}
				}
			}
		}
		for i := 0; i < 4; i++ {
			for k := 0; k < 3; k++ {
				pos[i][k] += vel[i][k]
			}
		}
	}
	var sum int
	for i := 0; i < 4; i++ {
		var pot, kin int
		for k := 0; k < 3; k++ {
			pot += util.Abs(pos[i][k])
			kin += util.Abs(vel[i][k])
		}
		sum += pot * kin
	}
	return sum
}

func repeat(pos, vel [4][3]int, dim int) int {
	var steps int
	var p0, v0 [4]int
	for i := 0; i < 4; i++ {
		p0[i], v0[i] = pos[i][dim], vel[i][dim]
	}
	for {
		steps++
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				if pos[i][dim] > pos[j][dim] {
					vel[i][dim]--
					vel[j][dim]++
				} else if pos[i][dim] < pos[j][dim] {
					vel[i][dim]++
					vel[j][dim]--
				}
			}
		}
		seen := true
		for i := 0; i < 4; i++ {
			pos[i][dim] += vel[i][dim]
			if pos[i][dim] != p0[i] || vel[i][dim] != v0[i] {
				seen = false
			}
		}
		if seen {
			break
		}
	}
	return steps
}

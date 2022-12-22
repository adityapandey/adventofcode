package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

const rockstr = `####

 # 
###
 # 

  #
  #
###

#
#
#
#

##
##`

func main() {
	jetPattern := []byte(util.ReadAll())

	rocks := getRocks()
	grid := map[image.Point]struct{}{}
	for x := 0; x < 7; x++ {
		grid[image.Pt(x, 0)] = struct{}{}
	}
	floor, j := 0, 0
	repeat := map[[2]int][2]int{}

	for i, curr := 0, 0; ; i, curr = i+1, (curr+1)%len(rocks) {
		if i == 2022 {
			fmt.Println(floor)
		}
		key := [2]int{curr, j}
		if r, ok := repeat[key]; ok {
			previ, prevFloor := r[0], r[1]
			if (1000000000000-i)%(i-previ) == 0 {
				fmt.Println(floor + (1000000000000-i)/(i-previ)*(floor-prevFloor))
				break
			}
		}
		repeat[key] = [2]int{i, floor}
		currRock := rocks[curr]
		pos := image.Pt(2, floor+4)
		for {
			jet := jetPattern[j]
			j = (j + 1) % len(jetPattern)
			pos = pos.Add(util.DirFromByte(jet).Point())
			if collision(grid, currRock, pos) {
				pos = pos.Sub(util.DirFromByte(jet).Point())
			}
			pos = pos.Add(util.S.Point())
			if collision(grid, currRock, pos) {
				pos = pos.Sub(util.S.Point())
				for p := range currRock {
					grid[p.Add(pos)] = struct{}{}
					if p.Add(pos).Y > floor {
						floor = p.Add(pos).Y
					}
				}
				break
			}
		}
	}
}

func collision(grid, rock map[image.Point]struct{}, pos image.Point) bool {
	for p := range rock {
		_, ok := grid[p.Add(pos)]
		if ok || p.Add(pos).X < 0 || p.Add(pos).X > 6 {
			return true
		}
	}
	return false
}

func getRocks() []map[image.Point]struct{} {
	rocks := []map[image.Point]struct{}{}
	for i, rock := range strings.Split(rockstr, "\n\n") {
		rocks = append(rocks, map[image.Point]struct{}{})
		lines := strings.Split(rock, "\n")
		for y, line := range lines {
			for x := 0; x < len(line); x++ {
				if line[x] == '#' {
					rocks[i][image.Pt(x, len(lines)-1-y)] = struct{}{}
				}
			}
		}
	}
	return rocks
}

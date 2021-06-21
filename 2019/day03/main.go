package main

import (
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	s := util.ScanAll()
	s.Scan()
	w1 := getPath(s.Text())
	s.Scan()
	w2 := getPath(s.Text())

	intersections := make(map[image.Point]struct{})
	for p := range w1 {
		if _, ok := w2[p]; ok {
			intersections[p] = struct{}{}
		}
	}

	var origin image.Point
	minDist, minDelay := math.MaxInt32, math.MaxInt32
	for x := range intersections {
		minDist = util.Min(minDist, util.Manhattan(x, origin))
		minDelay = util.Min(minDelay, w1[x][0]+w2[x][0])
	}
	fmt.Println(minDist)
	fmt.Println(minDelay)
}

func getPath(w string) map[image.Point][]int {
	steps := strings.Split(w, ",")
	path := make(map[image.Point][]int)
	var p image.Point
	var i int
	for _, s := range steps {
		var b byte
		var n int
		fmt.Sscanf(s, "%c%d", &b, &n)
		for j := 0; j < n; j++ {
			i++
			p = p.Add(util.DirFromByte(b).Point())
			path[p] = append(path[p], i)
		}
	}
	return path
}

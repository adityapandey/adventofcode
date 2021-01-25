package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var lr = map[byte]int{'R': 1, 'L': 3}

func main() {
	var p, hq image.Point
	var d util.Dir
	seen := make(map[image.Point]struct{})
	var hqSeen bool
	for _, step := range strings.Split(util.ReadAll(), ", ") {
		var b byte
		var n int
		fmt.Sscanf(step, "%c%d", &b, &n)
		d = d.Add(lr[b])
		for i := 0; i < n; i++ {
			p = p.Add(d.Point())
			if _, ok := seen[p]; ok && !hqSeen {
				hq, hqSeen = p, true
			}
			seen[p] = struct{}{}
		}
	}
	fmt.Println(util.Manhattan(p, image.Pt(0, 0)))
	fmt.Println(util.Manhattan(hq, image.Pt(0, 0)))
}

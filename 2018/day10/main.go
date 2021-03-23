package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var ps, vs []image.Point
	s := util.ScanAll()
	for s.Scan() {
		var p, v image.Point
		fmt.Sscanf(s.Text(),
			"position=<%d,%d> velocity=<%d,%d>",
			&p.X, &p.Y, &v.X, &v.Y)
		ps = append(ps, p)
		vs = append(vs, v)
	}

	var lastSpread, spread, t int
	for {
		spread = util.Bounds(ps).Dx()
		if t > 0 && spread > lastSpread {
			break
		}
		for i := range ps {
			ps[i] = ps[i].Add(vs[i])
		}
		lastSpread = spread
		t++
	}

	for i := range ps {
		ps[i] = ps[i].Sub(vs[i])
	}
	r := util.Bounds(ps)
	screen := make([]byte, (r.Dx()+1)*(r.Dy()+1))
	for i := range screen {
		screen[i] = ' '
	}
	for _, p := range ps {
		screen[(p.Y-r.Min.Y)*(r.Dx()+1)+(p.X-r.Min.X)] = '#'
	}
	for y := 0; y <= r.Dy(); y++ {
		for x := 0; x <= r.Dx(); x++ {
			fmt.Printf("%c", screen[y*(r.Dx()+1)+x])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println(t - 1)
}

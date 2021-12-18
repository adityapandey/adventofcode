package main

import (
	"fmt"
	"image"
	"os"
)

func main() {
	var x0, x1, y0, y1 int
	fmt.Fscanf(os.Stdin, "target area: x=%d..%d, y=%d..%d", &x0, &x1, &y0, &y1)
	target := image.Rect(x0, y0, x1, y1)
	fmt.Println(-target.Min.Y * (-target.Min.Y - 1) / 2)
	var trajectories []image.Point
	for x := 0; x <= target.Max.X; x++ {
		for y := target.Min.Y; y < -target.Min.Y; y++ {
			trajectory := image.Pt(x, y)
			if hits(trajectory, target) {
				trajectories = append(trajectories, trajectory)
			}
		}
	}
	fmt.Println(len(trajectories))
}

func hits(v image.Point, target image.Rectangle) bool {
	target.Max.X++
	target.Max.Y++
	p := image.Pt(0, 0)
	for {
		next := p.Add(v)
		if next.In(target) {
			return true
		}
		if v.X > 0 {
			v.X--
		}
		v.Y--
		if (v.X == 0 && next.X < target.Min.X) ||
			next.X > target.Max.X ||
			(v.Y <= 0 && next.Y < target.Min.Y) {
			return false
		}
		p = next
	}
}

package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	grid := map[image.Point]struct{}{}
	s := util.ScanAll()
	for s.Scan() {
		sp := strings.Split(s.Text(), " -> ")
		pts := make([]image.Point, len(sp))
		for i := range sp {
			fmt.Sscanf(sp[i], "%d,%d", &pts[i].X, &pts[i].Y)
		}
		for i := 0; i < len(pts)-1; i++ {
			if pts[i].X == pts[i+1].X {
				for y := util.Min(pts[i].Y, pts[i+1].Y); y <= util.Max(pts[i].Y, pts[i+1].Y); y++ {
					grid[image.Pt(pts[i].X, y)] = struct{}{}
				}
			} else {
				for x := util.Min(pts[i].X, pts[i+1].X); x <= util.Max(pts[i].X, pts[i+1].X); x++ {
					grid[image.Pt(x, pts[i].Y)] = struct{}{}
				}
			}
		}
	}

	fmt.Println(fill(grid))
}

func fill(grid map[image.Point]struct{}) (int, int) {
	floor := bounds(grid).Max.Y + 1
	sands, firstFloorTouch := 0, 0
	for full := false; !full; _, full = grid[image.Pt(500, 0)] {
		for sand, settled := image.Pt(500, 0), false; !settled; sand, settled = next(grid, sand) {
			if sand.Y == floor {
				if firstFloorTouch == 0 {
					firstFloorTouch = sands
				}
				grid[sand] = struct{}{}
				break
			}
		}
		sands++
	}
	return firstFloorTouch, sands
}

var D = util.DirFromByte('D').PointR()
var L = util.DirFromByte('L').PointR()
var R = util.DirFromByte('R').PointR()

func next(grid map[image.Point]struct{}, sand image.Point) (image.Point, bool) {
	for _, n := range []image.Point{sand.Add(D), sand.Add(D).Add(L), sand.Add(D).Add(R)} {
		if _, ok := grid[n]; !ok {
			return n, false
		}
	}
	grid[sand] = struct{}{}
	return sand, true
}

func bounds(grid map[image.Point]struct{}) image.Rectangle {
	pts := make([]image.Point, 0, len(grid))
	for p := range grid {
		pts = append(pts, p)
	}
	return util.Bounds(pts)
}

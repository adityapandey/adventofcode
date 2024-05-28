package main

import (
	_ "embed"
	"fmt"
	"image"
	"slices"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

func main() {
	grid := map[image.Point]byte{}
	for y, line := range strings.Split(util.ReadAll(), "\n") {
		for x := 0; x < len(line); x++ {
			grid[image.Pt(x, y)] = line[x]
		}
	}
	bounds := util.Bounds(maps.Keys(grid))
	start := image.Pt(1, 0)
	end := image.Pt(bounds.Max.X-2, bounds.Max.Y-1)

	fmt.Println(dfs(start, end, grid, 0, map[image.Point]struct{}{}))

	intersections := getIntersections(start, end, grid)
	weightsByIntersection := getWeights(intersections, grid)

	fmt.Println(dfs2(start, end, weightsByIntersection, 0, map[image.Point]struct{}{}))
}

func dfs(start, end image.Point, grid map[image.Point]byte, startDist int, visited map[image.Point]struct{}) int {
	if start == end {
		return startDist
	}
	longest := 0
	for _, next := range nexts(start, grid, false) {
		if _, ok := visited[next]; ok {
			continue
		}
		visited[next] = struct{}{}
		longest = util.Max(longest, dfs(next, end, grid, startDist+1, visited))
		delete(visited, next)
	}
	return longest
}

func nexts(p image.Point, grid map[image.Point]byte, dry bool) []image.Point {
	var ns []image.Point
	if dry || grid[p] == '.' {
		for _, n := range util.Neighbors4 {
			ns = append(ns, p.Add(n))
		}
	} else {
		ns = []image.Point{p.Add(util.DirFromByte(grid[p]).PointR())}
	}
	for i := 0; i < len(ns); i++ {
		if b, ok := grid[ns[i]]; !ok || b == '#' {
			ns = append(ns[:i], ns[i+1:]...)
			i--
		}
	}
	return ns
}

func getIntersections(start, end image.Point, grid map[image.Point]byte) []image.Point {
	var intersections []image.Point
	for p := range grid {
		if grid[p] == '#' {
			continue
		}
		if len(nexts(p, grid, true)) > 2 {
			intersections = append(intersections, p)
		}
	}
	intersections = append(intersections, []image.Point{start, end}...)
	return intersections
}

type branchWeight struct {
	branch image.Point
	weight int
}

func getWeights(intersections []image.Point, grid map[image.Point]byte) map[image.Point][]branchWeight {
	weightsByIntersection := map[image.Point][]branchWeight{}
	for _, intersection := range intersections {
		frontier := []image.Point{intersection}
		weight := 0
		visited := map[image.Point]struct{}{intersection: {}}
		for len(frontier) > 0 {
			weight++
			var nextFrontier []image.Point
			for _, p := range frontier {
				for _, next := range nexts(p, grid, true) {
					if _, ok := visited[next]; !ok {
						visited[next] = struct{}{}
						if slices.Contains(intersections, next) {
							weightsByIntersection[intersection] = append(weightsByIntersection[intersection], branchWeight{next, weight})
						} else {
							nextFrontier = append(nextFrontier, next)
						}
					}
				}
			}
			frontier = nextFrontier
		}
	}
	return weightsByIntersection
}

func dfs2(start, end image.Point, weightsByIntersection map[image.Point][]branchWeight, startWeight int, visited map[image.Point]struct{}) int {
	if start == end {
		return startWeight
	}
	longest := 0
	for _, bw := range weightsByIntersection[start] {
		if _, ok := visited[bw.branch]; ok {
			continue
		}
		visited[bw.branch] = struct{}{}
		longest = util.Max(longest, dfs2(bw.branch, end, weightsByIntersection, startWeight+bw.weight, visited))
		delete(visited, bw.branch)
	}
	return longest
}

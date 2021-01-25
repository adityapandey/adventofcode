package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type Point struct {
	x, y int
}

type Closest struct {
	distance      int
	point         *Point
	totalDistance int
}

func main() {
	var points []Point
	var minX, minY, maxX, maxY int
	minX, minY = math.MaxInt32, math.MaxInt32
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var p Point
		fmt.Sscanf(s.Text(), "%d, %d", &p.x, &p.y)
		points = append(points, p)
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	width := maxX - minX + 1
	height := maxY - minY + 1

	grid := make([][]Closest, width)
	for i := range grid {
		grid[i] = make([]Closest, height)
		for j := range grid[i] {
			grid[i][j].distance = width + height
		}
	}

	// Part 1
	for i, p := range points {
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				distance := abs(x-(p.x-minX)) + abs(y-(p.y-minY))
				if distance == grid[x][y].distance {
					grid[x][y].point = nil
				}
				if distance < grid[x][y].distance {
					grid[x][y].point = &points[i]
					grid[x][y].distance = distance
				}
			}
		}
	}

	m := make(map[*Point]int)
	border := make(map[*Point]struct{})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[x][y].point != nil {
				m[grid[x][y].point]++
				if x == 0 || x == width-1 || y == 0 || y == height-1 {
					border[grid[x][y].point] = struct{}{}
				}
			}
		}
	}

	var maxArea int
	for p, area := range m {
		_, ok := border[p]
		if !ok && area > maxArea {
			maxArea = area
		}
	}

	fmt.Println(maxArea)

	// Part 2
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for _, p := range points {
				grid[x][y].totalDistance += abs(x-(p.x-minX)) + abs(y-(p.y-minY))
			}
		}
	}

	var safeArea int
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[x][y].totalDistance < 10000 {
				safeArea++
			}
		}
	}

	fmt.Println(safeArea)
}

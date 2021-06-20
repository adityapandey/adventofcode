package main

import (
	"container/heap"
	"fmt"
	"image"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

type regionType int

const (
	rocky regionType = iota
	wet
	narrow
)

type tool int

const (
	neither tool = iota
	torch
	climber
)

type region struct {
	p    image.Point
	tool tool
}

type cave struct {
	target     image.Point
	depth      int
	erosionMap map[image.Point]int
}

func (c *cave) getErosionLevel(p image.Point) int {
	if e, ok := c.erosionMap[p]; ok {
		return e
	}
	var g int
	if p == image.Pt(0, 0) || p == c.target {
		g = 0
	} else if p.X == 0 {
		g = p.Y * 48271
	} else if p.Y == 0 {
		g = p.X * 16807
	} else {
		g = c.getErosionLevel(p.Add(image.Pt(-1, 0))) * c.getErosionLevel(p.Add(image.Pt(0, -1)))
	}
	e := (g + c.depth) % 20183
	c.erosionMap[p] = e
	return e
}

func (c *cave) getRegionType(p image.Point) regionType {
	return regionType(c.getErosionLevel(p) % 3)
}

func (c *cave) getNeighbors(r region) []region {
	var regions []region
	for _, n := range util.Neighbors4 {
		p := r.p.Add(n)
		if p.X < 0 || p.Y < 0 || int(c.getRegionType(p)) == int(r.tool) {
			continue
		}
		regions = append(regions, region{r.p.Add(n), r.tool})
	}
	return regions
}

func main() {
	c := cave{erosionMap: make(map[image.Point]int)}
	fmt.Fscanf(os.Stdin, "depth: %d\n", &c.depth)
	fmt.Fscanf(os.Stdin, "target: %d,%d", &c.target.X, &c.target.Y)

	// Part 1
	var riskLevel int
	for x := 0; x <= c.target.X; x++ {
		for y := 0; y <= c.target.Y; y++ {
			riskLevel += int(c.getRegionType(image.Pt(x, y)))
		}
	}
	fmt.Println(riskLevel)

	// Part 2
	mouth := region{image.Pt(0, 0), torch}
	dist := map[region]int{mouth: 0}
	pq := util.PQ{&util.Item{mouth, 0}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(region)
		currDist := dist[curr]
		if curr == (region{c.target, torch}) {
			break
		}
		for _, n := range c.getNeighbors(curr) {
			if d, ok := dist[n]; !ok || currDist+1 < d {
				d = currDist + 1
				heap.Push(&pq, &util.Item{n, -d})
				dist[n] = d
			}
		}
		changeTool := curr
		changeTool.tool = tool((3 - int(c.getRegionType(changeTool.p))) - int(changeTool.tool))
		if d, ok := dist[changeTool]; !ok || currDist+7 < d {
			d = currDist + 7
			heap.Push(&pq, &util.Item{changeTool, -d})
			dist[changeTool] = d
		}
	}
	fmt.Println(dist[region{c.target, torch}])
}

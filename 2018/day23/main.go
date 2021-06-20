package main

import (
	"container/heap"
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

var origin util.Pt3

type nanobot struct {
	p util.Pt3
	r int
}

func (n nanobot) inRangeOf(other nanobot) bool {
	return util.Manhattan3(n.p, other.p) <= n.r
}

func (n nanobot) touches(b box) bool {
	p0 := b.p0
	p1 := b.p1.Add(util.Pt3{-1, -1, -1}) // exclude end point, it belongs to the next adjacent box
	return util.Manhattan3(p0, n.p)+util.Manhattan3(p1, n.p)-(util.Manhattan3(p1, origin)-util.Manhattan3(p0, origin)) <= 2*n.r
}

type box struct {
	p0 util.Pt3
	p1 util.Pt3
}

func (b box) size() int {
	return b.p1.X - b.p0.X
}

func (b box) distance() int {
	return util.Manhattan3(b.p0, origin)
}

type item struct {
	b     box
	nBots int
}

type pq []*item

func (q pq) Len() int { return len(q) }
func (q pq) Less(i, j int) bool {
	if q[i].nBots == q[j].nBots {
		if q[i].b.size() == q[j].b.size() {
			return q[i].b.distance() < q[j].b.distance()
		}
		return q[i].b.size() > q[j].b.size()
	}
	return q[i].nBots > q[j].nBots
}
func (q pq) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *pq) Push(x interface{}) { *q = append(*q, x.(*item)) }
func (q *pq) Pop() interface{} {
	old := *q
	n := len(*q)
	i := old[n-1]
	*q = old[:n-1]
	return i
}

func main() {
	var nanobots []nanobot
	s := util.ScanAll()
	for s.Scan() {
		var n nanobot
		fmt.Sscanf(s.Text(), "pos=<%d,%d,%d>, r=%d\n", &n.p.X, &n.p.Y, &n.p.Z, &n.r)
		nanobots = append(nanobots, n)
	}

	var maxR, maxI int
	for i, n := range nanobots {
		if n.r > maxR {
			maxR, maxI = n.r, i
		}
	}

	var inRangeSum int
	for _, n := range nanobots {
		if nanobots[maxI].inRangeOf(n) {
			inRangeSum++
		}
	}

	fmt.Println(inRangeSum)

	var maxBoxDim int
	for _, n := range nanobots {
		maxBoxDim = util.Max(maxBoxDim, util.Abs(n.p.X)+n.r)
		maxBoxDim = util.Max(maxBoxDim, util.Abs(n.p.Y)+n.r)
		maxBoxDim = util.Max(maxBoxDim, util.Abs(n.p.Z)+n.r)
	}

	boxSize := 1
	for boxSize < maxBoxDim {
		boxSize *= 2
	}

	startBox := box{util.Pt3{-boxSize, -boxSize, -boxSize}, util.Pt3{boxSize, boxSize, boxSize}}
	q := pq{&item{b: startBox, nBots: len(nanobots)}}
	heap.Init(&q)
	for q.Len() > 0 {
		curr := heap.Pop(&q).(*item)
		if curr.b.size() == 1 {
			fmt.Println(curr.b.distance())
			break
		}
		nextSize := curr.b.size() / 2
		for _, octant := range [8][3]int{
			{0, 0, 0},
			{0, 0, 1},
			{0, 1, 0},
			{0, 1, 1},
			{1, 0, 0},
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 1},
		} {
			var neighbor box
			neighbor.p0 = curr.b.p0.Add(util.Pt3{octant[0], octant[1], octant[2]}.Mul(nextSize))
			neighbor.p1 = neighbor.p0.Add(util.Pt3{nextSize, nextSize, nextSize})
			var nBots int
			for _, n := range nanobots {
				if n.touches(neighbor) {
					nBots++
				}
			}
			heap.Push(&q, &item{b: neighbor, nBots: nBots})
		}
	}
}

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type P struct{ x, y, z int }

type Bot struct {
	P
	r int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func distance(p0, p1 P) int {
	return abs(p0.x-p1.x) + abs(p0.y-p1.y) + abs(p0.z-p1.z)
}

func (b Bot) InRange(other Bot) bool {
	return distance(b.P, other.P) <= b.r
}

func (b Bot) Touches(box Box) bool {
	p0 := box.p0
	p1 := P{box.p1.x - 1, box.p1.y - 1, box.p1.z - 1} // exclude end point, it belongs to the next adjacent box
	return distance(p0, b.P)+distance(p1, b.P)-(distance(p1, P{0, 0, 0})-distance(p0, P{0, 0, 0})) <= 2*b.r
}

type Box struct {
	p0 P
	p1 P
}

func (box Box) Size() int {
	return box.p1.x - box.p0.x
}

func (box Box) Distance() int {
	return distance(box.p0, P{0, 0, 0})
}

type Item struct {
	box           Box
	intersections int
	index         int
}

type Q []*Item

func (q Q) Len() int { return len(q) }
func (q Q) Less(i, j int) bool {
	if q[i].intersections == q[j].intersections {
		if q[i].box.Size() == q[j].box.Size() {
			return q[i].box.Distance() < q[j].box.Distance()
		}
		return q[i].box.Size() > q[j].box.Size()
	}
	return q[i].intersections > q[j].intersections
}
func (q Q) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}
func (q *Q) Push(x interface{}) {
	l := len(*q)
	item := x.(*Item)
	item.index = l
	*q = append(*q, item)
}
func (q *Q) Pop() interface{} {
	old := *q
	l := len(*q)
	x := old[l-1]
	x.index = -1
	*q = old[:l-1]
	return x
}

func main() {
	var bots []Bot
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var bot Bot
		fmt.Sscanf(s.Text(), "pos=<%d,%d,%d>, r=%d\n", &bot.x, &bot.y, &bot.z, &bot.r)
		bots = append(bots, bot)
	}

	var maxR, maxI int
	for i, bot := range bots {
		if bot.r > maxR {
			maxR, maxI = bot.r, i
		}
	}

	var inRangeSum int
	for _, bot := range bots {
		if bots[maxI].InRange(bot) {
			inRangeSum++
		}
	}

	fmt.Println(inRangeSum)

	var maxBoxDim int
	for _, b := range bots {
		maxBoxDim = max(maxBoxDim, abs(b.x)+b.r)
		maxBoxDim = max(maxBoxDim, abs(b.y)+b.r)
		maxBoxDim = max(maxBoxDim, abs(b.z)+b.r)
	}

	boxSize := 1
	for boxSize < maxBoxDim {
		boxSize *= 2
	}

	startBox := Box{P{-boxSize, -boxSize, -boxSize}, P{boxSize, boxSize, boxSize}}
	var q Q
	heap.Init(&q)
	heap.Push(&q, &Item{box: startBox, intersections: len(bots)})

	for q.Len() > 0 {
		item := heap.Pop(&q).(*Item)
		if item.box.Size() == 1 {
			fmt.Println(item.box.Distance())
			break
		}
		nextSize := item.box.Size() / 2
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
			var newBox Box
			newBox.p0.x = item.box.p0.x + octant[0]*nextSize
			newBox.p1.x = newBox.p0.x + nextSize
			newBox.p0.y = item.box.p0.y + octant[1]*nextSize
			newBox.p1.y = newBox.p0.y + nextSize
			newBox.p0.z = item.box.p0.z + octant[2]*nextSize
			newBox.p1.z = newBox.p0.z + nextSize
			intersections := 0
			for _, b := range bots {
				if b.Touches(newBox) {
					intersections++
				}
			}
			heap.Push(&q, &Item{box: newBox, intersections: intersections})
		}
	}
}

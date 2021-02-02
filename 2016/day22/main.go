package main

import (
	"container/heap"
	"fmt"
	"image"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`-x(\d+)-y(\d+)`)

type node struct {
	used, avail int
}

func main() {
	nodes := make(map[image.Point]node)
	input := strings.Split(util.ReadAll(), "\n")[2:]
	for i := range input {
		f := strings.Fields(input[i])
		matches := re.FindAllStringSubmatch(f[0], -1)[0]
		var n node
		p := image.Pt(util.Atoi(matches[1]), util.Atoi(matches[2]))
		n.used = util.Atoi(f[2][:len(f[2])-1])
		n.avail = util.Atoi(f[3][:len(f[3])-1])
		nodes[p] = n
	}
	fmt.Println(viable(nodes))
	fmt.Println(minmoves(nodes))
}

func viable(nodes map[image.Point]node) int {
	nodesByAvail := make([]node, 0, len(nodes))
	for _, n := range nodes {
		nodesByAvail = append(nodesByAvail, n)
	}
	sort.Slice(nodesByAvail, func(i, j int) bool { return nodesByAvail[i].avail < nodesByAvail[j].avail })
	var viable int
	for _, n := range nodes {
		if n.used == 0 {
			continue
		}
		i := sort.Search(len(nodesByAvail), func(i int) bool { return nodesByAvail[i].avail >= n.used })
		if i != len(nodesByAvail) {
			viable += (len(nodesByAvail) - i)
			if n.avail >= n.used {
				viable-- // double-counted self
			}
		}
	}
	return viable
}

func minmoves(nodes map[image.Point]node) int {
	w, _ := dim(nodes)
	goal := image.Pt(w, 0)
	hole, err := findHole(nodes)
	if err != nil {
		log.Fatal(err)
	}
	// move hole to adjacent to goal
	// swap goal and hole
	// repeat until goal at origin
	var sum int
	for goal != image.Pt(0, 0) {
		next := goal.Add(image.Pt(-1, 0))
		m, err := moves(nodes, goal, hole, next)
		if err != nil {
			log.Fatal(err)
		}
		sum += m
		hole = next
		m, err = moves(nodes, goal, goal, hole)
		if err != nil {
			log.Fatal(err)
		}
		sum += m
		goal, hole = hole, goal
	}
	return sum
}

func findHole(nodes map[image.Point]node) (image.Point, error) {
	for p, n := range nodes {
		if n.used == 0 {
			return p, nil
		}
	}
	return image.Pt(0, 0), fmt.Errorf("no hole")
}

func moves(nodes map[image.Point]node, goal, from, to image.Point) (int, error) {
	w, h := dim(nodes)
	depth := map[image.Point]int{from: 0}
	pq := util.PQ{&util.Item{from, 0}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*util.Item).Obj.(image.Point)
		if p == to {
			return depth[p], nil
		}
		currdepth := depth[p] + 1
		for _, n := range util.Neighbors4 {
			next := p.Add(n)
			if next.X < 0 || next.Y < 0 || next.X > w || next.Y > h || nodes[next].used > 400 || next == goal {
				continue
			}
			if d, ok := depth[next]; !ok || currdepth < d {
				depth[next] = currdepth
				heap.Push(&pq, &util.Item{next, -currdepth})
			}
		}
	}
	return -1, fmt.Errorf("no possible path")
}

func dim(m map[image.Point]node) (w int, h int) {
	for p := range m {
		if p.X > w {
			w = p.X
		}
		if p.Y > h {
			h = p.Y
		}
	}
	return
}

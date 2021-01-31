package main

import (
	"container/heap"
	"crypto/md5"
	"fmt"
	"image"
	"log"

	"github.com/adityapandey/adventofcode/util"
)

type node struct {
	p    image.Point
	path string
}

func main() {
	passcode := util.ReadAll()
	path, err := shortestPath(passcode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
	fmt.Println(longestPath(passcode, node{image.Pt(0, 0), ""}))
}

func shortestPath(passcode string) (string, error) {
	var p image.Point
	depth := map[node]int{{p, ""}: 0}
	pq := util.PQ{&util.Item{node{p, ""}, 0}}
	heap.Init(&pq)
	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*util.Item).Obj.(node)
		if n.p == image.Pt(3, 3) {
			return string(n.path), nil
		}
		for _, c := range openings(fmt.Sprintf("%v%v", passcode, string(n.path))) {
			next := n.p.Add(util.DirFromByte(c).PointR())
			if next.X < 0 || next.Y < 0 || next.X > 3 || next.Y > 3 {
				continue
			}
			nextnode := node{next, fmt.Sprintf("%v%c", n.path, c)}
			currdepth := depth[n] + 1
			if d, ok := depth[nextnode]; !ok || currdepth < d {
				depth[nextnode] = currdepth
				heap.Push(&pq, &util.Item{nextnode, -currdepth})
			}
		}
	}
	return "", fmt.Errorf("no path")
}

func longestPath(passcode string, n node) int {
	if n.p == image.Pt(3, 3) {
		return 0
	}
	var max int
	for _, c := range openings(fmt.Sprintf("%v%v", passcode, string(n.path))) {
		next := n.p.Add(util.DirFromByte(c).PointR())
		if next.X < 0 || next.Y < 0 || next.X > 3 || next.Y > 3 {
			continue
		}
		depth := 1 + longestPath(passcode, node{next, fmt.Sprintf("%v%c", n.path, c)})
		if depth > max {
			max = depth
		}
	}
	return max
}

func openings(s string) []byte {
	m := []byte("UDLR")
	var cs []byte
	for i, c := range hash(s)[:4] {
		if c >= 'b' {
			cs = append(cs, m[i])
		}
	}
	return cs
}

func hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

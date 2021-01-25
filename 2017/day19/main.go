package main

import (
	"bufio"
	"fmt"
	"image"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	path := make(map[image.Point]byte)
	s := bufio.NewScanner(os.Stdin)
	var y int
	var pos image.Point
	for s.Scan() {
		line := s.Text()
		for x := range line {
			if line[x] != ' ' {
				if y == 0 {
					pos = image.Pt(x, y)
				}
				path[image.Pt(x, y)] = line[x]
			}
		}
		y++
	}
	var letters []byte
	d := util.S
	var steps int
	for {
		steps++
		var ok bool
		if d, ok = next(pos, d, path); !ok {
			break
		}
		pos = pos.Add(d.PointR())
		if path[pos] >= 'A' && path[pos] <= 'Z' {
			letters = append(letters, path[pos])
		}
	}
	fmt.Println(string(letters))
	fmt.Println(steps)
}

func next(p image.Point, prevd util.Dir, path map[image.Point]byte) (util.Dir, bool) {
	for _, d := range []util.Dir{prevd, prevd.Next(), prevd.Prev()} {
		if _, ok := path[p.Add(d.PointR())]; ok {
			return d, true
		}
	}
	return 0, false
}

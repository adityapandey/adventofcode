package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

type dir int

const (
	S dir = iota
	W
	N
	E
)

var dirs = map[dir]image.Point{
	S: {0, 1},
	W: {-1, 0},
	N: {0, -1},
	E: {1, 0},
}

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
	var d dir
	var steps int
	for {
		steps++
		var ok bool
		if d, ok = next(pos, d, path); !ok {
			break
		}
		pos = pos.Add(dirs[d])
		if path[pos] >= 'A' && path[pos] <= 'Z' {
			letters = append(letters, path[pos])
		}
	}
	fmt.Println(string(letters))
	fmt.Println(steps)
}

func next(p image.Point, prevd dir, path map[image.Point]byte) (dir, bool) {
	for _, d := range []dir{prevd, (prevd + 1) % 4, (prevd + 3) % 4} {
		if _, ok := path[p.Add(dirs[d])]; ok {
			return d, true
		}
	}
	return 0, false
}

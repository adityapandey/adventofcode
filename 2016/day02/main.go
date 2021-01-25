package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var dirmap = map[byte]byte{
	'U': 'N',
	'D': 'S',
	'R': 'E',
	'L': 'W',
}

var pad1 = `123
456
789`

var pad2 = `  1
 234
56789
 ABC
  D`

func main() {
	var rows []string
	s := util.ScanAll()
	for s.Scan() {
		rows = append(rows, s.Text())
	}

	var code1, code2 []byte
	m1 := padmap(pad1)
	m2 := padmap(pad2)
	p1 := image.Pt(1, 1)
	p2 := image.Pt(0, 2)
	for _, row := range rows {
		for i := range row {
			p := util.DirFromByte(dirmap[row[i]]).PointR()
			next1, next2 := p1.Add(p), p2.Add(p)
			if _, ok := m1[next1]; ok {
				p1 = next1
			}
			if _, ok := m2[next2]; ok {
				p2 = next2
			}
		}
		code1 = append(code1, m1[p1])
		code2 = append(code2, m2[p2])
	}
	fmt.Println(string(code1))
	fmt.Println(string(code2))
}

func padmap(pad string) map[image.Point]byte {
	m := make(map[image.Point]byte)
	var y int
	for _, row := range strings.Split(pad, "\n") {
		for x := range row {
			if row[x] != ' ' {
				m[image.Pt(x, y)] = row[x]
			}
		}
		y++
	}
	return m
}

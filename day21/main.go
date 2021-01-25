package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(.+) => (.+)`)

const seed = `.#.
..#
###`

type grid struct {
	px   map[image.Point]byte
	size int
}

func (g grid) flip() grid {
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size/2; x++ {
			g.px[image.Pt(g.size-x-1, y)], g.px[image.Pt(x, y)] = g.px[image.Pt(x, y)], g.px[image.Pt(g.size-x-1, y)]
		}
	}
	return g
}

func (g grid) rotate() grid {
	for i := 0; i < g.size/2; i++ {
		for j := i; j < g.size-i-1; i++ {
			tmp := g.px[image.Pt(i, j)]
			g.px[image.Pt(i, j)] = g.px[image.Pt(g.size-1-j, i)]
			g.px[image.Pt(g.size-1-j, i)] = g.px[image.Pt(g.size-1-i, g.size-1-j)]
			g.px[image.Pt(g.size-1-i, g.size-1-j)] = g.px[image.Pt(j, g.size-1-i)]
			g.px[image.Pt(j, g.size-1-i)] = tmp
		}
	}
	return g
}

func (g grid) subgrid(p image.Point, size int) grid {
	sub := grid{make(map[image.Point]byte), size}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sub.px[image.Pt(i, j)] = g.px[p.Add(image.Pt(i, j))]
		}
	}
	return sub
}

func (g grid) merge(origin image.Point, s grid) {
	for p := range s.px {
		g.px[p.Add(origin)] = s.px[p]
	}
}

func (g grid) String() string {
	var sb strings.Builder
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			fmt.Fprintf(&sb, "%c", g.px[image.Pt(x, y)])
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}

func main() {
	transforms := make(map[string]string)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		m := re.FindAllStringSubmatch(s.Text(), -1)[0]
		transforms[m[1]] = m[2]
	}

	g := grid{make(map[image.Point]byte), 3}
	c := 0
	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			g.px[image.Pt(x, y)] = seed[c]
			c++
		}
		c++
	}

	var sum5, sum18 int
	for i := 1; i <= 18; i++ {
		size := getSubgridSize(g.size)
		ng := grid{make(map[image.Point]byte), g.size / size * (size + 1)}
		for x := 0; x < g.size; x += size {
			for y := 0; y < g.size; y += size {
				origin := image.Pt(x, y)
				subgrid := next(g.subgrid(origin, size), transforms)
				ng.merge(image.Pt(x/size*(size+1), y/size*(size+1)), subgrid)
			}
		}
		g = ng
		switch i {
		case 5:
			sum5 = count(g)
		case 18:
			sum18 = count(g)
		}
	}
	fmt.Println(sum5)
	fmt.Println(sum18)
}

func getSubgridSize(size int) int {
	if size%2 == 0 {
		return 2
	}
	return 3
}

func next(g grid, transform map[string]string) grid {
	for _, o := range orientations(g) {
		if e, ok := transform[o]; ok {
			return toGrid(e)
		}
	}
	log.Fatal("No match")
	return grid{}
}

func orientations(g grid) []string {
	var entries []string
	for i := 0; i < 4; i++ {
		entries = append(entries, toEntry(g.rotate()))
	}
	entries = append(entries, toEntry(g.flip()))
	for i := 0; i < 3; i++ {
		entries = append(entries, toEntry(g.rotate()))
	}
	return entries
}

func toEntry(g grid) string {
	var s []string
	for y := 0; y < g.size; y++ {
		var sb strings.Builder
		for x := 0; x < g.size; x++ {
			fmt.Fprintf(&sb, "%c", g.px[image.Pt(x, y)])
		}
		s = append(s, sb.String())
	}
	return strings.Join(s, "/")
}

func toGrid(entry string) grid {
	px := make(map[image.Point]byte)
	var x, y int
	for i := range entry {
		if entry[i] == '/' {
			y++
			x = 0
			continue
		}
		px[image.Pt(x, y)] = entry[i]
		x++
	}
	return grid{px, x}
}

func count(g grid) int {
	var sum int
	for _, b := range g.px {
		if b == '#' {
			sum++
		}
	}
	return sum
}

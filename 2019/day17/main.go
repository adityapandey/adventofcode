package main

import (
	"bufio"
	"fmt"
	"image"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}
	scaffolding, robot, dir := parse(program)

	fmt.Println(sumAlign(scaffolding))

	fmt.Println(dust(program, scaffolding, robot, dir))
}

func parse(program []int) (map[image.Point]struct{}, image.Point, util.Dir) {
	in := make(chan int)
	close(in)
	out := machine.Run(program, in)
	var sb strings.Builder
	for o := range out {
		fmt.Fprintf(&sb, "%c", o)
	}
	scaffolding := make(map[image.Point]struct{})
	var robot image.Point
	var dir util.Dir
	s := bufio.NewScanner(strings.NewReader(sb.String()))
	var y int
	for s.Scan() {
		line := s.Text()
		for x := 0; x < len(line); x++ {
			switch line[x] {
			case '^', 'v', '<', '>':
				robot = image.Pt(x, y)
				dir = util.DirFromByte(line[x])
				fallthrough
			case '#':
				scaffolding[image.Pt(x, y)] = struct{}{}
			}
		}
		y++
	}
	return scaffolding, robot, dir
}

func sumAlign(grid map[image.Point]struct{}) int {
	var sum int
	for p := range grid {
		x := true
		for _, n := range util.Neighbors4 {
			if _, ok := grid[p.Add(n)]; !ok {
				x = false
			}
		}
		if x {
			sum += p.X * p.Y
		}
	}
	return sum
}

func path(scaffolding map[image.Point]struct{}, robot image.Point, dir util.Dir) string {
	var dist int
	var d byte
	var sections []string
	for {
		if _, ok := scaffolding[robot.Add(dir.PointR())]; ok {
			robot = robot.Add(dir.PointR())
			dist++
			continue
		}
		if dist > 0 {
			sections = append(sections, fmt.Sprintf("%c,%d", d, dist))
		}
		if _, ok := scaffolding[robot.Add(dir.Next().PointR())]; ok {
			robot = robot.Add(dir.Next().PointR())
			dir = dir.Next()
			dist = 1
			d = 'R'
		} else if _, ok := scaffolding[robot.Add(dir.Prev().PointR())]; ok {
			robot = robot.Add(dir.Prev().PointR())
			dir = dir.Prev()
			dist = 1
			d = 'L'
		} else {
			break
		}
	}
	return strings.Join(sections, ",")
}

func encode(path string) (seq, a, b, c string) {
loop:
	for i := 2; i <= 21; i++ {
		for j := 2; j <= 21; j++ {
			for k := 2; k <= 21; k++ {
				next := path + ","
				a = next[:i]
				next = strings.ReplaceAll(next, a, "")
				b = next[:j]
				next = strings.ReplaceAll(next, b, "")
				c = next[:k]
				next = strings.ReplaceAll(next, c, "")
				if next == "" {
					break loop
				}
			}
		}
	}
	a, b, c = strings.Trim(a, ","), strings.Trim(b, ","), strings.Trim(c, ",")
	path = strings.ReplaceAll(path, a, "A")
	path = strings.ReplaceAll(path, b, "B")
	path = strings.ReplaceAll(path, c, "C")
	path = strings.Trim(path, ",")
	return path, a, b, c
}

func dust(program []int, scaffolding map[image.Point]struct{}, robot image.Point, dir util.Dir) int {
	seq, a, b, c := encode(path(scaffolding, robot, dir))
	input := fmt.Sprintf("%s\n%s\n%s\n%s\nn\n", seq, a, b, c)
	in := make(chan int, len(input))
	for i := 0; i < len(input); i++ {
		in <- int(input[i])
	}
	close(in)
	program[0] = 2
	out := machine.Run(program, in)
	var o int
	for o = range out {
	}
	return o
}

package main

import (
	"fmt"
	"image"
	"math"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var faces = map[int]image.Rectangle{
	1: image.Rect(50, 0, 100, 50),
	2: image.Rect(100, 0, 150, 50),
	3: image.Rect(50, 50, 100, 100),
	4: image.Rect(50, 100, 100, 150),
	5: image.Rect(0, 100, 50, 150),
	6: image.Rect(0, 150, 50, 200),
}

// For each (face, facing), func( (pos, facing) -> (new face, new pos, new facing) )
var transitions2d = map[int]map[util.Dir]func(image.Point, util.Dir) (int, image.Point, util.Dir){
	1: {
		util.N: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 4, image.Pt(p.X, 49), util.N },
		util.W: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 2, image.Pt(49, p.Y), util.W },
	},
	2: {
		util.N: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 2, image.Pt(p.X, 49), util.N },
		util.E: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 1, image.Pt(0, p.Y), util.E },
		util.S: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 2, image.Pt(p.X, 0), util.S },
	},
	3: {
		util.E: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 3, image.Pt(0, p.Y), util.E },
		util.W: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 3, image.Pt(49, p.Y), util.W },
	},
	4: {
		util.E: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 5, image.Pt(0, p.Y), util.E },
		util.S: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 1, image.Pt(p.X, 0), util.S },
	},
	5: {
		util.N: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 6, image.Pt(p.X, 49), util.N },
		util.W: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 4, image.Pt(49, p.Y), util.W },
	},
	6: {
		util.E: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 6, image.Pt(49, p.Y), util.E },
		util.S: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 5, image.Pt(p.X, 0), util.S },
		util.W: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 6, image.Pt(0, p.Y), util.W },
	},
}

// For each (face, facing), func( (pos, facing) -> (new face, new pos, new facing) )
var transitions3d = map[int]map[util.Dir]func(image.Point, util.Dir) (int, image.Point, util.Dir){
	1: {
		util.N: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 6, image.Pt(0, p.X), util.E },
		util.W: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 5, image.Pt(0, 49-p.Y), util.E },
	},
	2: {
		util.N: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 6, image.Pt(p.X, 49), util.N },
		util.E: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 4, image.Pt(49, 49-p.Y), util.W },
		util.S: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 3, image.Pt(49, p.X), util.W },
	},
	3: {
		util.E: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 2, image.Pt(p.Y, 49), util.N },
		util.W: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 5, image.Pt(p.Y, 0), util.S },
	},
	4: {
		util.E: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 2, image.Pt(49, 49-p.Y), util.W },
		util.S: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 6, image.Pt(49, p.X), util.W },
	},
	5: {
		util.N: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 3, image.Pt(0, p.X), util.E },
		util.W: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 1, image.Pt(0, 49-p.Y), util.E },
	},
	6: {
		util.E: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 4, image.Pt(p.Y, 49), util.N },
		util.S: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 2, image.Pt(p.X, 0), util.S },
		util.W: func(p image.Point, d util.Dir) (int, image.Point, util.Dir) { return 1, image.Pt(p.Y, 0), util.S },
	},
}

func main() {
	input := strings.Split(util.ReadFile("input"), "\n\n")
	grid := map[image.Point]byte{}
	y := 0
	for _, line := range strings.Split(input[0], "\n") {
		for x := range line {
			if line[x] == '.' || line[x] == '#' {
				grid[image.Pt(x, y)] = line[x]
			}
		}
		y++
	}
	start := image.Pt(math.MaxInt, 0)
	for p := range grid {
		if p.Y == 0 {
			start.X = util.Min(start.X, p.X)
		}
	}

	pos1, facing1 := start, util.DirFromByte('>')
	pos2, facing2 := pos1, facing1
	for _, path := range parse(input[1]) {
		switch path.dir {
		case 'R':
			facing1 = facing1.Next()
			facing2 = facing2.Next()
		case 'L':
			facing1 = facing1.Prev()
			facing2 = facing2.Prev()
		default:
			for i := 0; i < path.val; i++ {
				pos1, facing1 = next(pos1, facing1, grid, transitions2d)
				pos2, facing2 = next(pos2, facing2, grid, transitions3d)
			}
		}
	}

	score := map[util.Dir]int{util.E: 0, util.S: 1, util.W: 2, util.N: 3}
	fmt.Println(1000*(pos1.Y+1) + 4*(pos1.X+1) + score[facing1])
	fmt.Println(1000*(pos2.Y+1) + 4*(pos2.X+1) + score[facing2])
}

func next(pos image.Point, facing util.Dir, grid map[image.Point]byte, transition map[int]map[util.Dir]func(image.Point, util.Dir) (int, image.Point, util.Dir)) (image.Point, util.Dir) {
	newPos := pos.Add(facing.PointR())
	newFacing := facing
	if _, ok := grid[newPos]; !ok {
		var face int
		for f := range faces {
			if pos.In(faces[f]) {
				face = f
				break
			}
		}
		var newFace int
		newFace, newPos, newFacing = transition[face][facing](pos.Sub(faces[face].Min), facing)
		newPos = newPos.Add(faces[newFace].Min)
	}
	if grid[newPos] == '#' {
		return pos, facing
	}
	return newPos, newFacing
}

type instr struct {
	val int
	dir byte
}

func parse(path string) []instr {
	val := 0
	var instrs []instr
	for i := range path {
		switch path[i] {
		case 'R', 'L':
			instrs = append(instrs, instr{val: val})
			instrs = append(instrs, instr{dir: path[i]})
			val = 0
		default:
			val *= 10
			val += int(path[i]) - '0'
		}
	}
	instrs = append(instrs, instr{val: val})
	return instrs
}

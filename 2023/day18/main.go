package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	s := util.ReadAll()
	fmt.Println(area(s, false))
	fmt.Println(area(s, true))
}

func area(s string, swap bool) int {
	sum := 0
	curr := image.Pt(0, 0)
	for _, line := range strings.Split(s, "\n") {
		sp := strings.Fields(line)
		var d util.Dir
		var n int
		if swap {
			d = map[byte]util.Dir{
				'0': util.E,
				'1': util.S,
				'2': util.W,
				'3': util.N,
			}[sp[2][7]]
			nn, _ := strconv.ParseInt(sp[2][2:7], 16, 64)
			n = int(nn)
		} else {
			d = util.DirFromByte(sp[0][0])
			n = util.Atoi(sp[1])
		}
		next := curr.Add(d.PointR().Mul(n))
		sum += curr.X*next.Y - curr.Y*next.X + n
		curr = next
	}
	return sum/2 + 1
}

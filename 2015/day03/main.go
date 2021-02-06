package main

import (
	"fmt"
	"image"

	"github.com/adityapandey/adventofcode/util"
)

var dirmap = map[byte]util.Dir{}

func main() {
	var p image.Point
	var santa, robosanta image.Point
	visited1 := map[image.Point]struct{}{p: {}}
	visited2 := map[image.Point]struct{}{p: {}}
	input := util.ReadAll()
	for i := range input {
		p = p.Add(util.DirFromByte(input[i]).Point())
		visited1[p] = struct{}{}
		switch i % 2 {
		case 0:
			santa = santa.Add(util.DirFromByte(input[i]).Point())
			visited2[santa] = struct{}{}
		case 1:
			robosanta = robosanta.Add(util.DirFromByte(input[i]).Point())
			visited2[robosanta] = struct{}{}
		}
	}
	fmt.Println((len(visited1)))
	fmt.Println((len(visited2)))
}

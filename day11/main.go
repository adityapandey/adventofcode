package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode2017-go/util"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	dirs := map[string]image.Point{
		"n":  image.Pt(0, 2),
		"ne": image.Pt(1, 1),
		"se": image.Pt(1, -1),
		"s":  image.Pt(0, -2),
		"sw": image.Pt(-1, -1),
		"nw": image.Pt(-1, 1),
	}
	furthest := 0
	var p image.Point
	for _, d := range strings.Split(string(input), ",") {
		p = p.Add(dirs[d])
		if distance(p) > furthest {
			furthest = distance(p)
		}
	}
	fmt.Println(distance(p))
	fmt.Println(furthest)
}

func distance(p image.Point) int {
	return (util.Abs(p.X) + util.Abs(p.Y)) / 2
}

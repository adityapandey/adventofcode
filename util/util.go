package util

import (
	"image"
	"log"
	"strconv"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ManhattanDistance(p1, p2 image.Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

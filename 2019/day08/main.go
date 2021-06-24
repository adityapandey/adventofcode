package main

import (
	"fmt"
	"image"
	"math"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := []byte(util.ReadAll())
	var layers []map[image.Point]byte
	for i := 0; i < len(input); {
		layer := make(map[image.Point]byte)
		for y := 0; y < 6; y++ {
			for x := 0; x < 25; x++ {
				layer[image.Pt(x, y)] = input[i]
				i++
			}
		}
		layers = append(layers, layer)
	}

	minZeros := math.MaxInt32
	var minZerosLayer int
	for i := range layers {
		var zeros int
		for _, b := range layers[i] {
			if b == '0' {
				zeros++
			}
		}
		if zeros < minZeros {
			minZeros, minZerosLayer = zeros, i
		}
	}
	var ones, twos int
	for _, b := range layers[minZerosLayer] {
		switch b {
		case '1':
			ones++
		case '2':
			twos++
		}
	}
	fmt.Println(ones * twos)

	for y := 0; y < 6; y++ {
		for x := 0; x < 25; x++ {
			var lit bool
			for _, layer := range layers {
				p := image.Pt(x, y)
				if layer[p] != '2' {
					lit = layer[p] == '1'
					break
				}
			}
			if lit {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

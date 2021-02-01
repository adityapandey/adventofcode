package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

const l = 8

func main() {
	b := []byte("abcdefgh")
	input := strings.Split(util.ReadAll(), "\n")
	m := rotationMap()
	for _, line := range input {
		b = apply(line, b, false, m)
	}
	fmt.Println(string(b))
	b = []byte("fbgdceah")
	for i := len(input) - 1; i >= 0; i-- {
		b = apply(input[i], b, true, m)
	}
	fmt.Println(string(b))
}

func rotationMap() map[bool]map[int]int {
	m := map[bool]map[int]int{false: {}, true: {}}
	for i := 0; i < l; i++ {
		rot := i + 1
		if i >= 4 {
			rot++
		}
		m[false][i] = rot
		m[true][(i+rot)%l] = rot
	}
	return m
}

func apply(rule string, b []byte, unscramble bool, rotationMap map[bool]map[int]int) []byte {
	f := strings.Fields(rule)
	switch f[0] {
	case "swap":
		switch f[1] {
		case "position":
			x, y := util.Atoi(f[2]), util.Atoi(f[5])
			b[x], b[y] = b[y], b[x]
		case "letter":
			x, y := f[2][0], f[5][0]
			i, j := bytes.IndexByte(b, x), bytes.IndexByte(b, y)
			b[i], b[j] = b[j], b[i]
		}
	case "rotate":
		c := make([]byte, l)
		copy(c, b)
		var rot int
		var left bool
		switch f[1] {
		case "left":
			rot = util.Atoi(f[2])
			left = !unscramble
		case "right":
			rot = util.Atoi(f[2])
			left = unscramble
		case "based":
			x := bytes.IndexByte(b, f[6][0])
			rot = rotationMap[unscramble][x]
			left = unscramble
		}
		for i := range b {
			if left {
				b[i] = c[(i+rot)%l]
			} else {
				b[(i+rot)%l] = c[i]
			}
		}
	case "reverse":
		x, y := util.Atoi(f[2]), util.Atoi(f[4])
		bs := b[x : y+1]
		for i := 0; i < len(bs)/2; i++ {
			bs[i], bs[len(bs)-i-1] = bs[len(bs)-i-1], bs[i]
		}
	case "move":
		x, y := util.Atoi(f[2]), util.Atoi(f[5])
		if unscramble {
			x, y = y, x
		}
		c := b[x]
		b = append(b[:x], b[x+1:]...)
		b = append(b[:y], append([]byte{c}, b[y:]...)...)
	}
	return b
}

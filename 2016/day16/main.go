package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var a []bool
	input := util.ReadAll()
	for i := range input {
		a = append(a, input[i] == '1')
	}
	fmt.Println(str(checksum(a, 272)))
	fmt.Println(str(checksum(a, 35651584)))
}

func step(a []bool) []bool {
	b := make([]bool, 2*len(a)+1)
	for i := range b {
		switch {
		case i < len(a):
			b[i] = a[i]
		case i == len(a):
			b[i] = false
		case i > len(a):
			b[i] = !a[2*len(a)-i]
		}
	}
	return b
}

func checksum(a []bool, diskLen int) []bool {
	for len(a) < diskLen {
		a = step(a)
	}
	a = a[:diskLen]
	l := diskLen / 2
	cksum := make([]bool, l)
	for {
		for i := 0; i < 2*l; i += 2 {
			cksum[i/2] = a[i] == a[i+1]
		}
		if l%2 != 0 {
			return cksum[:l]
		}
		a = cksum[:l]
		l /= 2
	}
}

func str(a []bool) string {
	m := map[bool]byte{false: '0', true: '1'}
	var sb strings.Builder
	for _, b := range a {
		fmt.Fprintf(&sb, "%c", m[b])
	}
	return sb.String()
}

package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var code, decode, encode int
	s := util.ScanAll()
	for s.Scan() {
		code += len(s.Text())
		decode += decodeLen(s.Text())
		encode += encodeLen(s.Text())
	}
	fmt.Println(code - decode)
	fmt.Println(encode - code)
}

func decodeLen(s string) int {
	var l int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '"':
		case '\\':
			switch s[i+1] {
			case '\\':
				l++
				i++
			case '"':
				l++
				i++
			case 'x':
				l++
				i += 3
			}
		default:
			l++
		}
	}
	return l
}

func encodeLen(s string) int {
	return len(s) + strings.Count(s, `\`) + strings.Count(s, `"`) + len(`""`)
}

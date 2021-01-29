package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := util.ReadAll()
	hashes := make(map[int]string)
	stretchHashes := make(map[int]string)
	fmt.Println(findKey(hashes, input, false))
	fmt.Println(findKey(stretchHashes, input, true))
}

func findKey(hashes map[int]string, s string, stretch bool) int {
	var suffix, key int
	for key < 64 {
		b, tripleIndex := findTriple(hashes, s, suffix, stretch)
		suffix = tripleIndex + 1
		if found := findPentuple(hashes, s, suffix, b, 1000, stretch); found {
			key++
		}
	}
	return suffix - 1
}

func hash(hashes map[int]string, s string, suffix int, stretch bool) string {
	if h, ok := hashes[suffix]; ok {
		return h
	}
	h := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, suffix))))
	if stretch {
		for i := 0; i < 2016; i++ {
			h = fmt.Sprintf("%x", md5.Sum([]byte(h)))
		}
	}
	hashes[suffix] = h
	return h
}

func findTriple(hashes map[int]string, s string, suffix int, stretch bool) (byte, int) {
	for {
		h := hash(hashes, s, suffix, stretch)
		for i := 0; i < len(h)-2; i++ {
			if h[i] == h[i+1] && h[i] == h[i+2] {
				return h[i], suffix
			}
		}
		suffix++
	}
}

func findPentuple(hashes map[int]string, s string, start int, b byte, steps int, stretch bool) bool {
	pentuple := string([]byte{b, b, b, b, b})
	for i := 0; i < steps; i++ {
		h := hash(hashes, s, start+i, stretch)
		if strings.Contains(h, pentuple) {
			return true
		}
	}
	return false
}

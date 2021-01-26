package main

import (
	"bytes"
	"crypto/md5"
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

var zeros = []byte("00000")

func main() {
	input := util.ReadAll()
	var password1 []byte
	password2 := make(map[int]byte)
	for i := 0; len(password1) < 8 || len(password2) < 8; i++ {
		h := []byte(fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))))
		if bytes.Equal(h[:5], zeros) {
			if len(password1) < 8 {
				password1 = append(password1, h[5])
			}
			if h[5] >= '0' && h[5] <= '7' && len(password2) < 8 {
				pos := int(h[5] - '0')
				if _, ok := password2[pos]; !ok {
					password2[pos] = h[6]
				}
			}
		}
	}
	fmt.Println(string(password1))
	for i := 0; i < 8; i++ {
		fmt.Printf("%c", password2[i])
	}
	fmt.Println()
}

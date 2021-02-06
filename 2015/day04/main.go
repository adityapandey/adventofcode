package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := util.ReadAll()
	var hash5, hash6 int
	for i := 0; hash5 == 0 || hash6 == 0; i++ {
		if hash5 == 0 && strings.HasPrefix(hash(fmt.Sprintf("%s%d", input, i)), "00000") {
			hash5 = i
		}
		if hash6 == 0 && strings.HasPrefix(hash(fmt.Sprintf("%s%d", input, i)), "000000") {
			hash6 = i
		}
	}
	fmt.Println(hash5)
	fmt.Println(hash6)
}

func hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

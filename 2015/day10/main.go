package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	a := util.ReadAll()
	for i := 0; i < 40; i++ {
		a = looksay(a)
	}
	fmt.Println(len(a))
	for i := 0; i < 10; i++ {
		a = looksay(a)
	}
	fmt.Println(len(a))
}

func looksay(a string) string {
	var sb strings.Builder
	prev, count := a[0], 1
	for i := 1; i < len(a); i++ {
		if a[i] == prev {
			count++
		} else {
			fmt.Fprintf(&sb, "%d%c", count, prev)
			prev, count = a[i], 1
		}
	}
	fmt.Fprintf(&sb, "%d%c", count, prev)
	return sb.String()
}

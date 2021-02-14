package main

import (
	"fmt"
	"os"
)

func main() {
	var r, c int
	fmt.Fscanf(os.Stdin, "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.", &r, &c)
	n := nth(r, c)
	code := 20151125
	for i := 1; i < n; i++ {
		code *= 252533
		code %= 33554393
	}
	fmt.Println(code)
}

func nth(r, c int) int {
	n := c * (c + 1) / 2
	for y := 2; y <= r; y++ {
		n += y - 2 + c
	}
	return n
}

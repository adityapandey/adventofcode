package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func ReadAll() string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func ScanAll() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

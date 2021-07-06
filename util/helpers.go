package util

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
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

func Min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func Gcd(a, b int) int {
	a, b = Abs(a), Abs(b)
	if a == 0 || b == 0 {
		return a + b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a * (b / Gcd(a, b))
}

func ReadAll() string {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func ScanAll() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

func ReadFile(fn string) string {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(input)
}

package util

import (
	"bufio"
	"io/ioutil"
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

func ReadAll() string {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func ScanAll() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type decl struct {
	ans []string
}

func main() {
	var decls []decl
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range strings.Split(string(input), "\n\n") {
		var d decl
		for _, line := range strings.Split(t, "\n") {
			d.ans = append(d.ans, line)
		}
		decls = append(decls, d)
	}

	// Part 1
	sum := 0
	for _, d := range decls {
		sum += anyYes(d)
	}
	fmt.Println(sum)

	// Part 2
	sum = 0
	for _, d := range decls {
		sum += allYes(d)
	}
	fmt.Println(sum)
}

func anyYes(d decl) int {
	m := make(map[byte]struct{})
	for _, a := range d.ans {
		for i := 0; i < len(a); i++ {
			m[a[i]] = struct{}{}
		}
	}
	return len(m)
}

func allYes(d decl) int {
	m := make(map[byte]int)
	for _, a := range d.ans {
		for i := 0; i < len(a); i++ {
			m[a[i]]++
		}
	}
	sum := 0
	for _, v := range m {
		if v == len(d.ans) {
			sum++
		}
	}
	return sum
}

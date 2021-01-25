package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	sum1, sum2 := 0, 0
	for s.Scan() {
		if valid(s.Text()) {
			sum1++
		}
		if valid2(s.Text()) {
			sum2++
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func valid(s string) bool {
	words := strings.Split(s, " ")
	m := make(map[string]struct{})
	for _, w := range words {
		if _, ok := m[w]; ok {
			return false
		}
		m[w] = struct{}{}
	}
	return true
}

func valid2(s string) bool {
	words := strings.Split(s, " ")
	m := make(map[string]map[byte]int)
	for i := range words {
		if _, ok := m[words[i]]; ok {
			return false
		}
		m[words[i]] = make(map[byte]int)
		for j := 0; j < len(words[i]); j++ {
			m[words[i]][words[i][j]]++
		}
		for j := i - 1; j >= 0; j-- {
			if samemap(m[words[i]], m[words[j]]) {
				return false
			}
		}
	}
	return true
}

func samemap(m1, m2 map[byte]int) bool {
	fmt.Println(m1, m2)
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

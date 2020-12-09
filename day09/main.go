// https://adventofcode.com/2020/day/9
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	var input []int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var i int
		fmt.Sscanf(s.Text(), "%d", &i)
		input = append(input, i)
	}

	// Part 1
	var invalid int
	preambleLen := 25
	m := make(map[int]struct{})
	for _, i := range input[:preambleLen] {
		m[i] = struct{}{}
	}
	for i := preambleLen; i < len(input); i++ {
		if !sumFromSet(input[i], m) {
			invalid = input[i]
			fmt.Println(invalid)
			break
		}
		m[input[i]] = struct{}{}
		delete(m, input[i-preambleLen])
	}

	// Part 2
	i, j, err := findRangeSum(invalid, input)
	if err != nil {
		log.Fatal(err)
	}
	o := make([]int, j+1-i)
	copy(o, input[i:j+1])
	sort.Ints(o)
	fmt.Println(o[0] + o[len(o)-1])
}

func sumFromSet(n int, m map[int]struct{}) bool {
	for k := range m {
		if _, ok := m[n-k]; ok {
			return true
		}
	}
	return false
}

func findRangeSum(n int, arr []int) (int, int, error) {
	for i := 0; i < len(arr)-1; i++ {
		sum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			if sum == n {
				return i, j, nil
			}
			if sum > n {
				break
			}
		}
	}
	return -1, -1, errors.New("No viable range")
}

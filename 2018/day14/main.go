package main

import (
	"fmt"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var input string
	fmt.Fscanf(os.Stdin, "%s", &input)

	// Part 1
	iters := util.Atoi(input)
	scores := []int{3, 7}
	i, j := 0, 1
	for len(scores) < iters+11 {
		sum := scores[i] + scores[j]
		if sum >= 10 {
			scores = append(scores, sum/10)
		}
		scores = append(scores, sum%10)
		i += scores[i] + 1
		i %= len(scores)
		j += scores[j] + 1
		j %= len(scores)
	}
	for _, n := range scores[iters : iters+10] {
		fmt.Print(n)
	}
	fmt.Println()

	// Part 2
	arr := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		arr[i] = int(input[i] - '0')
	}

	scores = []int{3, 7}
	i, j = 0, 1
	for {
		sum := scores[i] + scores[j]
		if sum >= 10 {
			scores = append(scores, sum/10)
			if matches(scores, arr) {
				return
			}
		}
		scores = append(scores, sum%10)
		if matches(scores, arr) {
			return
		}
		i += scores[i] + 1
		i %= len(scores)
		j += scores[j] + 1
		j %= len(scores)
	}
}

func matches(scores, arr []int) bool {
	if len(scores) < len(arr) {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if scores[len(scores)-1-i] != arr[len(arr)-1-i] {
			return false
		}
	}
	fmt.Println(len(scores) - len(arr))
	return true
}

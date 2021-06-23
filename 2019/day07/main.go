package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}

	var max int
	for _, phase := range permutations([]int{0, 1, 2, 3, 4}) {
		max = util.Max(max, signal(program, phase))
	}
	fmt.Println(max)

	max = 0
	for _, phase := range permutations([]int{5, 6, 7, 8, 9}) {
		max = util.Max(max, signalFeedback(program, phase))
	}
	fmt.Println(max)
}

func permutations(arr []int) [][]int {
	res := [][]int{}
	var helper func([]int, int)
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func signal(program []int, phase []int) int {
	var prevOut int
	for i := 0; i < 5; i++ {
		in := make(chan int, 2)
		in <- phase[i]
		in <- prevOut
		close(in)
		out := machine.Run(program, in)
		prevOut = <-out
	}
	return prevOut
}

func signalFeedback(program []int, phase []int) int {
	in0 := make(chan int, 2)
	in0 <- phase[0]
	in0 <- 0
	out := make(chan int)
	m := machine.New(program, in0, out)
	go m.Run()

	for i := 1; i < 5; i++ {
		in := make(chan int)
		go func(i int, in chan<- int, prevOut <-chan int) {
			in <- phase[i]
			for o := range prevOut {
				in <- o
			}
			close(in)
		}(i, in, out)
		out = make(chan int)
		m = machine.New(program, in, out)
		go m.Run()
	}

	var o int
	for o = range out {
		in0 <- o
	}
	close(in0)
	return o
}

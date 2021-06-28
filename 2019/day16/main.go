package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := util.ReadAll()
	signal := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		signal[i] = int(input[i] - '0')
	}

	output := make([]int, len(signal))
	copy(output, signal)
	for steps := 0; steps < 100; steps++ {
		output = fft(output)
	}
	fmt.Println(output[:8])

	fmt.Println(fftRepeat(signal))
}

func fft(signal []int) []int {
	output := make([]int, len(signal))
	n := len(signal)
	m := map[int]int{0: 0, 1: 1, 2: 0, 3: -1}
	for i := 0; i < n; i++ {
		for j := 1; j <= n; j++ {
			output[i] += m[(j/(i+1))%4] * signal[j-1]
		}
		output[i] %= 10
		output[i] = util.Abs(output[i])
	}
	return output
}

func fftRepeat(signal []int) []int {
	var offset int
	for i := 0; i < 7; i++ {
		offset *= 10
		offset += signal[i]
	}
	l := len(signal)
	n := 10000 * l
	// If offset >= n/2, output digits can be computed from the back of the
	// input, and depend only on the output beyond that point. The delta
	// between coefficients of the "FFT" looks like this (n=8):
	// 0:  +-= ++
	// 1:   + ----
	// 2:    +  --
	// 3:     +   -
	// 4:      +     <- n/2
	// 5:       +
	// 6:        +
	// 7:         +
	if offset < n/2 {
		panic("offset too large")
	}
	output := make([]int, n-offset)
	for i := n - 1; i >= offset; i-- {
		output[i-offset] = signal[i%l]
	}
	for steps := 0; steps < 100; steps++ {
		for i := len(output) - 2; i >= 0; i-- {
			output[i] = (output[i+1] + output[i]) % 10
		}
	}
	return output[:8]
}

package main

import (
	"fmt"
	"os"
)

func main() {
	var grid int
	fmt.Fscanf(os.Stdin, "%d", &grid)

	const S = 300
	var arr [S][S]int

	for x := 0; x < S; x++ {
		for y := 0; y < S; y++ {
			rack := x + 11
			arr[x][y] = ((((rack*(y+1) + grid) * rack) / 100) % 10) - 5
		}
	}
	var maxX, maxY, maxSum int
	for x := 0; x < S-3; x++ {
		for y := 0; y < S-3; y++ {
			var sum int
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					sum += arr[x+i][y+j]
				}
			}
			if sum > maxSum {
				maxX, maxY, maxSum = x, y, sum
			}
		}
	}
	fmt.Printf("%d,%d\n", maxX+1, maxY+1)

	var sums [S][S]int
	maxX, maxY, maxSum, maxSize := 0, 0, 0, 0
	for size := 1; size <= S; size++ {
		for x := 0; x < S-size+1; x++ {
			for y := 0; y < S-size+1; y++ {
				if size > 1 {
					for i := 0; i < size; i++ {
						sums[x][y] += arr[x+size-1][y+i] + arr[x+i][y+size-1]
					}
				}
				sums[x][y] += arr[x+size-1][y+size-1]
				if sums[x][y] > maxSum {
					maxX, maxY, maxSum, maxSize = x, y, sums[x][y], size
				}
			}
		}
	}
	fmt.Printf("%d,%d,%d\n", maxX+1, maxY+1, maxSize)
}

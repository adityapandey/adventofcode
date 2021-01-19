package main

import (
	"fmt"
	"os"
)

var factors = [2]uint64{16807, 48271}

func main() {
	var val [2]uint64
	fmt.Fscanf(os.Stdin, "Generator A starts with %d", &val[0])
	fmt.Fscanf(os.Stdin, "Generator B starts with %d", &val[1])

	fmt.Println(judge(val, [2]uint64{1, 1}, 40000000))
	fmt.Println(judge(val, [2]uint64{4, 8}, 5000000))
}

func judge(val, criteria [2]uint64, iter int) int {
	var c int
	for i := 0; i < iter; i++ {
		val = next(val, criteria)
		if val[0]&0xFFFF == val[1]&0xFFFF {
			c++
		}
	}
	return c
}

func next(prev [2]uint64, criteria [2]uint64) [2]uint64 {
	a := prev
	for i := range prev {
		for {
			a[i] *= factors[i]
			// 2^n === 1 (mod 2^n-1)
			// So if x = p*2^n + q, then x === p + q (mod 2^n-1)
			// p = x << n, q = x && 2^n-1
			// p + q can overflow, so mod again until < 2^n - 1
			for a[i] > 0x7FFFFFFF {
				a[i] = (a[i] & 0x7FFFFFFF) + (a[i] >> 31)
			}
			if a[i]%criteria[i] == 0 {
				break
			}
		}
	}
	return a
}

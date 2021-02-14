package main

import (
	"fmt"
	"os"
)

func main() {
	var presents int
	fmt.Fscanf(os.Stdin, "%d", &presents)
	i := 1
	for numPresents1(i) < presents {
		i++
	}
	fmt.Println(i)
	i = 1
	for numPresents2(i) < presents {
		i++
	}
	fmt.Println(i)
}

func numPresents1(n int) int {
	var sum int
	for _, f := range factors(n) {
		sum += f
	}
	return 10 * sum
}

func numPresents2(n int) int {
	f := factors(n)
	var sum int
	for i := 0; i < len(f); i++ {
		if n > 50*f[i] {
			f = append(f[:i], f[i+1:]...)
			i--
		} else {
			sum += f[i]
		}
	}
	return 11 * sum
}

func factors(n int) []int {
	f := []int{1}
	expand := func(prime, power int) {
		l := len(f)
		i, prod := 0, prime
		for i < power {
			for j := 0; j < l; j++ {
				f = append(f, f[j]*prod)
			}
			i++
			prod *= prime
		}
	}
	var power int
	for n&1 == 0 {
		n >>= 1
		power++
	}
	expand(2, power)
	for p := 3; n > 1; p += 2 {
		if p*p > n {
			p = n
		}
		for power = 0; n%p == 0; power++ {
			n /= p
		}
		if power > 0 {
			expand(p, power)
		}
	}
	return f
}

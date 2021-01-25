package main

import (
	"fmt"
	"os"
)

type ring struct {
	a    []int
	l    int
	curr int
}

func (r *ring) next(n int) {
	r.curr += n
	r.curr %= r.l
}

func (r *ring) add(n int) {
	r.curr++
	r.a = append(r.a[:r.curr], append([]int{n}, r.a[r.curr:]...)...)
	r.l++
}

func main() {
	var step int
	fmt.Fscanf(os.Stdin, "%d", &step)
	r := &ring{a: make([]int, 1), l: 1}
	r.a[0] = 0
	for i := 1; i <= 2017; i++ {
		r.next(step)
		r.add(i)
	}
	fmt.Println(r.a[(r.curr+1)%r.l])

	for i := 2018; i <= 50000000; i++ {
		r.next(step)
		if r.curr == 0 {
			r.add(i)
		} else {
			r.curr++
			r.l++
		}
	}
	fmt.Println(r.a[1])
}

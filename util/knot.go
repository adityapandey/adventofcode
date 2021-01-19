package util

import (
	"fmt"
	"strings"
)

type Knot struct {
	List     []int
	currPos  int
	stepSize int
}

func NewKnot(n int) *Knot {
	r := &Knot{List: make([]int, n)}
	for i := 0; i < n; i++ {
		r.List[i] = i
	}
	return r
}

func (r *Knot) get(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = r.List[(r.currPos+i)%len(r.List)]
	}
	return a
}

func (r *Knot) set(a []int) {
	for i := range a {
		r.List[(r.currPos+i)%len(r.List)] = a[i]
	}
}

func (r *Knot) move(n int) {
	r.currPos = (r.currPos + n + r.stepSize) % len(r.List)
	r.stepSize++
}

func (r *Knot) Hash(l int) {
	a := r.get(l)
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	r.set(a)
	r.move(l)
}

func (r *Knot) DenseHash(input []byte) string {
	input = append(input, []byte{17, 31, 73, 47, 23}...)
	for i := 0; i < 64; i++ {
		for _, c := range input {
			r.Hash(int(c))
		}
	}
	a := make([]int, 16)
	for i := 0; i < 16; i++ {
		xor := r.List[16*i]
		for j := 1; j < 16; j++ {
			xor ^= r.List[16*i+j]
		}
		a[i] = xor
	}
	s := make([]string, 16)
	for i := 0; i < 16; i++ {
		s[i] = fmt.Sprintf("%02x", a[i])
	}
	return strings.Join(s, "")
}

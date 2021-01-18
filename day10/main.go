package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode2017-go/util"
)

type ring struct {
	list     []int
	currPos  int
	stepSize int
}

func newRing(n int) *ring {
	r := &ring{list: make([]int, n)}
	for i := 0; i < n; i++ {
		r.list[i] = i
	}
	return r
}

func (r *ring) get(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = r.list[(r.currPos+i)%len(r.list)]
	}
	return a
}

func (r *ring) set(a []int) {
	for i := range a {
		r.list[(r.currPos+i)%len(r.list)] = a[i]
	}
}

func (r *ring) move(n int) {
	r.currPos = (r.currPos + n + r.stepSize) % len(r.list)
	r.stepSize++
}

func (r *ring) hash(l int) {
	a := r.get(l)
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	r.set(a)
	r.move(l)
}

func (r *ring) denseHash() string {
	a := make([]int, 16)
	for i := 0; i < 16; i++ {
		xor := r.list[16*i]
		for j := 1; j < 16; j++ {
			xor ^= r.list[16*i+j]
		}
		a[i] = xor
	}
	s := make([]string, 16)
	for i := 0; i < 16; i++ {
		s[i] = fmt.Sprintf("%x", a[i])
	}
	return strings.Join(s, "")
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var lengths []int
	for _, s := range strings.Split(string(input), ",") {
		lengths = append(lengths, util.Atoi(s))
	}

	// Part 1
	listLen := 256
	r := newRing(listLen)
	for _, l := range lengths {
		r.hash(l)
	}
	fmt.Println(r.list[0] * r.list[1])

	// Part 2
	r = newRing(listLen)
	input = append(input, []byte{17, 31, 73, 47, 23}...)
	for i := 0; i < 64; i++ {
		for _, c := range input {
			r.hash(int(c))
		}
	}
	fmt.Println(r.denseHash())
}

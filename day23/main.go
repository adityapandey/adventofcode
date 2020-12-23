// https://adventofcode.com/2020/day/23
package main

import (
	"container/ring"
	"fmt"
	"os"
)

func main() {
	var input string
	fmt.Fscanln(os.Stdin, &input)

	// Part 1
	r := ring.New(9)
	m := make(map[int]*ring.Ring)
	for i := 0; i < len(input); i++ {
		n := int(input[i]) - '0'
		m[n] = r
		r.Value = n
		r = r.Next()
	}

	for i := 0; i < 100; i++ {
		move(r, m, len(input))
		r = r.Next()
	}
	r = m[1]
	var remains []int
	r.Do(func(i interface{}) { remains = append(remains, i.(int)) })
	for _, i := range remains[1:] {
		fmt.Printf("%d", i)
	}
	fmt.Println()

	// Part 2
	r = ring.New(1000000)
	m = make(map[int]*ring.Ring)
	for i := 0; i < len(input); i++ {
		n := int(input[i]) - '0'
		m[n] = r
		r.Value = n
		r = r.Next()
	}
	for i := 10; i <= 1000000; i++ {
		m[i] = r
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < 10000000; i++ {
		move(r, m, 1000000)
		r = r.Next()
	}
	r = m[1]
	fmt.Println(r.Next().Value.(int) * r.Next().Next().Value.(int))
}

func move(r *ring.Ring, m map[int]*ring.Ring, max int) {
	pickedCups := r.Unlink(3)
	picked := make(map[int]struct{})
	pickedCups.Do(func(i interface{}) { picked[i.(int)] = struct{}{} })
	dest := r.Value.(int)
	for {
		dest--
		if dest == 0 {
			dest = max
		}
		if _, ok := picked[dest]; !ok {
			break
		}
	}
	m[dest].Link(pickedCups)
}

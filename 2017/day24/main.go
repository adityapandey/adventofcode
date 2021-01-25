package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type adapter struct {
	a, b int
}

func (a adapter) reverse(p int) int {
	if p == a.a {
		return a.b
	}
	return a.a
}

func (a adapter) strength() int {
	return a.a + a.b
}

func main() {
	var adpts []adapter
	pins := make(map[int][]int)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var a adapter
		fmt.Sscanf(s.Text(), "%d/%d", &a.a, &a.b)
		adpts = append(adpts, a)
	}
	for i, a := range adpts {
		if a.a == a.b {
			pins[a.a] = append(pins[a.a], i)
		} else {
			pins[a.a] = append(pins[a.a], i)
			pins[a.b] = append(pins[a.b], i)
		}
	}

	chains := connect(0, pins, adpts, map[int]struct{}{})
	var strongest int
	for _, c := range chains {
		var sum int
		for _, a := range c {
			sum += adpts[a].strength()
		}
		if sum > strongest {
			strongest = sum
		}
	}
	fmt.Println(strongest)

	sort.Slice(chains, func(i, j int) bool {
		var sumi, sumj int
		for _, a := range chains[i] {
			sumi += adpts[a].strength()
		}
		for _, a := range chains[j] {
			sumj += adpts[a].strength()
		}
		leni := len(chains[i])
		lenj := len(chains[j])
		if leni == lenj {
			return sumi > sumj
		}
		return leni > lenj
	})
	var strength int
	for _, a := range chains[0] {
		strength += adpts[a].strength()
	}
	if strength > strongest {
		strongest = strength
	}
	fmt.Println(strength)
}

func connect(p int, pins map[int][]int, adpts []adapter, used map[int]struct{}) [][]int {
	if len(pins[p]) == 0 {
		return [][]int{{}}
	}
	chains := [][]int{{}}
	candidates := pins[p]
	for _, c := range candidates {
		if _, ok := used[c]; ok {
			continue
		}
		u := copymap(used)
		u[c] = struct{}{}
		subchains := connect(adpts[c].reverse(p), pins, adpts, u)
		for i := range subchains {
			chains = append(chains, append([]int{c}, subchains[i]...))
		}
	}
	return chains
}

func copymap(m map[int]struct{}) map[int]struct{} {
	c := make(map[int]struct{})
	for k := range m {
		c[k] = struct{}{}
	}
	return c
}

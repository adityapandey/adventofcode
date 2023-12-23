package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/adityapandey/adventofcode/util"
	"golang.org/x/exp/maps"
)

type pulse bool

const (
	low  pulse = false
	high pulse = true
)

func (p pulse) String() string {
	if p {
		return "high"
	}
	return "low"
}

type op struct {
	src  string
	p    pulse
	dest string
}

func (o op) String() string {
	return fmt.Sprintf("%v -%v-> %v", o.src, o.p, o.dest)
}

type module interface {
	signal(src string, p pulse) []op
}
type flipflop struct {
	on   bool
	dest []string
}

func (f *flipflop) signal(src string, p pulse) []op {
	if p == high {
		return []op{}
	}
	f.on = !f.on
	var ops []op
	for _, d := range f.dest {
		ops = append(ops, op{p: pulse(f.on), dest: d})
	}
	return ops
}

type conjunction struct {
	inputs map[string]pulse
	dest   []string
}

func (c *conjunction) signal(src string, p pulse) []op {
	c.inputs[src] = p
	ret := low
	if slices.Contains(maps.Values(c.inputs), low) {
		ret = high
	}
	var ops []op
	for _, d := range c.dest {
		ops = append(ops, op{p: ret, dest: d})
	}
	return ops
}

type broadcast struct {
	dest []string
}

func (b *broadcast) signal(src string, p pulse) []op {
	var ops []op
	for _, d := range b.dest {
		ops = append(ops, op{p: p, dest: d})
	}
	return ops
}

func main() {
	modules := map[string]module{}
	parents := map[string][]string{}
	for _, line := range strings.Split(util.ReadAll(), "\n") {
		sp := strings.Split(line, " -> ")
		var id string
		dest := strings.Split(sp[1], ", ")
		switch sp[0][0] {
		case '%':
			id = sp[0][1:]
			modules[id] = &flipflop{dest: dest}
		case '&':
			id = sp[0][1:]
			modules[id] = &conjunction{inputs: map[string]pulse{}, dest: dest}
		default:
			id = sp[0]
			modules[id] = &broadcast{dest: dest}
		}
		for _, d := range dest {
			parents[d] = append(parents[d], id)
		}
	}
	for m := range modules {
		if c, ok := modules[m].(*conjunction); ok {
			for _, d := range parents[m] {
				c.inputs[d] = low
			}
		}
	}

	c := map[pulse]int{}
	cyclelen := map[string]int{}
	parentsOfParents := parents[parents["rx"][0]]
	for i := 0; len(cyclelen) < len(parentsOfParents); i++ {
		q := []op{{src: "button", p: low, dest: "broadcaster"}}
		for len(q) > 0 {
			curr := q[0]
			if i < 1000 {
				c[curr.p]++
			}
			q = q[1:]
			if _, ok := modules[curr.dest]; !ok {
				continue
			}
			for _, o := range modules[curr.dest].signal(curr.src, curr.p) {
				q = append(q, op{src: curr.dest, p: o.p, dest: o.dest})
			}
			for _, pp := range parentsOfParents {
				// Assuming "rx" is derived from a single conjunction, all of whose parents are also conjunctions.
				if _, ok := cyclelen[pp]; !ok && modules[parents["rx"][0]].(*conjunction).inputs[pp] == high {
					cyclelen[pp] = i + 1
				}
			}
		}
	}
	fmt.Println(c[low] * c[high])
	lcm := 1
	for _, v := range maps.Values(cyclelen) {
		lcm = util.Lcm(lcm, v)
	}
	fmt.Println(lcm)
}

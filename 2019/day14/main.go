package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type reaction struct {
	qty      int
	reagents map[string]int
}

func main() {
	formulae := make(map[string]reaction)
	s := util.ScanAll()
	for s.Scan() {
		sp := strings.Split(s.Text(), " => ")
		name, qty := parse(sp[1])
		rxn := reaction{
			qty:      qty,
			reagents: make(map[string]int),
		}
		for _, r := range strings.Split(sp[0], ", ") {
			name, qty := parse(r)
			rxn.reagents[name] = qty
		}
		formulae[name] = rxn
	}

	fmt.Println(ore("FUEL", 1, formulae, map[string]int{}))

	fmt.Println(sort.Search(1_000_000_000_000, func(i int) bool {
		return ore("FUEL", i, formulae, map[string]int{}) > 1_000_000_000_000
	}) - 1)
}

func parse(s string) (string, int) {
	var name string
	var qty int
	fmt.Sscanf(s, "%d %s", &qty, &name)
	return name, qty
}

func ore(name string, qty int, formulae map[string]reaction, surplus map[string]int) int {
	if name == "ORE" {
		return qty
	}
	if surplus[name] > 0 {
		min := util.Min(qty, surplus[name])
		qty -= min
		surplus[name] -= min
	}
	if qty == 0 {
		return 0
	}
	n := qty / formulae[name].qty
	if qty%formulae[name].qty != 0 {
		n++
	}
	surplus[name] += n*formulae[name].qty - qty
	var sum int
	for r, qty := range formulae[name].reagents {
		sum += ore(r, n*qty, formulae, surplus)
	}
	return sum
}

package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input := strings.Split(util.ReadAll(), "\n\n")
	transformations, medicine := strings.Split(input[0], "\n"), input[1]
	replacements := make(map[string][]string)
	for _, t := range transformations {
		var from, to string
		fmt.Sscanf(t, "%s => %s", &from, &to)
		replacements[from] = append(replacements[from], to)
	}

	molecules := make(map[string]struct{})
	c := components(medicine)
	for i := range c {
		for _, r := range replacements[c[i]] {
			m := splice(c, i, r)
			molecules[m] = struct{}{}
		}
	}
	fmt.Println(len(molecules))

	fmt.Println(shortestPath(medicine))
}

func components(medicine string) []string {
	var ret []string
	for i := 0; i < len(medicine); i++ {
		atom := []byte{medicine[i]}
		if i+1 < len(medicine) && medicine[i+1] >= 'a' {
			atom = append(atom, medicine[i+1])
		}
		ret = append(ret, string(atom))
		if len(atom) == 2 {
			i++
		}
	}
	return ret
}

func splice(a []string, i int, r string) string {
	c := make([]string, len(a))
	copy(c, a)
	c = append(c[:i], append([]string{r}, c[i+1:]...)...)
	return strings.Join(c, "")
}

func shortestPath(medicine string) int {
	c := components(medicine)
	m := make(map[string]int)
	for i := range c {
		m[c[i]]++
	}
	return len(c) - m["Rn"] - 2*m["Y"] - m["Ar"] - 1
}

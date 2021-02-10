package main

import (
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

type ingredient [5]int

type mix struct {
	name   string
	amount int
}

func main() {
	ingredients := make(map[string]ingredient)
	s := util.ScanAll()
	for s.Scan() {
		var i ingredient
		var name string
		fmt.Sscanf(s.Text(), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &i[0], &i[1], &i[2], &i[3], &i[4])
		ingredients[name[:len(name)-1]] = i
	}
	var max1, max2 int
	for _, m := range recipes(keys(ingredients), 100) {
		s, cal := score(m, ingredients)
		if s > max1 {
			max1 = s
		}
		if cal == 500 && s > max2 {
			max2 = s
		}
	}
	fmt.Println(max1)
	fmt.Println(max2)
}

func recipes(ingredients []string, total int) [][]mix {
	if len(ingredients) == 1 {
		return [][]mix{{{ingredients[0], total}}}
	}
	var ret [][]mix
	for i := 0; i <= total; i++ {
		m := recipes(ingredients[1:], total-i)
		for j := range m {
			m[j] = append(m[j], mix{ingredients[0], i})
		}
		ret = append(ret, m...)
	}
	return ret
}

func keys(m map[string]ingredient) []string {
	var ret []string
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func score(mixes []mix, ingredients map[string]ingredient) (int, int) {
	var sum ingredient
	for _, m := range mixes {
		for j := 0; j < 5; j++ {
			sum[j] += m.amount * ingredients[m.name][j]
		}
	}
	prod := 1
	for j := 0; j < 4; j++ {
		if sum[j] <= 0 {
			return 0, sum[4]
		}
		prod *= sum[j]
	}
	return prod, sum[4]
}

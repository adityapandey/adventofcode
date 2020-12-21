// https://adventofcode.com/2020/day/21
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`(.+) \(contains (.+)\)`)

type food struct {
	ingredients []string
	allergens   []string
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	var foods []food
	for s.Scan() {
		matches := re.FindAllStringSubmatch(s.Text(), -1)
		var f food
		f.ingredients = strings.Split(matches[0][1], " ")
		sort.Strings(f.ingredients)
		f.allergens = strings.Split(matches[0][2], ", ")
		sort.Strings(f.allergens)
		foods = append(foods, f)
	}

	// Part 1
	candidates := make(map[string]map[string]struct{})
	for _, f := range foods {
		for _, i := range f.ingredients {
			if _, ok := candidates[i]; !ok {
				candidates[i] = make(map[string]struct{})
			}
			for _, a := range f.allergens {
				candidates[i][a] = struct{}{}
			}
		}
	}

	cleanIngredients := make(map[string]struct{})
	for ingredient := range candidates {
		for allergen := range candidates[ingredient] {
			// Rule out <allergen = ingredient> from foods list
			for _, f := range foods {
				if contains(f.allergens, allergen) {
					if !contains(f.ingredients, ingredient) {
						delete(candidates[ingredient], allergen)
						if len(candidates[ingredient]) == 0 {
							cleanIngredients[ingredient] = struct{}{}
						}
						break
					}
				}
			}
		}
	}

	sum := 0
	for _, f := range foods {
		for _, i := range f.ingredients {
			if _, ok := cleanIngredients[i]; ok {
				sum++
			}
		}
	}
	fmt.Println(sum)

	// Part 2
	var dangerous []struct {
		ingredient string
		candidates map[string]struct{}
	}
	for ingredient := range candidates {
		if len(candidates[ingredient]) > 0 {
			dangerous = append(dangerous,
				struct {
					ingredient string
					candidates map[string]struct{}
				}{ingredient, candidates[ingredient]})
		}
	}

	sort.Slice(dangerous, func(i, j int) bool {
		return len(dangerous[i].candidates) > len(dangerous[j].candidates)
	})
	for len(dangerous[0].candidates) > 1 {
		for i := 0; i < len(dangerous); i++ {
			if len(dangerous[i].candidates) > 1 {
				continue
			}
			// Rule out ingredients already mapped to exactly one allergen from other ingredients.
			for k := range dangerous[i].candidates {
				for j := 0; j < len(dangerous); j++ {
					if j == i {
						continue
					}
					delete(dangerous[j].candidates, k)
				}
			}
		}
		sort.Slice(dangerous, func(i, j int) bool {
			return len(dangerous[i].candidates) > len(dangerous[j].candidates)
		})
	}

	var output []string
	sort.Slice(dangerous, func(i, j int) bool {
		for ki := range dangerous[i].candidates {
			for kj := range dangerous[j].candidates {
				return ki < kj
			}
		}
		return false
	})
	for _, i := range dangerous {
		output = append(output, i.ingredient)
	}
	fmt.Println(strings.Join(output, ","))
}

func contains(a []string, x string) bool {
	q := sort.SearchStrings(a, x)
	return q < len(a) && a[q] == x
}

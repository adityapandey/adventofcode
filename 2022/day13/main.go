package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	s := util.ReadAll()
	var packets []any
	sum := 0
	for i, pair := range strings.Split(s, "\n\n") {
		sp := strings.Split(pair, "\n")
		var first, second any
		json.Unmarshal([]byte(sp[0]), &first)
		json.Unmarshal([]byte(sp[1]), &second)
		packets = append(packets, first, second)
		if compare(first, second) == -1 {
			sum += i + 1
		}
	}
	fmt.Println(sum)

	var divider1, divider2 any
	json.Unmarshal([]byte("[[2]]"), &divider1)
	json.Unmarshal([]byte("[[6]]"), &divider2)
	packets = append(packets, divider1, divider2)
	sort.Slice(packets, func(i, j int) bool { return compare(packets[i], packets[j]) < 0 })
	divider1Pos := sort.Search(len(packets), func(i int) bool { return compare(packets[i], divider1) >= 0 })
	divider2Pos := sort.Search(len(packets), func(i int) bool { return compare(packets[i], divider2) >= 0 })
	fmt.Println((divider1Pos + 1) * (divider2Pos + 1))
}

func compare(a, b any) int {
	_, anum := a.(float64)
	_, bnum := b.(float64)
	switch {
	case anum && bnum:
		return util.Sign(int(a.(float64)) - int(b.(float64)))
	case anum:
		return compare([]any{a}, b)
	case bnum:
		return compare(a, []any{b})
	default:
		aa, bb := a.([]any), b.([]any)
		for i := 0; i < len(aa) && i < len(bb); i++ {
			if c := compare(aa[i], bb[i]); c != 0 {
				return c
			}
		}
		return util.Sign(len(aa) - len(bb))
	}
}

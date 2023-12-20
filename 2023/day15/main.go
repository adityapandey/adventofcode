package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type lens struct {
	label string
	focal int
}

func main() {
	h := 0
	sum := 0
	input := util.ReadAll()
	input = strings.ReplaceAll(input, "\n", "")
	for _, s := range strings.Split(input, ",") {
		h = hash(s)
		sum += h
	}
	fmt.Println(sum)

	var box [256][]lens
	for _, s := range strings.Split(input, ",") {
		oper := strings.IndexAny(s, "=-")
		label := s[:oper]
		h := hash(label)
		switch s[oper] {
		case '-':
			if i := slices.IndexFunc(box[h], func(x lens) bool { return x.label == label }); i != -1 {
				box[h] = append(box[h][:i], box[h][i+1:]...)
			}
		case '=':
			focal := int(s[oper+1]) - int('0')
			if i := slices.IndexFunc(box[h], func(x lens) bool { return x.label == label }); i != -1 {
				box[h][i].focal = focal
			} else {
				box[h] = append(box[h], lens{label, focal})
			}
		}
	}
	sum2 := 0
	for i := range box {
		for j := range box[i] {
			sum2 += (i + 1) * (j + 1) * box[i][j].focal
		}
	}
	fmt.Println(sum2)
}

func hash(s string) int {
	h := 0
	for i := range s {
		h += int(s[i])
		h *= 17
		h %= 256
	}
	return h
}

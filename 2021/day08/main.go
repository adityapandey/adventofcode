package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type entry struct {
	inputs  []string
	outputs []string
}

var digits = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

var matches = map[int]string{
	2: "cf",
	3: "acf",
	4: "bcdf",
}

type table [7][7]bool

func makeTable() table {
	var t [7][7]bool
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			t[i][j] = true
		}
	}
	return t
}

func (t table) String() string {
	var sb strings.Builder
	for i := 0; i < 7; i++ {
		fmt.Fprintf(&sb, "%c:", 'a'+i)
		for j := 0; j < 7; j++ {
			ok := t[i][j]
			fmt.Fprintf(&sb, " %c", map[bool]byte{true: byte('a' + j), false: ' '}[ok])
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()

}

func (t *table) match(a, b string) {
	aa := map[int]struct{}{}
	bb := map[int]struct{}{}
	for i := range a {
		aa[int(a[i]-'a')] = struct{}{}
	}
	for i := range b {
		bb[int(b[i]-'a')] = struct{}{}
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			_, aok := aa[i]
			_, bok := bb[j]
			if (aok && !bok) || (!aok && bok) {
				t[i][j] = false
			}
		}
	}
}

func (t table) permutations() map[table]struct{} {
	r := map[table]struct{}{}
	dup := false
	for i := 0; i < 7; i++ {
		var trues []int
		for j := 0; j < 7; j++ {
			if t[i][j] {
				trues = append(trues, j)
			}
		}
		if len(trues) == 1 {
			continue
		}
		dup = true
		for _, j := range trues {
			tc := t
			for jj := 0; jj < 7; jj++ {
				if jj != j {
					tc[i][jj] = false
				}
			}
			for ii := 0; ii < 7; ii++ {
				if i == ii {
					continue
				}
				tc[ii][j] = false
			}
			for tt := range tc.permutations() {
				r[tt] = struct{}{}
			}
		}
	}
	if !dup {
		r[t] = struct{}{}
	}
	return r
}

func (t table) translate(s string) string {
	tr := make(map[byte]byte)
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if t[i][j] {
				tr[byte('a'+i)] = byte('a' + j)
			}
		}
	}
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = tr[s[i]]
	}
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func main() {
	s := util.ScanAll()
	var entries []entry
	for s.Scan() {
		var e entry
		sp := strings.Split(s.Text(), " | ")
		e.inputs = strings.Split(sp[0], " ")
		e.outputs = strings.Split(sp[1], " ")
		entries = append(entries, e)
	}

	sum := 0
	for _, e := range entries {
		for _, o := range e.outputs {
			switch len(o) {
			case 2, 3, 4, 7:
				sum++
			}
		}
	}
	fmt.Println(sum)

	sum = 0
	for _, e := range entries {
		t := makeTable()
		for _, p := range append(e.inputs, e.outputs...) {
			switch len(p) {
			case 2, 3, 4:
				t.match(p, matches[len(p)])
			}
		}
		for tp := range t.permutations() {
			var invalid bool
			for _, i := range append(e.inputs, e.outputs...) {
				if _, ok := digits[tp.translate(i)]; !ok {
					invalid = true
					break
				}
			}
			if !invalid {
				var num string
				for _, o := range e.outputs {
					num += fmt.Sprint(digits[tp.translate(o)])
				}
				sum += util.Atoi(num)
				break
			}
		}
	}
	fmt.Println(sum)
}

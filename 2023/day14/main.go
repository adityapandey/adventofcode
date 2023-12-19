package main

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	in := util.ReadAll()
	fmt.Println(load(north(in)))

	old, new := getRepeat(in)
	for i := 0; i < old+(1000000000-old)%(new-old); i++ {
		in = cycle(in)
	}
	fmt.Println(load(in))
}

func getRepeat(in string) (int, int) {
	var old, new int
	cache := map[string]int{}
	for {
		new++
		in = cycle(in)
		var ok bool
		if old, ok = cache[in]; ok {
			break
		} else {
			cache[in] = new
		}
	}
	return old, new
}

func load(in string) int {
	sum := 0
	rows := strings.Split(in, "\n")
	for i := range rows {
		for j := range rows[i] {
			if rows[i][j] == 'O' {
				sum += len(rows) - i
			}
		}
	}
	return sum
}

func cycle(in string) string {
	in = north(in)
	in = west(in)
	in = south(in)
	in = east(in)
	return in
}

func north(in string) string {
	rows := strings.Split(in, "\n")
	cols := make([]string, len(rows[0]))
	for i := 0; i < len(rows[0]); i++ {
		for j := 0; j < len(rows); j++ {
			cols[i] += string(rows[j][i])
		}
	}

	tilt(cols)

	rows = make([]string, len(cols[0]))
	for i := range cols[0] {
		var sb strings.Builder
		for j := range cols {
			fmt.Fprintf(&sb, "%c", cols[j][i])
		}
		rows[i] = sb.String()
	}
	return strings.Join(rows, "\n")
}

func south(in string) string {
	rows := strings.Split(in, "\n")
	cols := make([]string, len(rows[0]))
	slices.Reverse(rows)
	for i := 0; i < len(rows[0]); i++ {
		for j := 0; j < len(rows); j++ {
			cols[i] += string(rows[j][i])
		}
	}

	tilt(cols)
	for i := range cols {
		b := []byte(cols[i])
		slices.Reverse(b)
		cols[i] = string(b)
	}

	rows = make([]string, len(cols[0]))
	for i := range cols[0] {
		var sb strings.Builder
		for j := range cols {
			fmt.Fprintf(&sb, "%c", cols[j][i])
		}
		rows[i] = sb.String()
	}
	return strings.Join(rows, "\n")
}

func west(in string) string {
	rows := strings.Split(in, "\n")
	tilt(rows)
	return strings.Join(rows, "\n")
}

func east(in string) string {
	rows := strings.Split(in, "\n")
	for i := range rows {
		b := []byte(rows[i])
		slices.Reverse(b)
		rows[i] = string(b)
	}
	tilt(rows)
	for i := range rows {
		b := []byte(rows[i])
		slices.Reverse(b)
		rows[i] = string(b)
	}
	return strings.Join(rows, "\n")
}

func tilt(a []string) {
	for i := range a {
		b := []byte(a[i])
		c := 0
		for j := bytes.IndexByte(b, 'O'); j != -1; j = bytes.IndexByte(b[c:], 'O') {
			k := bytes.LastIndexAny(b[:c+j], "#O")
			b[c+j] = '.'
			b[k+1] = 'O'
			c = k + 2
		}
		a[i] = string(b)
	}
}

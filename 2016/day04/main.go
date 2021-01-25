package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`([a-z-]+)([0-9]+)\[([a-z]+)\]`)

type room struct {
	name   string
	sector int
	cksum  string
}

func main() {
	var sum, secret int
	s := util.ScanAll()
	for s.Scan() {
		r := parseRoom(s.Text())
		if isReal(r) {
			sum += r.sector
			if strings.Contains(rotate(r.name, r.sector), "north") {
				secret = r.sector
			}
		}
	}
	fmt.Println(sum)
	fmt.Println(secret)
}

func parseRoom(s string) room {
	matches := re.FindAllStringSubmatch(s, -1)[0]
	return room{matches[1], util.Atoi(matches[2]), matches[3]}
}

func isReal(r room) bool {
	m := make(map[byte]int)
	for i := range r.name {
		m[r.name[i]]++
	}
	delete(m, '-')
	var a []struct {
		b byte
		f int
	}
	for b, f := range m {
		a = append(a, struct {
			b byte
			f int
		}{b, f})
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].f == a[j].f {
			return a[i].b < a[j].b
		}
		return a[i].f > a[j].f
	})
	cksum := make([]byte, 5)
	for i := 0; i < 5; i++ {
		cksum[i] = a[i].b
	}
	return string(cksum) == r.cksum
}

func rotate(s string, n int) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		var next byte
		if s[i] == '-' {
			next = ' '
		} else {
			next = 'a' + byte((int(s[i])-'a'+n)%('z'-'a'+1))
		}
		if err := sb.WriteByte(next); err != nil {
			log.Fatal(err)
		}
	}
	return sb.String()
}

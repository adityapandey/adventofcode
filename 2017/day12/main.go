package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"os"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode2017-go/util"
)

type set map[int]struct{}

func (s set) hash() int {
	a := make([]int, 0, len(s))
	for k := range s {
		a = append(a, k)
	}
	sort.Ints(a)
	return int(crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v", a))))
}

func main() {
	m := make(map[int]set)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		matches := strings.Split(s.Text(), " <-> ")
		from := util.Atoi(matches[0])
		m[from] = map[int]struct{}{from: {}}
		tos := strings.Split(matches[1], ", ")
		for i := range tos {
			m[from][util.Atoi(tos[i])] = struct{}{}
		}
	}
	fmt.Println(len(reachable(m, 0, set{})))

	groups := make(set)
	for k := range m {
		groups[reachable(m, k, set{}).hash()] = struct{}{}
	}
	fmt.Println(len(groups))

}

func reachable(m map[int]set, from int, seen set) set {
	seen[from] = struct{}{}
	tos := make(set)
	for k := range m[from] {
		tos[k] = struct{}{}
		if _, ok := seen[k]; !ok {
			for k2 := range reachable(m, k, seen) {
				tos[k2] = struct{}{}
			}
		}
	}
	return tos
}

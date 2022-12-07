package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	root := []string{""}
	dirs := map[string]int{}
	files := map[string]int{}
	var curr []string
	s := util.ScanAll()
	for s.Scan() {
		txt := strings.Fields(s.Text())
		if txt[0] == "$" {
			if txt[1] == "cd" {
				if txt[2] == "/" {
					curr = root
				} else if txt[2] == ".." {
					curr = curr[:len(curr)-1]
				} else {
					curr = append(curr, txt[2])
				}
				dirs[strings.Join(curr, "/")] = 0
			}
		} else {
			if txt[0] != "dir" {
				files[strings.Join(append(curr, txt[1]), "/")] = util.Atoi(txt[0])
			}
		}
	}

	for f, s := range files {
		path := strings.Split(f, "/")
		for i := 1; i < len(path); i++ {
			dirs[strings.Join(path[:i], "/")] += s
		}
	}

	var sortedSizes []int
	sum := 0
	for _, s := range dirs {
		sortedSizes = append(sortedSizes, s)
		if s <= 100000 {
			sum += s
		}
	}
	fmt.Println(sum)

	sort.Ints(sortedSizes)
	total, want := 70000000, 30000000
	available := total - dirs[""]
	fmt.Println(sortedSizes[sort.SearchInts(sortedSizes, want-available)])
}

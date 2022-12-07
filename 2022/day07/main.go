package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type dir struct {
	path    string
	parent  *dir
	subdirs map[string]*dir
	files   []int
}

func main() {
	root := &dir{
		path:    "",
		parent:  nil,
		subdirs: map[string]*dir{},
		files:   []int{},
	}
	var curr *dir
	s := util.ScanAll()
	for s.Scan() {
		txt := strings.Fields(s.Text())
		if txt[0] == "$" {
			if txt[1] == "cd" {
				if txt[2] == "/" {
					curr = root
				} else if txt[2] == ".." {
					curr = curr.parent
				} else {
					curr = curr.subdirs[txt[2]]
				}
			}
		} else {
			if txt[0] == "dir" {
				dirname := txt[1]
				d := &dir{
					path:    curr.path + "/" + dirname,
					parent:  curr,
					subdirs: map[string]*dir{},
					files:   []int{},
				}
				curr.subdirs[dirname] = d
			} else {
				curr.files = append(curr.files, util.Atoi(txt[0]))
			}
		}
	}

	dirsizes := map[string]int{}
	computeSize(root, dirsizes)

	var sortedSizes []int
	sum := 0
	for _, s := range dirsizes {
		sortedSizes = append(sortedSizes, s)
		if s <= 100000 {
			sum += s
		}
	}
	fmt.Println(sum)

	sort.Ints(sortedSizes)
	total, want := 70000000, 30000000
	available := total - dirsizes[root.path]
	fmt.Println(sortedSizes[sort.SearchInts(sortedSizes, want-available)])
}

func computeSize(d *dir, dirsizes map[string]int) int {
	if s, ok := dirsizes[d.path]; ok {
		return s
	}
	s := 0
	for _, f := range d.files {
		s += f
	}
	for _, sd := range d.subdirs {
		s += computeSize(sd, dirsizes)
	}
	dirsizes[d.path] = s
	return s
}

package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type file struct {
	name string
	size int
}

type dir struct {
	path    string
	parent  *dir
	subdirs map[string]*dir
	files   []file
}

func main() {
	root := &dir{
		path:    "",
		parent:  nil,
		subdirs: map[string]*dir{},
		files:   []file{},
	}
	var curr *dir
	s := util.ScanAll()
	for s.Scan() {
		f := strings.Fields(s.Text())
		if f[0] == "$" {
			if f[1] == "cd" {
				if f[2] == "/" {
					curr = root
				} else if f[2] == ".." {
					curr = curr.parent
				} else {
					curr = curr.subdirs[f[2]]
				}
			}
		} else {
			if f[0] == "dir" {
				d := &dir{
					path:    curr.path + "/" + f[1],
					parent:  curr,
					subdirs: map[string]*dir{},
					files:   []file{},
				}
				curr.subdirs[f[1]] = d
			} else {
				curr.files = append(curr.files, file{
					name: f[1],
					size: util.Atoi(f[0]),
				})
			}
		}
	}

	dirs := map[string]int{}
	computeSize(root, dirs)

	var dirsizes []int
	sum := 0
	for _, s := range dirs {
		dirsizes = append(dirsizes, s)
		if s <= 100000 {
			sum += s
		}
	}
	fmt.Println(sum)

	total, want := 70000000, 30000000
	available := total - dirs[root.path]
	sort.Ints(dirsizes)
	fmt.Println(dirsizes[sort.SearchInts(dirsizes, want-available)])
}

func computeSize(d *dir, dirs map[string]int) int {
	if s, ok := dirs[d.path]; ok {
		return s
	}
	s := 0
	for _, f := range d.files {
		s += f.size
	}
	for _, sd := range d.subdirs {
		s += computeSize(sd, dirs)
	}
	dirs[d.path] = s
	return s
}

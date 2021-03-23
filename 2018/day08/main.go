package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type node struct {
	children []*node
	metadata []int
}

func (n *node) value() int {
	sum := 0
	if len(n.children) == 0 {
		for _, m := range n.metadata {
			sum += m
		}
		return sum
	}
	for _, m := range n.metadata {
		if m > 0 && m <= len(n.children) {
			sum += n.children[m-1].value()
		}
	}
	return sum
}

func makeNode(vals []int, pos int) (*node, int) {
	n := &node{}
	numChildren := vals[pos]
	numMetadata := vals[pos+1]
	pos += 2
	for i := 0; i < numChildren; i++ {
		child, childPos := makeNode(vals, pos)
		pos = childPos
		n.children = append(n.children, child)
	}
	for i := 0; i < numMetadata; i++ {
		n.metadata = append(n.metadata, vals[pos])
		pos++
	}
	return n, pos
}

func metadataSum(root *node) int {
	sum := 0
	for _, m := range root.metadata {
		sum += m
	}
	for _, c := range root.children {
		sum += metadataSum(c)
	}
	return sum
}

func main() {
	var vals []int
	for _, f := range strings.Fields(util.ReadAll()) {
		vals = append(vals, util.Atoi(f))
	}
	root, _ := makeNode(vals, 0)
	fmt.Println(metadataSum(root))
	fmt.Println(root.value())
}

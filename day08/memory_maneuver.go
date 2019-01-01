package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	children []*Node
	metadata []int
}

func (n *Node) Value() int {
	sum := 0
	if len(n.children) == 0 {
		for _, m := range n.metadata {
			sum += m
		}
		return sum
	}
	for _, m := range n.metadata {
		if m > 0 && m <= len(n.children) {
			sum += n.children[m-1].Value()
		}
	}
	return sum
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func makeNode(vals []int, pos int) (*Node, int) {
	n := &Node{}
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

func metadataSum(root *Node) int {
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
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		for _, f := range fields {
			vals = append(vals, atoi(f))
		}
	}

	root, _ := makeNode(vals, 0)
	// Part 1
	fmt.Println(metadataSum(root))
	// Part 2
	fmt.Println(root.Value())
}

package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type node struct {
	left   *node
	right  *node
	parent *node
	val    int
	level  int
}

func (n *node) isVal() bool {
	return n.left == nil && n.right == nil
}

func parse(s string) *node {
	n := &node{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			n.left = &node{parent: n, level: n.level + 1}
			n = n.left
		case ',':
			n = n.parent
			n.right = &node{parent: n, level: n.level + 1}
			n = n.right
		case ']':
			n = n.parent
		default:
			n.val = int(s[i] - '0')
		}
	}
	return n
}

func (n *node) add(k *node) *node {
	r := &node{left: n, right: k}
	n.parent = r
	k.parent = r
	n.forEach(func(x *node) { x.level++ })
	k.forEach(func(x *node) { x.level++ })
	r.reduce()
	return r
}

func (n *node) reduce() {
	for {
		if explodeNode := n.leftMost(func(x *node) bool { return x.level >= 4 && !x.isVal() && x.left.isVal() && x.right.isVal() }); explodeNode != nil {
			explodeNode.explode()
			continue
		}
		if splitNode := n.leftMost(func(x *node) bool { return x.isVal() && x.val >= 10 }); splitNode != nil {
			splitNode.split()
			continue
		}
		break
	}
}

func (n *node) explode() {
	if l := n.nextLeftVal(); l != nil {
		l.val += n.left.val
	}
	if r := n.nextRightVal(); r != nil {
		r.val += n.right.val
	}
	n.left = nil
	n.right = nil
	n.val = 0
}

func (n *node) nextLeftVal() *node {
	for n.parent != nil && n.parent.left == n {
		n = n.parent
	}
	if n.parent == nil {
		return nil
	}
	n = n.parent.left
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *node) nextRightVal() *node {
	for n.parent != nil && n.parent.right == n {
		n = n.parent
	}
	if n.parent == nil {
		return nil
	}
	n = n.parent.right
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *node) split() {
	n.left = &node{parent: n, val: n.val / 2, level: n.level + 1}
	n.right = &node{parent: n, val: (n.val + 1) / 2, level: n.level + 1}
}

func (n *node) forEach(fn func(*node)) {
	if n == nil {
		return
	}
	fn(n)
	n.left.forEach(fn)
	n.right.forEach(fn)
}

func (n *node) leftMost(fn func(*node) bool) *node {
	if n.left != nil {
		if l := n.left.leftMost(fn); l != nil {
			return l
		}
	}
	if fn(n) {
		return n
	}
	if n.right != nil {
		return n.right.leftMost(fn)
	}
	return nil
}

func (n *node) magnitude() int {
	if n.isVal() {
		return n.val
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func main() {
	nums := strings.Split(util.ReadAll(), "\n")
	n := parse(nums[0])
	for i := 1; i < len(nums); i++ {
		k := parse(nums[i])
		n = n.add(k)
	}
	fmt.Println(n.magnitude())

	max := 0
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			n := parse(nums[i])
			k := parse(nums[j])
			n = n.add(k)
			max = util.Max(max, n.magnitude())
		}
	}
	fmt.Println(max)
}

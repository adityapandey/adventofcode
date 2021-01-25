package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`(\w+) \((\d+)\)( -> (.+))*`)

type node struct {
	name     string
	weight   int
	balance  int
	children []*node
	parent   *node
}

func main() {
	nodes := make(map[string]*node)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		matches := re.FindAllStringSubmatch(s.Text(), -1)[0]
		n := findOrCreate(nodes, matches[1])
		n.weight = util.Atoi(matches[2])
		if len(matches[4]) > 0 {
			for _, c := range strings.Split(matches[4], ", ") {
				child := findOrCreate(nodes, c)
				n.children = append(n.children, child)
				child.parent = n
				nodes[child.name] = child
			}
		}
		nodes[n.name] = n
	}
	var root *node
	for _, n := range nodes {
		if n.parent == nil {
			root = n
			break
		}
	}
	fmt.Println(root.name)
	computeBalance(root)
	b, err := findUnbalanced(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

func findOrCreate(nodes map[string]*node, name string) *node {
	if n, ok := nodes[name]; ok {
		return n
	}
	return &node{name: name}
}

func computeBalance(n *node) {
	if n.balance != 0 {
		return
	}
	if len(n.children) == 0 {
		n.balance = n.weight
		return
	}
	sum := n.weight
	for _, c := range n.children {
		computeBalance(c)
		sum += c.balance
	}
	n.balance = sum
}

func findUnbalanced(n *node) (int, error) {
	m := make(map[int][]*node)
	for _, c := range n.children {
		m[c.balance] = append(m[c.balance], c)
	}
	if len(m) != 2 {
		return 0, fmt.Errorf("More than one child with unexpected balance: %v", m)
	}
	var expectedBalance int
	var unbalancedNode *node
	for b := range m {
		if len(m[b]) == 1 {
			unbalancedNode = m[b][0]
		} else {
			expectedBalance = b
		}
	}
	return findCulprit(unbalancedNode, expectedBalance)
}

func findCulprit(n *node, expectedBalance int) (int, error) {
	fmt.Println(n.name, expectedBalance)
	m := make(map[int][]*node)
	for _, c := range n.children {
		m[c.balance] = append(m[c.balance], c)
	}
	fmt.Println(m)
	switch len(m) {
	case 0:
		// leaf node
		return expectedBalance, nil
	case 1:
		// all children have equal balances, so this is the culprit
		return expectedBalance - len(n.children)*n.children[0].balance, nil
	case 2:
		expectedBalance -= n.weight
		expectedBalance /= len(n.children)
		for b := range m {
			if b == expectedBalance {
				continue
			}
			if len(m[b]) != 1 {
				return 0, fmt.Errorf("More than one child with unexpected balance: %v", m)
			}
			return findCulprit(m[b][0], expectedBalance)
		}
		return 0, fmt.Errorf("Unreachable statement: %v", m)
	default:
		return 0, fmt.Errorf("No possible consensus balance: %v", m)
	}
}

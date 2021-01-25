// https://adventofcode.com/2020/day/18
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	exprs := bytes.Split(input, []byte("\n"))

	// Part 1
	sum := 0
	for _, e := range exprs {
		sum += eval(e, map[byte]int{'*': 1, '+': 1})
	}
	fmt.Println(sum)

	// Part 2
	sum = 0
	for _, e := range exprs {
		sum += eval(e, map[byte]int{'*': 1, '+': 2})
	}
	fmt.Println(sum)
}

func eval(expr []byte, precedence map[byte]int) int {
	var opStack []byte
	var postfix []byte
	for _, token := range expr {
		switch token {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			postfix = append(postfix, token)
		case '(':
			opStack = append(opStack, token)
		case ')':
			for len(opStack) > 0 {
				var op byte
				op, opStack = opStack[len(opStack)-1], opStack[:len(opStack)-1]
				if op == '(' {
					break
				}
				postfix = append(postfix, op)
			}
		case '+', '*':
			for len(opStack) > 0 {
				op := opStack[len(opStack)-1]
				if op == '(' || precedence[token] > precedence[op] {
					break
				}
				opStack = opStack[:len(opStack)-1]
				postfix = append(postfix, op)
			}
			opStack = append(opStack, token)
		}
	}
	for len(opStack) > 0 {
		postfix = append(postfix, opStack[len(opStack)-1])
		opStack = opStack[:len(opStack)-1]
	}
	return evalPostfix(postfix)
}

func evalPostfix(postfix []byte) int {
	var stack []int
	for _, token := range postfix {
		switch token {
		case '+':
			l := len(stack)
			stack[l-2] += stack[l-1]
			stack = stack[:l-1]
		case '*':
			l := len(stack)
			stack[l-2] *= stack[l-1]
			stack = stack[:l-1]
		default:
			stack = append(stack, ord(token))
		}
	}
	return stack[0]
}

func ord(b byte) int {
	return int(b) - '0'
}

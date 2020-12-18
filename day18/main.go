// https://adventofcode.com/2020/day/18
package main

import (
	"bytes"
	"errors"
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
		sum += eval(e, [][]byte{{'+', '*'}})
	}
	fmt.Println(sum)

	// Part 2
	sum = 0
	for _, e := range exprs {
		sum += eval(e, [][]byte{{'+'}, {'*'}})
	}
	fmt.Println(sum)
}

func eval(expr []byte, precedence [][]byte) int {
	var ops []byte
	var vals []int
	for i := 0; i < len(expr); i++ {
		switch expr[i] {
		case '+', '*':
			ops = append(ops, expr[i])
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			vals = append(vals, ord(expr[i]))
		case '(':
			end, err := findCloseParen(expr, i)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, eval(expr[i+1:end], precedence))
			i = end
		}
	}
	for _, currOps := range precedence {
		for i := 0; i < len(ops); i++ {
			if !bytes.Contains(currOps, []byte{ops[i]}) {
				continue
			}
			switch ops[i] {
			case '+':
				vals[i+1] = vals[i] + vals[i+1]
			case '*':
				vals[i+1] = vals[i] * vals[i+1]
			}
			vals = append(vals[:i], vals[i+1:]...)
			ops = append(ops[:i], ops[i+1:]...)
			i--
		}
	}
	return vals[0]
}

func ord(b byte) int {
	return int(b) - '0'
}

func findCloseParen(expr []byte, i int) (int, error) {
	var parens int
	for ; i < len(expr); i++ {
		switch expr[i] {
		case '(':
			parens++
		case ')':
			parens--
			if parens == 0 {
				return i, nil
			}
		}
	}
	return -1, errors.New("No matching closing parenthesis")
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type group struct {
	start, end, score int
}

type state int

const (
	NONE state = iota
	GROUP
	GARBAGE
)

func main() {
	var groups []group
	var stack []group
	garbage := 0
	st := NONE
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(input); i++ {
		c := input[i]
		switch st {
		case NONE:
			switch c {
			case '{':
				st = GROUP
				stack = append(stack, group{start: i, score: 1})
			case '<':
				st = GARBAGE
			default:
				log.Fatal(c)
			}
		case GROUP:
			switch c {
			case '{':
				stack = append(stack, group{start: i, score: stack[len(stack)-1].score + 1})
			case '}':
				stack[len(stack)-1].end = i
				groups = append(groups, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			case '<':
				st = GARBAGE
			}
		case GARBAGE:
			switch c {
			case '!':
				i++
			case '>':
				st = GROUP
			default:
				garbage++
			}
		}
	}
	sum := 0
	for _, g := range groups {
		sum += g.score
	}
	fmt.Println(sum)
	fmt.Println(garbage)
}

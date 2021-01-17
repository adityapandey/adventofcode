package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type state int

const (
	none state = iota
	group
	garbage
)

func main() {
	var sumScore, currScore, garbageCount int
	st := none
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(input); i++ {
		c := input[i]
		switch st {
		case none:
			switch c {
			case '{':
				st = group
				currScore = 1
			case '<':
				st = garbage
			default:
				log.Fatal(c)
			}
		case group:
			switch c {
			case '{':
				currScore++
			case '}':
				sumScore += currScore
				currScore--
			case '<':
				st = garbage
			}
		case garbage:
			switch c {
			case '!':
				i++
			case '>':
				st = group
			default:
				garbageCount++
			}
		}
	}
	fmt.Println(sumScore)
	fmt.Println(garbageCount)
}

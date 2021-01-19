package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode2017-go/util"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var lengths []int
	for _, s := range strings.Split(string(input), ",") {
		lengths = append(lengths, util.Atoi(s))
	}

	// Part 1
	listLen := 256
	knot := util.NewKnot(listLen)
	for _, l := range lengths {
		knot.Hash(l)
	}
	fmt.Println(knot.List[0] * knot.List[1])

	// Part 2
	knot = util.NewKnot(listLen)
	fmt.Println(knot.DenseHash(input))
}

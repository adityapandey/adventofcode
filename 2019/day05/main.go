package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadAll(), ",") {
		program = append(program, util.Atoi(n))
	}
	m := machine.New(program, []int{1})
	m.Run()
	out := m.Output()
	for _, z := range out[:len(out)-1] {
		if z != 0 {
			log.Fatal("Expected zero")
		}
	}
	fmt.Println(out[len(out)-1])
	m = machine.New(program, []int{5})
	m.Run()
	fmt.Println(m.Output()[0])
}

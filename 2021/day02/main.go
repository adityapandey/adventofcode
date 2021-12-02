package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	var horiz, depth1, aim, depth2 int
	for _, line := range strings.Split(util.ReadAll(), "\n") {
		var cmd string
		var x int
		fmt.Sscanf(line, "%s %d", &cmd, &x)
		switch cmd {
		case "forward":
			horiz += x
			depth2 += aim * x
		case "down":
			depth1 += x
			aim += x
		case "up":
			depth1 -= x
			aim -= x
		}
	}
	fmt.Println(horiz * depth1)
	fmt.Println(horiz * depth2)
}

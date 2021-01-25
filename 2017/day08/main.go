package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`(\w+) (inc|dec) ([0-9-]+) if (\w+) (>|<|<=|>=|==|!=) ([0-9-]+)`)

func main() {
	max := math.MinInt32
	registers := make(map[string]int)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		matches := re.FindAllStringSubmatch(s.Text(), -1)[0]
		cond, err := eval(registers[matches[4]], matches[5], util.Atoi(matches[6]))
		if err != nil {
			log.Fatal(err)
		}
		if cond {
			switch matches[2] {
			case "inc":
				registers[matches[1]] += util.Atoi(matches[3])
			case "dec":
				registers[matches[1]] -= util.Atoi(matches[3])
			default:
				log.Fatal("Umkmowm verb", matches[2])
			}
		}
		m := maxRegister(registers)
		if m > max {
			max = m
		}
	}

	fmt.Println(maxRegister(registers))
	fmt.Println(max)
}

func eval(register int, condition string, value int) (bool, error) {
	switch condition {
	case ">":
		return register > value, nil
	case "<":
		return register < value, nil
	case ">=":
		return register >= value, nil
	case "<=":
		return register <= value, nil
	case "==":
		return register == value, nil
	case "!=":
		return register != value, nil
	default:
		return false, fmt.Errorf("Unknown operator %v", condition)
	}
}

func maxRegister(regs map[string]int) int {
	max := math.MinInt32
	for _, v := range regs {
		if v > max {
			max = v
		}
	}
	return max
}

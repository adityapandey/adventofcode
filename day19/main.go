// https://adventofcode.com/2020/day/19
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/adityapandey/adventofcode2020-go/util"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	inputSplit := strings.Split(string(input), "\n\n")

	// Part 1
	inputRules := make(map[int]string)
	regexRules := make(map[int]string)
	for _, rule := range strings.Split(inputSplit[0], "\n") {
		split := strings.Split(rule, ": ")
		i := util.Atoi(split[0])
		inputRules[i] = split[1]
		if quote := strings.Index(rule, `"`); quote != -1 {
			regexRules[i] = string(rule[quote+1])
		}
	}
	re0 := regexp.MustCompile(fmt.Sprintf("^%s$", expand(0, inputRules, regexRules)))

	sum := 0
	for _, msg := range strings.Split(inputSplit[1], "\n") {
		if re0.MatchString(msg) {
			sum++
		}
	}
	fmt.Println(sum)

	// Part 2
	//
	// 0: 8 11
	// 8: 42 | 42  8
	// 11: 42 31 | 42 11 31
	//
	// minimal 0: 42 42 31
	//
	// 8: (42)+, 11: (42){n}(31){n}
	// Thus 0: (42){m}(31){n} m>n, m>=2, n>=1
	re42 := regexp.MustCompile(fmt.Sprintf("^%s", regexRules[42]))
	re31 := regexp.MustCompile(fmt.Sprintf("^%s", regexRules[31]))
	sum = 0
	for _, msg := range strings.Split(inputSplit[1], "\n") {
		var m, n int
		for len(msg) > 0 && re42.MatchString(msg) {
			msg = msg[re42.FindStringIndex(msg)[1]:]
			m++
		}
		for len(msg) > 0 && re31.MatchString(msg) {
			msg = msg[re31.FindStringIndex(msg)[1]:]
			n++
		}
		if len(msg) == 0 && m >= 2 && n >= 1 && m > n {
			sum++
		}
	}
	fmt.Println(sum)

}

func expand(i int, inputRules, regexRules map[int]string) string {
	if v, ok := regexRules[i]; ok {
		return v
	}
	var regexSubRules []string
	for _, subRule := range strings.Split(inputRules[i], " | ") {
		var sb strings.Builder
		for _, num := range strings.Split(subRule, " ") {
			sb.WriteString(expand(util.Atoi(num), inputRules, regexRules))
		}
		regexSubRules = append(regexSubRules, sb.String())
	}
	regexRules[i] = fmt.Sprintf("(%s)", strings.Join(regexSubRules, "|"))
	return regexRules[i]
}

// https://adventofcode.com/2020/day/4
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var kvRegexp = regexp.MustCompile(`[^ ]+:[^ ]+`)
var hgtRegexp = regexp.MustCompile(`^(\d+)(cm|in)$`)
var hclRegexp = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var pidRegexp = regexp.MustCompile(`^\d{9}$`)
var eclRegexp = regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)

type passport struct {
	kv map[string]string
}

func main() {
	var passports []passport
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range strings.Split(string(input), "\n\n") {
		p := passport{make(map[string]string)}
		for _, line := range strings.Split(t, "\n") {
			for _, kv := range kvRegexp.FindAllString(line, -1) {
				kvSplit := strings.Split(kv, ":")
				p.kv[kvSplit[0]] = kvSplit[1]
			}
		}
		passports = append(passports, p)
	}

	// Part 1
	valid := 0
	for _, p := range passports {
		if isValid(p) {
			valid++
		}
	}
	fmt.Println(valid)

	// Part 2
	valid = 0
	for _, p := range passports {
		if strictValid(p) {
			valid++
		}
	}
	fmt.Println(valid)
}

func isValid(p passport) bool {
	delete(p.kv, "cid")
	return len(p.kv) == 7
}

func strictValid(p passport) bool {
	return isValid(p) &&
		between(p.kv["byr"], 1920, 2002) &&
		between(p.kv["iyr"], 2010, 2020) &&
		between(p.kv["eyr"], 2020, 2030) &&
		isValidHgt(p.kv["hgt"]) &&
		hclRegexp.MatchString(p.kv["hcl"]) &&
		eclRegexp.MatchString(p.kv["ecl"]) &&
		pidRegexp.MatchString(p.kv["pid"])
}

func between(s string, min, max int) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}

func isValidHgt(s string) bool {
	res := hgtRegexp.FindAllStringSubmatch(s, -1)
	if len(res) != 1 {
		return false
	}
	r := res[0]
	switch r[2] {
	case "cm":
		return between(r[1], 150, 193)
	case "in":
		return between(r[1], 59, 76)
	default:
		return false
	}
}

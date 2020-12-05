// https://adventofcode.com/2020/day/4
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var kvRegexp = regexp.MustCompile("[^ ]+:[^ ]+")
var hgtRegexp = regexp.MustCompile("^(\\d+)(cm|in)$")
var hclRegexp = regexp.MustCompile("^#[0-9a-f]{6}$")
var pidRegexp = regexp.MustCompile("^\\d{9}$")
var eclRegexp = regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")

type passport struct {
	kv map[string]string
}

func isValid(p passport) bool {
	delete(p.kv, "cid")
	return len(p.kv) == 7
}

func strictValid(p passport) bool {
	if !isValid(p) {
		return false
	}
	if !between(p.kv["byr"], 1920, 2002) {
		return false
	}
	if !between(p.kv["iyr"], 2010, 2020) {
		return false
	}
	if !between(p.kv["eyr"], 2020, 2030) {
		return false
	}
	if !validHgt(p.kv["hgt"]) {
		return false
	}
	if !hclRegexp.MatchString(p.kv["hcl"]) {
		return false
	}
	if !eclRegexp.MatchString(p.kv["ecl"]) {
		return false
	}
	if !pidRegexp.MatchString(p.kv["pid"]) {
		return false
	}
	return true
}

func between(s string, min, max int) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}

func validHgt(s string) bool {
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

func main() {
	var passports []passport
	s := bufio.NewScanner(os.Stdin)
	p := &passport{make(map[string]string)}
	for s.Scan() {
		if s.Text() == "" {
			passports = append(passports, *p)
			p = &passport{make(map[string]string)}
		}
		for _, kv := range kvRegexp.FindAllString(s.Text(), -1) {
			split := strings.Split(kv, ":")
			p.kv[split[0]] = split[1]
		}
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

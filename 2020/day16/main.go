// https://adventofcode.com/2020/day/16
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type rule struct {
	min, max int
}

type ticket []int

type validColsEntry struct {
	name string
	cols map[int]struct{}
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(input), "\n\n")
	mappingsString, yourTicket, nearby := split[0], split[1], split[2]

	// Part 1
	mappings := make(map[string][]rule)
	var allRules []rule
	for _, mapping := range strings.Split(mappingsString, "\n") {
		mappingSplit := strings.Split(mapping, ": ")
		rulesSplit := strings.Split(mappingSplit[1], " or ")
		rules := make([]rule, len(rulesSplit))
		for i := range rulesSplit {
			fmt.Sscanf(rulesSplit[i], "%d-%d", &rules[i].min, &rules[i].max)
			allRules = append(allRules, rules[i])
		}
		mappings[mappingSplit[0]] = rules
	}

	ticketLines := strings.Split(nearby, "\n")[1:]
	tickets := make([]ticket, len(ticketLines))
	for i, line := range ticketLines {
		tickets[i] = parseTicket(line)
	}

	var invalid []int
	for i := 0; i < len(tickets); i++ {
		hasInvalid := false
		for _, field := range tickets[i] {
			if !anyRule(field, allRules) {
				invalid = append(invalid, field)
				hasInvalid = true
			}
		}
		if hasInvalid {
			tickets = append(tickets[:i], tickets[i+1:]...)
			i--
		}
	}

	sum := 0
	for _, n := range invalid {
		sum += n
	}
	fmt.Println(sum)

	// Part 2
	yours := parseTicket(strings.Split(yourTicket, "\n")[1])
	tickets = append(tickets, yours)

	var entries []validColsEntry
	for name, rules := range mappings {
		cols := make(map[int]struct{})
		for i := 0; i < len(mappings); i++ {
			valid := true
			for _, t := range tickets {
				if !anyRule(t[i], rules) {
					valid = false
				}
			}
			if valid {
				cols[i] = struct{}{}
			}
		}
		entries = append(entries, validColsEntry{name, cols})
	}

	sort.Slice(entries, func(i, j int) bool { return len(entries[i].cols) < len(entries[j].cols) })
	for len(entries[len(entries)-1].cols) > 1 {
		if len(entries[0].cols) == 0 {
			log.Fatal("No solution: ", entries)
		}
		if len(entries[0].cols) > 1 {
			log.Fatal("Multiple solution: ", entries)
		}
		for i := 0; i < len(entries)-1; i++ {
			if len(entries[i].cols) > 1 {
				break
			}
			for j := i + 1; j < len(entries); j++ {
				for k := range entries[i].cols {
					delete(entries[j].cols, k)
				}
			}
		}
		sort.Slice(entries, func(i, j int) bool { return len(entries[i].cols) < len(entries[j].cols) })
	}

	prod := 1
	for i := range entries {
		if strings.HasPrefix(entries[i].name, "departure") {
			for k := range entries[i].cols {
				prod *= yours[k]
			}
		}
	}
	fmt.Println(prod)
}

func parseTicket(line string) ticket {
	var t ticket
	for _, field := range strings.Split(line, ",") {
		t = append(t, util.Atoi(field))
	}
	return t
}

func anyRule(n int, rules []rule) bool {
	for _, r := range rules {
		if n >= r.min && n <= r.max {
			return true
		}
	}
	return false
}

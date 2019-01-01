package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Roster struct {
	t     time.Time
	event string
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	var roster []Roster
	rosterRe := regexp.MustCompile(`^\[(.*)\] (.*)$`)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fields := rosterRe.FindAllStringSubmatch(s.Text(), -1)[0]
		t, err := time.Parse("2006-01-02 15:04", fields[1])
		if err != nil {
			log.Fatal("Unable to parse ", fields[1], ": ", err)
		}
		roster = append(roster, Roster{t, fields[2]})
	}
	sort.Slice(roster, func(i, j int) bool { return roster[i].t.Before(roster[j].t) })

	// (guardId, date) -> minutes slept
	m := make(map[int]map[string][]int)
	var date string
	var guardId, startSleep int
	for _, r := range roster {
		switch {
		case strings.HasPrefix(r.event, "Guard"):
			fmt.Sscanf(r.event, "Guard #%d begins shift", &guardId)
			if r.t.Hour() == 0 {
				date = r.t.Format("01-02")
			} else {
				date = r.t.Add(time.Hour).Format("01-02")
			}
			if _, ok := m[guardId]; !ok {
				m[guardId] = make(map[string][]int)
			}
			m[guardId][date] = make([]int, 60, 60)
		case r.event == "falls asleep":
			startSleep = r.t.Minute()
		case r.event == "wakes up":
			endSleep := r.t.Minute()
			for i := startSleep; i < endSleep; i++ {
				m[guardId][date][i] = 1
			}
		}
	}

	// Part 1
	var maxSlept int
	for guard, v := range m {
		var slept int
		for _, mins := range v {
			for i := range mins {
				slept += mins[i]
			}
		}
		if slept > maxSlept {
			maxSlept, guardId = slept, guard
		}
	}

	sleepsPerMinute := make([]int, 60, 60)
	for _, mins := range m[guardId] {
		for i := range mins {
			sleepsPerMinute[i] += mins[i]
		}
	}

	var maxSleepsPerMinute, minute int
	for i := range sleepsPerMinute {
		if sleepsPerMinute[i] > maxSleepsPerMinute {
			maxSleepsPerMinute, minute = sleepsPerMinute[i], i
		}
	}
	fmt.Println(guardId * minute)

	// Part 2
	maxSleepsPerMinute, minute = 0, 0
	for guard, v := range m {
		sleepsPerMinute = make([]int, 60, 60)
		for _, mins := range v {
			for i := range mins {
				sleepsPerMinute[i] += mins[i]
			}
		}
		for i, sleep := range sleepsPerMinute {
			if sleep > maxSleepsPerMinute {
				maxSleepsPerMinute, guardId, minute = sleep, guard, i
			}
		}
	}

	fmt.Println(guardId * minute)
}

package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/adityapandey/adventofcode/util"
)

type event struct {
	t   time.Time
	msg string
}

var eventRe = regexp.MustCompile(`^\[(.*)\] (.*)$`)

func main() {
	var events []event
	s := util.ScanAll()
	for s.Scan() {
		fields := eventRe.FindAllStringSubmatch(s.Text(), -1)[0]
		t, err := time.Parse("2006-01-02 15:04", fields[1])
		if err != nil {
			log.Fatal("Unable to parse ", fields[1], ": ", err)
		}
		events = append(events, event{t, fields[2]})
	}
	sort.Slice(events, func(i, j int) bool { return events[i].t.Before(events[j].t) })

	// (guardId, date) -> minutes slept
	m := make(map[int]map[string][]int)
	var date string
	var guardID, startSleep int
	for _, e := range events {
		switch {
		case strings.HasPrefix(e.msg, "Guard"):
			fmt.Sscanf(e.msg, "Guard #%d begins shift", &guardID)
			if e.t.Hour() == 0 {
				date = e.t.Format("01-02")
			} else {
				date = e.t.Add(time.Hour).Format("01-02")
			}
			if _, ok := m[guardID]; !ok {
				m[guardID] = make(map[string][]int)
			}
			m[guardID][date] = make([]int, 60, 60)
		case e.msg == "falls asleep":
			startSleep = e.t.Minute()
		case e.msg == "wakes up":
			endSleep := e.t.Minute()
			for i := startSleep; i < endSleep; i++ {
				m[guardID][date][i] = 1
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
			maxSlept, guardID = slept, guard
		}
	}

	sleepsPerMinute := make([]int, 60, 60)
	for _, mins := range m[guardID] {
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
	fmt.Println(guardID * minute)

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
				maxSleepsPerMinute, guardID, minute = sleep, guard, i
			}
		}
	}

	fmt.Println(guardID * minute)
}

package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	caves := map[string]map[string]struct{}{}
	s := util.ScanAll()
	for s.Scan() {
		sp := strings.Split(s.Text(), "-")
		from, to := sp[0], sp[1]
		if _, ok := caves[from]; !ok {
			caves[from] = map[string]struct{}{}
		}
		caves[from][to] = struct{}{}
		if _, ok := caves[to]; !ok {
			caves[to] = map[string]struct{}{}
		}
		caves[to][from] = struct{}{}
	}
	fmt.Println(
		paths(
			"start",
			caves,
			map[string]int{},
			func(next string, visited map[string]int) bool {
				return next == strings.ToUpper(next) || visited[next] == 0
			}))
	fmt.Println(
		paths(
			"start",
			caves,
			map[string]int{},
			func(next string, visited map[string]int) bool {
				if next == strings.ToUpper(next) || visited[next] == 0 {
					return true
				}
				for k, v := range visited {
					if k == strings.ToLower(k) && v > 1 {
						return false
					}
				}
				return true
			}))
}

func paths(start string, caves map[string]map[string]struct{}, visited map[string]int, visitable func(string, map[string]int) bool) int {
	var allPaths int
	for next := range caves[start] {
		switch next {
		case "start":
		case "end":
			allPaths++
		default:
			if visitable(next, visited) {
				nextVisited := map[string]int{}
				for k, v := range visited {
					nextVisited[k] = v
				}
				nextVisited[next]++
				allPaths += paths(next, caves, nextVisited, visitable)
			}
		}
	}
	return allPaths
}

package main

import (
	"fmt"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type valve struct {
	id      string
	flow    int
	tunnels map[string]int
}

func main() {
	valves := map[string]valve{}
	for _, line := range strings.Split(util.ReadAll(), "\n") {
		sp := strings.Split(line, "; ")
		var v valve
		fmt.Sscanf(sp[0], "Valve %s has flow rate=%d", &v.id, &v.flow)
		sp[1] = sp[1][len("tunnel leads to valve"):]
		if strings.HasPrefix(sp[1], "s") {
			sp[1] = sp[1][2:]
		} else {
			sp[1] = sp[1][1:]
		}
		v.tunnels = map[string]int{v.id: 0}
		for _, t := range strings.Split(sp[1], ", ") {
			v.tunnels[t] = 1
		}
		valves[v.id] = v
	}
	for k := range valves {
		for i := range valves {
			for j := range valves {
				dik, okik := valves[i].tunnels[k]
				dkj, okkj := valves[k].tunnels[j]
				if okik && okkj {
					dij, okij := valves[i].tunnels[j]
					if !okij || dij > dik+dkj {
						valves[i].tunnels[j] = dik + dkj
					}
				}
			}
		}
	}
	open := util.Set[string]{}
	for _, v := range valves {
		if v.flow > 0 {
			open[v.id] = struct{}{}
		}
	}

	fmt.Println(maxPressure(valves, "AA", 30, 0, open, 0))
}

func maxPressure(valves map[string]valve, curr string, minute int, pressure int, open util.Set[string], d int) int {
	max := pressure
	for v := range open {
		newopen := util.Set[string]{}
		for k := range open {
			newopen[k] = struct{}{}
		}
		delete(newopen, v)
		timeLeft := minute - valves[curr].tunnels[v] - 1
		if timeLeft > 0 {
			max = util.Max(max, maxPressure(valves, v, timeLeft, timeLeft*valves[v].flow+pressure, newopen, d+1))
		}
	}
	return max
}

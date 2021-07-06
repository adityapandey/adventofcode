package main

import (
	"container/heap"
	"fmt"
	"math/bits"
	"strings"

	"github.com/adityapandey/adventofcode/2019/machine"
	"github.com/adityapandey/adventofcode/util"
)

const checkpoint = "== Security Checkpoint =="

type room struct {
	name  string
	doors map[string]*room
}

func (r *room) unexplored() []string {
	// Manually explored later.
	if r.name == checkpoint {
		return []string{}
	}
	var ds []string
	for d := range r.doors {
		if r.doors[d] == nil {
			ds = append(ds, d)
		}
	}
	return ds
}

type game struct {
	in, out chan int
}

func newGame(program []int) *game {
	in := make(chan int, 25)
	out := machine.Run(program, in)
	return &game{in, out}
}

func (g *game) prompt() string {
	var sb strings.Builder
	for o := range g.out {
		fmt.Fprintf(&sb, "%c", o)
		if strings.HasSuffix(sb.String(), "Command?") {
			s := strings.Trim(sb.String(), "\n")
			sb.Reset()
			return s
		}
	}
	return sb.String()
}

func (g *game) send(cmd string) string {
	for i := 0; i < len(cmd); i++ {
		g.in <- int(cmd[i])
	}
	g.in <- int('\n')

	return g.prompt()
}

var reverse = map[string]string{"north": "south", "south": "north", "east": "west", "west": "east"}

func main() {
	var program []int
	for _, n := range strings.Split(util.ReadFile("input"), ",") {
		program = append(program, util.Atoi(n))
	}

	g := newGame(program)

	world := make(map[string]*room)
	var inventory []string

	prompt := g.prompt()
	r, _ := parseRoom(prompt)
	world[r.name] = &r

	curr := world[r.name]
	for {
		if len(curr.unexplored()) > 0 {
			// Choose some unexplored direction.
			dir := curr.unexplored()[0]
			prompt = g.send(dir)
			r, items := parseRoom(prompt)
			if _, ok := world[r.name]; !ok {
				world[r.name] = &r
			}
			next := world[r.name]
			for _, item := range items {
				g.send(fmt.Sprintf("take %s", item))
				inventory = append(inventory, item)
			}
			curr.doors[dir] = next
			next.doors[reverse[dir]] = curr
			curr = next
		} else {
			var next *room
			for r := range world {
				if len(world[r].unexplored()) > 0 {
					// Choose some room that is still unexplored.
					next = world[r]
					break
				}
			}
			if next == nil {
				// All rooms have been fully explored.
				break
			}
			p := path(world, curr.name, next.name)
			for _, dir := range p {
				g.send(dir)
			}
			curr = next
		}
	}

	// Get to the checkpoint.
	for _, p := range path(world, curr.name, checkpoint) {
		g.send(p)
	}

	// Try all the items.
	code := graycode(len(inventory))
	for i := 1; i < len(code); i++ {
		var cmd string
		x := code[i-1] ^ code[i]
		if code[i]&x == 0 {
			cmd = "drop"
		} else {
			cmd = "take"
		}
		g.send(fmt.Sprintf("%s %s", cmd, inventory[len(inventory)-1-bits.TrailingZeros(x)]))
		prompt := g.send("east")
		if strings.Contains(prompt, "proceed") {
			fmt.Println(prompt)
			break
		}
	}
}

var cursed = map[string]struct{}{"giant electromagnet": {}, "infinite loop": {}, "escape pod": {}, "molten lava": {}, "photons": {}}

func parseRoom(s string) (room, []string) {
	sp := strings.Split(s, "\n\n")
	r := room{
		name:  strings.Split(sp[0], "\n")[0],
		doors: make(map[string]*room),
	}
	for _, d := range strings.Split(sp[1], "\n")[1:] {
		r.doors[strings.Trim(d, "- ")] = nil
	}
	spp := strings.Split(sp[2], "\n")
	var items []string
	if strings.HasPrefix(spp[0], "Items") {
		for _, i := range spp[1:] {
			item := strings.Trim(i, "- ")
			if _, ok := cursed[item]; ok {
				continue
			}
			items = append(items, item)
		}
	}
	return r, items
}

func path(world map[string]*room, from, to string) []string {
	pq := util.PQ{&util.Item{world[from], 0}}
	heap.Init(&pq)
	pathmap := map[string][]string{from: {}}
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*util.Item).Obj.(*room)
		currpath := pathmap[curr.name]
		if curr.name == to {
			return currpath
		}
		for d := range curr.doors {
			next := curr.doors[d]
			if next == nil {
				continue
			}
			if p, ok := pathmap[next.name]; !ok || len(p) > 1+len(currpath) {
				nextpath := make([]string, len(currpath))
				copy(nextpath, currpath)
				nextpath = append(nextpath, d)
				pathmap[next.name] = nextpath
				heap.Push(&pq, &util.Item{next, -len(nextpath)})
			}
		}
	}
	panic("no path")
}

// n-bit graycode starting from all ones.
func graycode(n int) []uint {
	if n == 1 {
		return []uint{1, 0}
	}
	gn_1 := graycode(n - 1)
	var g []uint
	for i := range gn_1 {
		g = append(g, (1<<(n-1))+gn_1[i])
	}
	for i := range gn_1 {
		g = append(g, gn_1[len(gn_1)-1-i])
	}
	return g
}

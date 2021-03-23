package main

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/adityapandey/adventofcode/util"
)

type step struct {
	id   byte
	prev map[byte]struct{}
	next map[byte]struct{}
	time int
	w    *worker
}

func newStep(id byte) *step {
	s := &step{id: id}
	s.prev = make(map[byte]struct{})
	s.next = make(map[byte]struct{})
	s.time = int(id-'A') + 61
	return s
}

func (s *step) assigned() bool {
	return s.w != nil
}

type graph map[byte]*step

func newGraph() graph {
	return make(map[byte]*step)
}

func (g graph) size() int {
	return len(g)
}

func (g graph) add(from, to byte) {
	fromStep, ok := g[from]
	if !ok {
		fromStep = newStep(from)
	}
	toStep, ok := g[to]
	if !ok {
		toStep = newStep(to)
	}
	g[from] = fromStep
	g[to] = toStep
	fromStep.next[to] = struct{}{}
	toStep.prev[from] = struct{}{}
}

func (g graph) delete(id byte) {
	delStep := g[id]
	delete(g, id)
	for step := range delStep.next {
		delete(g[step].prev, id)
	}
}

func (g graph) readySteps() []*step {
	var candidates []*step
	for _, step := range g {
		if len(step.prev) == 0 {
			candidates = append(candidates, step)
		}
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i].id < candidates[j].id })
	return candidates
}

type worker struct {
	step *step
}

func (w *worker) assign(s *step) {
	w.step = s
	s.w = w
}

func (w *worker) unassign() {
	w.step = nil
}

func (w *worker) work() {
	if w.step != nil {
		w.step.time--
	}
}

func (w *worker) idle() bool {
	return w.step == nil
}

func main() {
	var log [][2]byte
	s := util.ScanAll()
	for s.Scan() {
		var from, to byte
		fmt.Sscanf(s.Text(), "Step %c must be finished before step %c can begin.", &from, &to)
		log = append(log, [2]byte{from, to})
	}

	// Part 1
	g := newGraph()
	for _, l := range log {
		g.add(l[0], l[1])
	}
	var b bytes.Buffer
	for g.size() > 0 {
		curr := g.readySteps()[0].id
		b.WriteByte(curr)
		g.delete(curr)
	}
	fmt.Println(b.String())

	// Part 2
	g = newGraph()
	for _, l := range log {
		g.add(l[0], l[1])
	}
	var seconds int
	workers := []*worker{{}, {}, {}, {}, {}}
	for g.size() > 0 {
		seconds++
		readySteps := g.readySteps()
		for _, step := range readySteps {
			if !step.assigned() {
				for _, w := range workers {
					if w.idle() {
						w.assign(step)
						break
					}
				}
			}
		}
		for _, w := range workers {
			w.work()
		}
		for _, step := range readySteps {
			if step.time == 0 {
				step.w.unassign()
				g.delete(step.id)
			}
		}
	}
	fmt.Println(seconds)
}

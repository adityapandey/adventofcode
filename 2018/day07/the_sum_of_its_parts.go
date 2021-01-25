package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

type E struct{}

type Step struct {
	id     byte
	prev   map[byte]E
	next   map[byte]E
	count  int
	worker *Worker
}

func NewStep(id byte) *Step {
	s := &Step{id: id}
	s.prev = make(map[byte]E)
	s.next = make(map[byte]E)
	s.count = int(id-'A') + 61
	return s
}

func (s *Step) Assigned() bool {
	return s.worker != nil
}

type Graph struct {
	m map[byte]*Step
}

func NewGraph() *Graph {
	g := new(Graph)
	g.m = make(map[byte]*Step)
	return g
}

func (g *Graph) Add(from, to byte) {
	fromStep, ok := g.m[from]
	if !ok {
		fromStep = NewStep(from)
	}
	toStep, ok := g.m[to]
	if !ok {
		toStep = NewStep(to)
	}
	g.m[from] = fromStep
	g.m[to] = toStep
	fromStep.next[to] = E{}
	toStep.prev[from] = E{}
}

func (g *Graph) Ready() []*Step {
	var candidates []*Step
	for _, step := range g.m {
		if len(step.prev) == 0 {
			candidates = append(candidates, step)
		}
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i].id < candidates[j].id })
	return candidates
}

func (g *Graph) Delete(id byte) {
	delStep := g.m[id]
	delete(g.m, id)
	for step, _ := range delStep.next {
		delete(g.m[step].prev, id)
	}
}

func (g *Graph) Size() int {
	return len(g.m)
}

type Worker struct {
	step *Step
}

func (w *Worker) Assign(s *Step) {
	w.step = s
	s.worker = w
}

func (w *Worker) Work() {
	if w.step != nil {
		w.step.count--
	}
}

func (w *Worker) IsIdle() bool {
	return w.step == nil
}

func main() {
	var log [][2]byte
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var from, to byte
		fmt.Sscanf(s.Text(), "Step %c must be finished before step %c can begin.", &from, &to)
		log = append(log, [2]byte{from, to})
	}

	// Part 1
	g := NewGraph()
	for _, l := range log {
		g.Add(l[0], l[1])
	}

	var b bytes.Buffer
	for g.Size() > 0 {
		ready := g.Ready()[0].id
		b.WriteByte(ready)
		g.Delete(ready)
	}

	// Part 2
	g = NewGraph()
	for _, l := range log {
		g.Add(l[0], l[1])
	}

	fmt.Println(b.String())
	b.Reset()
	var seconds int
	workers := []*Worker{{}, {}, {}, {}, {}}
	for g.Size() > 0 {
		seconds++
		ready := g.Ready()
		for _, step := range ready {
			if !step.Assigned() {
				for _, w := range workers {
					if w.IsIdle() {
						w.Assign(step)
						break
					}
				}
			}
		}
		for _, w := range workers {
			w.Work()
		}
		for _, step := range ready {
			if step.count == 0 {
				step.worker.step = nil
				g.Delete(step.id)
				b.WriteByte(step.id)
			}
		}
	}

	fmt.Println(seconds)
}

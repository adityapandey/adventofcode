package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

type group struct {
	nUnits     int
	hp         int
	ap         int
	attackType string
	initiative int
	immunities map[string]struct{}
	weaknesses map[string]struct{}
	label      string
	enemy      *[]*group
}

func (g group) effectivePower() int { return g.nUnits * g.ap }

func (g group) estimate(other group) int {
	if other.nUnits == 0 {
		return 0
	}
	if _, ok := other.immunities[g.attackType]; ok {
		return 0
	}
	estimate := g.effectivePower()
	if _, ok := other.weaknesses[g.attackType]; ok {
		estimate *= 2
	}
	return estimate
}

var groupRe = regexp.MustCompile(`(\d+) units each with (\d+) hit points( \((.*)\))* with an attack that does (\d+) (.+) damage at initiative (.+)`)

type world struct {
	groups       []*group
	immuneSystem []*group
	infection    []*group
}

func newWorld(groups []group) *world {
	var w world
	w.groups = make([]*group, len(groups))
	for i, g := range groups {
		w.groups[i] = new(group)
		*w.groups[i] = g
		if w.groups[i].label == "infection" {
			w.groups[i].enemy = &w.immuneSystem
			w.infection = append(w.infection, w.groups[i])
		} else {
			w.groups[i].enemy = &w.infection
			w.immuneSystem = append(w.immuneSystem, w.groups[i])
		}
	}
	return &w
}

func (w *world) run() (result string, remaining int) {
	for c := 0; c < 10000; c++ {
		// Target selection
		defenders := make(map[*group]struct{})
		type pair struct{ attacker, defender *group }
		var pairs []pair
		sort.Slice(w.groups, func(i, j int) bool {
			if w.groups[i].effectivePower() == w.groups[j].effectivePower() {
				return w.groups[i].initiative > w.groups[j].initiative
			}
			return w.groups[i].effectivePower() > w.groups[j].effectivePower()
		})

		for _, g := range w.groups {
			if g.nUnits == 0 {
				continue
			}
			enemy := make([]*group, len(*g.enemy))
			copy(enemy, *g.enemy)
			for i := 0; i < len(enemy); i++ {
				if _, ok := defenders[enemy[i]]; ok {
					enemy = append(enemy[:i], enemy[i+1:]...)
					i--
				}
			}
			sort.Slice(enemy, func(i, j int) bool {
				estI, estJ := g.estimate(*enemy[i]), g.estimate(*enemy[j])
				if estI == estJ {
					if enemy[i].effectivePower() == enemy[j].effectivePower() {
						return enemy[i].initiative > enemy[j].initiative
					}
					return enemy[i].effectivePower() > enemy[j].effectivePower()
				}
				return estI > estJ
			})
			selection := enemy[0]
			if g.estimate(*selection) > 0 {
				if _, ok := defenders[selection]; !ok {
					defenders[selection] = struct{}{}
					pairs = append(pairs, pair{g, selection})
				}
			}
		}

		// Attack
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].attacker.initiative > pairs[j].attacker.initiative
		})

		for _, p := range pairs {
			power := p.attacker.effectivePower()
			if _, ok := p.defender.weaknesses[p.attacker.attackType]; ok {
				power *= 2
			}
			p.defender.nUnits -= power / p.defender.hp
			if p.defender.nUnits < 0 {
				p.defender.nUnits = 0
			}
		}

		// Continue?
		var immuneSystemUnits, infectionUnits int
		for _, g := range w.groups {
			if g.enemy == &w.infection {
				immuneSystemUnits += g.nUnits
			} else {
				infectionUnits += g.nUnits
			}
		}
		if immuneSystemUnits == 0 {
			return w.infection[0].label, infectionUnits
		}
		if infectionUnits == 0 {
			return w.immuneSystem[0].label, immuneSystemUnits
		}
	}
	return "stalemate", -1
}

func main() {
	var label string
	var groups []group
	s := util.ScanAll()
	for s.Scan() {
		if s.Text() == "Immune System:" {
			label = "immuneSystem"
			continue
		}
		if s.Text() == "Infection:" {
			label = "infection"
			continue
		}
		ss := groupRe.FindAllStringSubmatch(s.Text(), -1)
		if len(ss) == 0 {
			continue
		}
		g := group{
			nUnits:     util.Atoi(ss[0][1]),
			hp:         util.Atoi(ss[0][2]),
			ap:         util.Atoi(ss[0][5]),
			attackType: ss[0][6],
			initiative: util.Atoi(ss[0][7]),
			label:      label,
		}
		g.immunities = make(map[string]struct{})
		g.weaknesses = make(map[string]struct{})

		if ss[0][4] != "" {
			for _, wi := range strings.Split(ss[0][4], "; ") {
				if strings.HasPrefix(wi, "weak to ") {
					for _, w := range strings.Split(wi[len("weak to "):], ", ") {
						g.weaknesses[w] = struct{}{}
					}
				}
				if strings.HasPrefix(wi, "immune to ") {
					for _, i := range strings.Split(wi[len("immune to "):], ", ") {
						g.immunities[i] = struct{}{}
					}
				}
			}
		}
		groups = append(groups, g)
	}

	// Part 1
	w := newWorld(groups)
	_, nUnits := w.run()
	fmt.Println(nUnits)

	// Part 2
	boost := 0
	for {
		boost++
		for i := range groups {
			if groups[i].label == "immuneSystem" {
				groups[i].ap++
			}
		}
		w := newWorld(groups)
		label, nUnits := w.run()
		if label == "immuneSystem" {
			fmt.Println(nUnits)
			break
		}
	}
}

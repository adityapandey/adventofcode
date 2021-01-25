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
)

type Group struct {
	nUnits     int
	hp         int
	ap         int
	attackType string
	initiative int
	immune     map[string]struct{}
	weakness   map[string]struct{}
	label      string
	enemy      *[]*Group
}

func (g Group) EffectivePower() int { return g.nUnits * g.ap }

func (g Group) Estimate(other Group) int {
	if other.nUnits == 0 {
		return 0
	}
	if _, ok := other.immune[g.attackType]; ok {
		return 0
	}
	estimate := g.EffectivePower()
	if _, ok := other.weakness[g.attackType]; ok {
		estimate *= 2
	}
	return estimate
}

var groupRe = regexp.MustCompile(`(\d+) units each with (\d+) hit points( \((.*)\))* with an attack that does (\d+) (.+) damage at initiative (.+)`)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

type World struct {
	groups       []*Group
	immuneSystem []*Group
	infection    []*Group
}

func NewWorld(groups []Group) *World {
	var w World
	w.groups = make([]*Group, len(groups))
	for i, g := range groups {
		w.groups[i] = new(Group)
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

func (w *World) Iterate() (result string, remaining int) {
	for c := 0; c < 10000; c++ {
		// Target selection
		defenders := make(map[*Group]struct{})
		type Pair struct{ attacker, defender *Group }
		var pairs []Pair
		sort.Slice(w.groups, func(i, j int) bool {
			if w.groups[i].EffectivePower() == w.groups[j].EffectivePower() {
				return w.groups[i].initiative > w.groups[j].initiative
			}
			return w.groups[i].EffectivePower() > w.groups[j].EffectivePower()
		})

		for _, g := range w.groups {
			if g.nUnits == 0 {
				continue
			}
			enemy := make([]*Group, len(*g.enemy))
			copy(enemy, *g.enemy)
			for i := 0; i < len(enemy); i++ {
				if _, ok := defenders[enemy[i]]; ok {
					enemy = append(enemy[:i], enemy[i+1:]...)
					i--
				}
			}
			sort.Slice(enemy, func(i, j int) bool {
				estI, estJ := g.Estimate(*enemy[i]), g.Estimate(*enemy[j])
				if estI == estJ {
					if enemy[i].EffectivePower() == enemy[j].EffectivePower() {
						return enemy[i].initiative > enemy[j].initiative
					}
					return enemy[i].EffectivePower() > enemy[j].EffectivePower()
				}
				return estI > estJ
			})
			selection := enemy[0]
			if g.Estimate(*selection) > 0 {
				if _, ok := defenders[selection]; !ok {
					defenders[selection] = struct{}{}
					pairs = append(pairs, Pair{g, selection})
				}
			}
		}

		// Attack
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].attacker.initiative > pairs[j].attacker.initiative
		})

		for _, p := range pairs {
			power := p.attacker.EffectivePower()
			if _, ok := p.defender.weakness[p.attacker.attackType]; ok {
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
	var groups []Group
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "Immune System:" {
			label = "immuneSystem"
			continue
		}
		if s.Text() == "Infection:" {
			label = "infection"
			continue
		}
		x := groupRe.FindAllStringSubmatch(s.Text(), -1)
		if len(x) == 0 {
			continue
		}
		g := Group{
			nUnits:     atoi(x[0][1]),
			hp:         atoi(x[0][2]),
			ap:         atoi(x[0][5]),
			attackType: x[0][6],
			initiative: atoi(x[0][7]),
			label:      label,
		}
		g.immune = make(map[string]struct{})
		g.weakness = make(map[string]struct{})

		if x[0][4] != "" {
			for _, wi := range strings.Split(x[0][4], "; ") {
				if strings.HasPrefix(wi, "weak to ") {
					for _, w := range strings.Split(wi[len("weak to "):], ", ") {
						g.weakness[w] = struct{}{}
					}
				}
				if strings.HasPrefix(wi, "immune to ") {
					for _, i := range strings.Split(wi[len("immune to "):], ", ") {
						g.immune[i] = struct{}{}
					}
				}
			}
		}
		groups = append(groups, g)
	}

	// Part 1
	w := NewWorld(groups)
	_, units := w.Iterate()
	fmt.Println(units)

	// Part 2
	boost := 0
	for {
		boost++
		for i := range groups {
			if groups[i].label == "immuneSystem" {
				groups[i].ap++
			}
		}
		w := NewWorld(groups)
		label, units := w.Iterate()
		if label == "immuneSystem" {
			fmt.Println(units)
			break
		}
	}
}

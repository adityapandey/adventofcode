package main

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

var spells = map[string]int{
	"missile":  53,
	"drain":    73,
	"shield":   113,
	"poison":   173,
	"recharge": 229,
}

type state struct {
	bossHP, playerHP        int
	bossDamage, playerArmor int
	playerMana              int
	effects                 map[string]int
	cost                    int
}

func (s state) copy() state {
	c := s
	c.effects = make(map[string]int)
	for e := range s.effects {
		c.effects[e] = s.effects[e]
	}
	return c
}

func (s *state) applyEffects() {
	for e := range s.effects {
		switch e {
		case "shield":
			s.playerArmor = 7
		case "poison":
			s.bossHP -= 3
		case "recharge":
			s.playerMana += 101
		}
		s.effects[e]--
		if s.effects[e] == 0 {
			if e == "shield" {
				s.playerArmor = 0
			}
			delete(s.effects, e)
		}
	}
}

func (s *state) applySpell(spell string) {
	switch spell {
	case "missile":
		s.bossHP -= 4
	case "drain":
		s.bossHP -= 2
		s.playerHP += 2
	case "shield", "poison":
		s.effects[spell] = 6
	case "recharge":
		s.effects[spell] = 5
	}
}

func main() {
	var bossHP, bossDamage int
	fmt.Fscanf(os.Stdin, "Hit Points: %d\n", &bossHP)
	fmt.Fscanf(os.Stdin, "Damage: %d", &bossDamage)

	s := state{
		bossHP:      bossHP,
		playerHP:    50,
		bossDamage:  bossDamage,
		playerArmor: 0,
		playerMana:  500,
		effects:     map[string]int{},
		cost:        0,
	}
	fmt.Println(leastCost(s, false))
	fmt.Println(leastCost(s, true))
}

func leastCost(s state, hardMode bool) int {
	pq := util.PQ{&util.Item{s, -s.cost}}
	heap.Init(&pq)
	var minCost int
	for pq.Len() > 0 {
		s := heap.Pop(&pq).(*util.Item).Obj.(state)
		if s.bossHP <= 0 {
			minCost = s.cost
			break
		}
		if hardMode {
			s.playerHP--
			if s.playerHP <= 0 {
				continue
			}
		}
		s.applyEffects()
		if s.bossHP <= 0 {
			heap.Push(&pq, &util.Item{s, -s.cost})
			continue
		}
		for spell, cost := range spells {
			if _, ok := s.effects[spell]; ok || cost > s.playerMana {
				continue
			}
			next := s.copy()
			next.cost += cost
			next.playerMana -= cost
			next.applySpell(spell)
			if next.bossHP <= 0 {
				heap.Push(&pq, &util.Item{next, -next.cost})
				continue
			}
			next.applyEffects()
			bossAttack := next.bossDamage - next.playerArmor
			if bossAttack < 1 {
				bossAttack = 1
			}
			next.playerHP -= bossAttack
			if next.playerHP > 0 {
				heap.Push(&pq, &util.Item{next, -next.cost})
			}
		}
	}
	return minCost
}

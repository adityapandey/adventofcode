package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

const shop = `Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`

type character struct {
	hp, attack, defense int
}

type item struct {
	cost, attack, defense int
}

func main() {
	var boss character
	fmt.Fscanf(os.Stdin, "Hit Points: %d\n", &boss.hp)
	fmt.Fscanf(os.Stdin, "Damage: %d\n", &boss.attack)
	fmt.Fscanf(os.Stdin, "Armor: %d\n", &boss.defense)

	min, max := math.MaxUint16, 0
	weapons, armor, rings := parseShop(shop)
	for _, w := range choose(weapons, 1) {
		for _, a := range chooseBetween(armor, 0, 1) {
			for _, r := range chooseBetween(rings, 0, 2) {
				var cost int
				var items []item
				items = append(items, append(w, append(a, r...)...)...)
				player := character{hp: 100}
				for _, i := range items {
					cost += i.cost
					player.attack += i.attack
					player.defense += i.defense
				}
				if win(player, boss) && cost < min {
					min = cost
				}
				if !win(player, boss) && cost > max {
					max = cost
				}
			}
		}
	}
	fmt.Println(min)
	fmt.Println(max)
}

func parseShop(shop string) ([]item, []item, []item) {
	s := strings.Split(shop, "\n\n")
	var weapons, armor, rings []item
	for _, ss := range strings.Split(s[0], "\n")[1:] {
		f := strings.Fields(ss)
		weapons = append(weapons, item{util.Atoi(f[1]), util.Atoi(f[2]), 0})
	}
	for _, ss := range strings.Split(s[1], "\n")[1:] {
		f := strings.Fields(ss)
		armor = append(armor, item{util.Atoi(f[1]), 0, util.Atoi(f[3])})
	}
	for _, ss := range strings.Split(s[2], "\n")[1:] {
		f := strings.Fields(ss)
		rings = append(rings, item{util.Atoi(f[2]), util.Atoi(f[3]), util.Atoi(f[4])})
	}
	return weapons, armor, rings
}

func win(player, boss character) bool {
	bossDelta := player.attack - boss.defense
	if bossDelta < 1 {
		bossDelta = 1
	}
	playerDelta := boss.attack - player.defense
	if playerDelta < 1 {
		playerDelta = 1
	}
	for {
		boss.hp -= bossDelta
		if boss.hp <= 0 {
			return true
		}
		player.hp -= playerDelta
		if player.hp <= 0 {
			return false
		}
	}
}

func chooseBetween(a []item, from, to int) [][]item {
	var ret [][]item
	for i := from; i <= to; i++ {
		ret = append(ret, choose(a, i)...)
	}
	return ret
}

func choose(a []item, n int) [][]item {
	if n == 0 {
		return [][]item{{}}
	}
	var ret [][]item
	for i := range a {
		b := make([]item, len(a)-1)
		copy(b, a)
		b = append(b[:i], a[i+1:]...)
		c := choose(b, n-1)
		for j := range c {
			c[j] = append(c[j], a[i])
			ret = append(ret, c[j])
		}
	}
	return ret
}

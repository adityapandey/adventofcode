package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`bot (\d+) gives low to ([a-z]+) (\d+) and high to ([a-z]+) (\d+)`)

type bot struct {
	chips []int
	toBot [2]bool
	to    [2]int
}

func (b *bot) give(n int) {
	if len(b.chips) >= 2 {
		log.Fatal(b)
	}
	b.chips = append(b.chips, n)
	sort.Ints(b.chips)
}

func main() {
	bots := make(map[int]*bot)
	outputs := make(map[int][]int)
	s := util.ScanAll()
	for s.Scan() {
		if strings.HasPrefix(s.Text(), "value") {
			var v, b int
			fmt.Sscanf(s.Text(), "value %d goes to bot %d", &v, &b)
			if _, ok := bots[b]; !ok {
				bots[b] = &bot{}
			}
			bots[b].give(v)
		} else {
			m := re.FindAllStringSubmatch(s.Text(), -1)[0]
			b := util.Atoi(m[1])
			if _, ok := bots[b]; !ok {
				bots[b] = &bot{}
			}
			bots[b].toBot[0] = m[2] == "bot"
			bots[b].to[0] = util.Atoi(m[3])
			bots[b].toBot[1] = m[4] == "bot"
			bots[b].to[1] = util.Atoi(m[5])
		}
	}

	for {
		halt := true
		for id, b := range bots {
			if len(b.chips) == 2 {
				halt = false
				if b.chips[0] == 17 && b.chips[1] == 61 {
					fmt.Println(id)
				}
				for i := 0; i < 2; i++ {
					if b.toBot[i] {
						bots[b.to[i]].give(b.chips[i])
					} else {
						outputs[b.to[i]] = append(outputs[b.to[i]], b.chips[i])
					}
				}
				b.chips = []int{}
			}
		}
		if halt {
			break
		}
	}
	fmt.Println(outputs[0][0] * outputs[1][0] * outputs[2][0])
}

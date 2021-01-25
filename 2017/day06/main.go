package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(input), "\t")
	banks := make([]int, len(s))
	for i := range s {
		banks[i] = util.Atoi(s[i])
	}
	seen := make(map[uint32]int)
	seen[hash(banks)] = 0
	l := len(banks)
	step, lastSeen := 0, 0
	for {
		maxi, max := 0, 0
		for i := 0; i < l; i++ {
			if banks[i] > max {
				maxi, max = i, banks[i]
			}
		}
		banks[maxi] = 0
		for i := (maxi + 1) % l; max > 0; max-- {
			banks[i]++
			i = (i + 1) % l
		}
		step++
		h := hash(banks)
		var ok bool
		if lastSeen, ok = seen[h]; ok {
			break
		}
		seen[h] = step
	}
	fmt.Println(step, step-lastSeen)
}

func hash(a []int) uint32 {
	return crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v", a)))
}

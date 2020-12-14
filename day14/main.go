package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mem1, mem2 := make(map[int64]int64), make(map[int64]int64)
	var mask string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "mask") {
			fmt.Sscanf(line, "mask = %s", &mask)
		} else {
			var address, val int64
			fmt.Sscanf(line, "mem[%d] = %d", &address, &val)
			mem1[address] = applyMask(mask, val)
			for _, addr := range applyFloatingMask(mask, address) {
				mem2[addr] = val
			}
		}
	}

	// Part 1
	var sum int64
	for _, v := range mem1 {
		sum += v
	}
	fmt.Println(sum)

	// Part 2
	sum = 0
	for _, v := range mem2 {
		sum += v
	}
	fmt.Println(sum)
}

func applyMask(mask string, val int64) int64 {
	var n, p int64 = 0, 1
	for i := len(mask) - 1; i >= 0; i-- {
		switch mask[i] {
		case 'X':
			n += p * (val % 2)
		case '1':
			n += p
		}
		val /= 2
		p *= 2
	}
	return n
}

func applyFloatingMask(mask string, address int64) []int64 {
	addresses := []int64{0}
	p := int64(1)
	for i := len(mask) - 1; i >= 0; i-- {
		switch mask[i] {
		case '0':
			for j := range addresses {
				addresses[j] += p * (address % 2)
			}
		case '1':
			for j := range addresses {
				addresses[j] += p
			}
		case 'X':
			l := len(addresses)
			for j := 0; j < l; j++ {
				addresses = append(addresses, addresses[j]+p)
			}
		}
		address /= 2
		p *= 2
	}
	return addresses
}

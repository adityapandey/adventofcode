package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var reactant map[byte]byte

func remove(b []byte, unit byte) []byte {
	molecule := make([]byte, len(b))
	copy(molecule, b)
	for i := 0; i < len(molecule); {
		if molecule[i] == unit || molecule[i] == reactant[unit] {
			molecule = append(molecule[:i], molecule[i+1:]...)
		} else {
			i++
		}
	}
	return molecule
}

func react(b []byte) []byte {
	molecule := make([]byte, len(b))
	copy(molecule, b)
	for i := 0; i < len(molecule)-1; {
		if molecule[i] == reactant[molecule[i+1]] {
			molecule = append(molecule[:i], molecule[i+2:]...)
			i--
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
	return molecule
}

func main() {
	reactant = make(map[byte]byte)
	lowers := []byte("abcdefghijklmnopqrstuvwxyz")
	uppers := bytes.ToUpper(lowers)
	for i := range lowers {
		reactant[lowers[i]] = uppers[i]
		reactant[uppers[i]] = lowers[i]
	}
	s := bufio.NewScanner(os.Stdin)
	var molecule []byte
	for s.Scan() {
		molecule = s.Bytes()
	}

	// Part 1
	minLen := len(react(molecule))
	fmt.Println(minLen)

	// Part 2
	for _, unit := range lowers {
		l := len(react(remove(molecule, unit)))
		if l < minLen {
			minLen = l
		}
	}
	fmt.Println(minLen)
}

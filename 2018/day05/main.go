package main

import (
	"bytes"
	"fmt"

	"github.com/adityapandey/adventofcode/util"
)

var neutralizer map[byte]byte

func main() {
	neutralizer = make(map[byte]byte)
	lowers := []byte("abcdefghijklmnopqrstuvwxyz")
	uppers := bytes.ToUpper(lowers)
	for i := range lowers {
		neutralizer[lowers[i]] = uppers[i]
		neutralizer[uppers[i]] = lowers[i]
	}

	molecule := util.ReadAll()
	minLen := len(react(molecule))
	fmt.Println(minLen)

	for _, unit := range lowers {
		l := len(react(removeAll(molecule, unit)))
		if l < minLen {
			minLen = l
		}
	}
	fmt.Println(minLen)
}

func react(b string) string {
	molecule := []byte(b)
	var i int
	for i < len(molecule)-1 {
		if molecule[i] == neutralizer[molecule[i+1]] {
			molecule = append(molecule[:i], molecule[i+2:]...)
			i--
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
	return string(molecule)
}

func removeAll(b string, unit byte) string {
	molecule := []byte(b)
	for i := 0; i < len(molecule); {
		if molecule[i] == unit || molecule[i] == neutralizer[unit] {
			molecule = append(molecule[:i], molecule[i+1:]...)
		} else {
			i++
		}
	}
	return string(molecule)
}

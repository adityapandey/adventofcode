package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

var re = regexp.MustCompile(`\[[a-z]+\]`)

func main() {
	var numTLS, numSSL int
	s := util.ScanAll()
	for s.Scan() {
		ip := s.Text()
		hypernetIndices := re.FindAllStringSubmatchIndex(ip, -1)
		var hypernets, supernets []string
		var start int
		for _, i := range hypernetIndices {
			hypernets = append(hypernets, ip[i[0]+1:i[1]-1])
			supernets = append(supernets, ip[start:i[0]])
			start = i[1]
		}
		supernets = append(supernets, ip[start:])
		tls, ssl := false, false
		var BABs []string
		for _, r := range supernets {
			if hasABBA(r) {
				tls = true
			}
			BABs = append(BABs, getBABs(r)...)
		}
		for _, h := range hypernets {
			if hasABBA(h) {
				tls = false
			}
			for _, bab := range BABs {
				if strings.Contains(h, bab) {
					ssl = true
				}
			}
		}
		if tls {
			numTLS++
		}
		if ssl {
			numSSL++
		}
	}
	fmt.Println(numTLS)
	fmt.Println(numSSL)
}

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func getBABs(s string) []string {
	var BABs []string
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			BABs = append(BABs, string([]byte{s[i+1], s[i], s[i+1]}))
		}
	}
	return BABs
}

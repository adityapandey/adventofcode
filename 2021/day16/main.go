package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

type typeid int

const (
	SUM typeid = iota
	PRODUCT
	MIN
	MAX
	LITERAL
	GT
	LT
	EQ
)

type packet struct {
	version    int
	id         typeid
	literal    int
	subpackets []packet
}

func (p packet) versionSum() int {
	sum := p.version
	for _, pp := range p.subpackets {
		sum += pp.versionSum()
	}
	return sum
}

func (p packet) value() int {
	switch p.id {
	case SUM:
		sum := 0
		for _, pp := range p.subpackets {
			sum += pp.value()
		}
		return sum
	case PRODUCT:
		prod := 1
		for _, pp := range p.subpackets {
			prod *= pp.value()
		}
		return prod
	case MIN:
		min := math.MaxInt
		for _, pp := range p.subpackets {
			min = util.Min(min, pp.value())
		}
		return min
	case MAX:
		max := 0
		for _, pp := range p.subpackets {
			max = util.Max(max, pp.value())
		}
		return max
	case LITERAL:
		return p.literal
	case GT:
		if p.subpackets[0].value() > p.subpackets[1].value() {
			return 1
		}
		return 0
	case LT:
		if p.subpackets[0].value() < p.subpackets[1].value() {
			return 1
		}
		return 0
	case EQ:
		if p.subpackets[0].value() == p.subpackets[1].value() {
			return 1
		}
		return 0
	}
	return 0
}

type parser struct {
	b   []byte
	pos int
}

func (p *parser) read(n int) []byte {
	b := p.b[p.pos : p.pos+n]
	p.pos += n
	return b
}

func (p *parser) parse() packet {
	var pkt packet
	pkt.version = dec(p.read(3))
	pkt.id = typeid(dec(p.read(3)))
	switch pkt.id {
	case LITERAL:
		val := 0
		more := byte('1')
		for more == '1' {
			more = p.read(1)[0]
			val = val<<4 + dec(p.read(4))
		}
		pkt.literal = val
	default:
		switch p.read(1)[0] {
		case '0':
			length := dec(p.read(15))
			currpos := p.pos
			for p.pos < currpos+length {
				pkt.subpackets = append(pkt.subpackets, p.parse())
			}
		case '1':
			for n := dec(p.read(11)); n > 0; n-- {
				pkt.subpackets = append(pkt.subpackets, p.parse())
			}
		}
	}
	return pkt
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	var b bytes.Buffer
	for _, c := range input {
		b.Write(bin(c))
	}

	parser2 := &parser{b.Bytes(), 0}
	pkt := parser2.parse()
	fmt.Println(pkt.versionSum())
	fmt.Println(pkt.value())
}

func bin(c byte) []byte {
	b := make([]byte, 4)
	var v int
	if c <= '9' {
		v = int(c - '0')
	} else {
		v = int(c-'A') + 10
	}
	for i := 0; i < 4; i++ {
		b[3-i] = '0' + byte(v%2)
		v /= 2
	}
	return b
}

func dec(b []byte) int {
	s, p := 0, 1
	for i := len(b) - 1; i >= 0; i-- {
		s += int(b[i]-'0') * p
		p *= 2
	}
	return s
}

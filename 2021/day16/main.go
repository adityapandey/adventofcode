package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"

	"github.com/adityapandey/adventofcode/util"
)

type state int

const (
	START state = iota
	LIT
	OP
	END
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
	lengthid   int
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
		} else {
			return 0
		}
	case LT:
		if p.subpackets[0].value() < p.subpackets[1].value() {
			return 1
		} else {
			return 0
		}
	case EQ:
		if p.subpackets[0].value() == p.subpackets[1].value() {
			return 1
		} else {
			return 0
		}
	}
	return 0
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	var b bytes.Buffer
	for _, c := range input {
		b.Write(bin(c))
	}
	packets, err := parse(&b, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(packets[0].versionSum())
	fmt.Println(packets[0].value())
}

func read(r io.Reader, n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := r.Read(b)
	return b, err
}

func parse(r io.Reader, n int) ([]packet, error) {
	var s state
	var packets []packet
	var p packet
	var literal []byte
	var done bool
	for !done {
		switch s {
		case START:
			p = packet{}
			b, err := read(r, 3)
			if err != nil {
				if err == io.EOF && n <= 0 {
					done = true
					break
				}
				return nil, err
			}
			p.version = dec(b)

			b, err = read(r, 3)
			if err != nil {
				return nil, err
			}
			p.id = typeid(dec(b))

			switch p.id {
			case LITERAL:
				s = LIT
			default:
				s = OP
			}
		case LIT:
			b, err := read(r, 1)
			if err != nil {
				return nil, err
			}
			l, err := read(r, 4)
			switch b[0] {
			case '1':
				if err != nil {
					return nil, err
				}
				literal = append(literal, l...)
				s = LIT
			case '0':
				if err != nil && err != io.EOF {
					return nil, err
				}
				literal = append(literal, l...)
				p.literal = dec(literal)
				literal = []byte{}
				s = END
			}
		case OP:
			b, err := read(r, 1)
			if err != nil {
				return nil, err
			}
			p.lengthid = dec(b)
			switch p.lengthid {
			case 0:
				b, err := read(r, 15)
				if err != nil {
					return nil, err
				}
				sublength := dec(b)
				b, err = read(r, sublength)
				if err != nil {
					return nil, err
				}
				subpackets, err := parse(bytes.NewBuffer(b), -1)
				if err != nil {
					return nil, err
				}
				p.subpackets = append(p.subpackets, subpackets...)
			case 1:
				b, err := read(r, 11)
				if err != nil {
					return nil, err
				}
				subpackets, err := parse(r, dec(b))
				if err != nil {
					return nil, err
				}
				p.subpackets = append(p.subpackets, subpackets...)
			}
			s = END
		case END:
			packets = append(packets, p)
			n--
			if n == 0 {
				done = true
			} else {
				s = START
			}
		}
	}
	return packets, nil
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

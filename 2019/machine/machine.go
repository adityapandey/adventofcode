package machine

import (
	"log"
)

type mode int

const (
	position mode = iota
	immediate
)

type opcode int

const (
	add opcode = 1 + iota
	mul
	input
	output
	jt
	jf
	lt
	eq
	halt opcode = 99
)

func decode(n int) (op opcode, modes [3]mode) {
	op = opcode(n % 100)
	n /= 100
	for i := 0; i < 3; i++ {
		modes[i] = mode(n % 10)
		n /= 10
	}
	return
}

type Machine struct {
	data map[int]int
	ip   int
	in   []int
	i    int
	out  []int
}

func New(program []int, in []int) *Machine {
	m := &Machine{
		data: make(map[int]int),
		in:   in,
	}
	for i, n := range program {
		m.data[i] = n
	}
	return m
}

func (m *Machine) get(i int, mo mode) int {
	switch mo {
	case immediate:
		return m.data[i]
	case position:
		return m.data[m.data[i]]
	default:
		log.Fatal("Unknown mode: ", mo)
	}
	return 0
}

func (m *Machine) set(i int, val int) {
	m.data[m.data[i]] = val
}

func (m *Machine) Step() bool {
	op, modes := decode(m.data[m.ip])
	switch op {
	case add:
		val := m.get(m.ip+1, modes[0]) + m.get(m.ip+2, modes[1])
		m.set(m.ip+3, val)
		m.ip += 4
	case mul:
		val := m.get(m.ip+1, modes[0]) * m.get(m.ip+2, modes[1])
		m.set(m.ip+3, val)
		m.ip += 4
	case input:
		m.set(m.ip+1, m.in[m.i])
		m.i++
		m.ip += 2
	case output:
		m.out = append(m.out, m.get(m.ip+1, modes[0]))
		m.ip += 2
	case jt:
		if m.get(m.ip+1, modes[0]) != 0 {
			m.ip = m.get(m.ip+2, modes[1])
		} else {
			m.ip += 3
		}
	case jf:
		if m.get(m.ip+1, modes[0]) == 0 {
			m.ip = m.get(m.ip+2, modes[1])
		} else {
			m.ip += 3
		}
	case lt:
		if m.get(m.ip+1, modes[0]) < m.get(m.ip+2, modes[1]) {
			m.set(m.ip+3, 1)
		} else {
			m.set(m.ip+3, 0)
		}
		m.ip += 4
	case eq:
		if m.get(m.ip+1, modes[0]) == m.get(m.ip+2, modes[1]) {
			m.set(m.ip+3, 1)
		} else {
			m.set(m.ip+3, 0)
		}
		m.ip += 4
	case halt:
		return false
	default:
		log.Fatal("Unknown opcode: ", op)
	}
	return true
}

func (m *Machine) Run() {
	for m.Step() {
	}
}

func (m *Machine) Output() []int {
	return m.out
}

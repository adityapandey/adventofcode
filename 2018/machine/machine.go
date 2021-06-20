package machine

type Machine struct {
	R []int
}

func New(n int) *Machine {
	return &Machine{
		R: make([]int, n),
	}
}

func (m *Machine) Execute(op string, a, b, c int) {
	Instructions[op](m, a, b, c)
}

type instruction func(m *Machine, a, b, c int)

var Instructions = map[string]instruction{
	"addr": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a] + m.R[b]
	},
	"addi": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a] + b
	},
	"mulr": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a] * m.R[b]
	},
	"muli": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a] * b
	},
	"banr": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a] & m.R[b]
	},
	"bani": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a] & b
	},
	"borr": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a] | m.R[b]
	},
	"bori": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a] | b
	},
	"setr": func(m *Machine, a, b, c int) {
		m.R[c] = m.R[a]
	},
	"seti": func(m *Machine, a, b, c int) {
		m.R[c] = a
	},
	"gtir": func(m *Machine, a, b, c int) {
		if a > m.R[b] {
			m.R[c] = 1
		} else {
			m.R[c] = 0
		}
	},
	"gtri": func(m *Machine, a, b, c int) {
		if m.R[a] > b {
			m.R[c] = 1
		} else {
			m.R[c] = 0
		}
	},
	"gtrr": func(m *Machine, a, b, c int) {
		if m.R[a] > m.R[b] {
			m.R[c] = 1
		} else {
			m.R[c] = 0
		}
	},
	"eqir": func(m *Machine, a, b, c int) {
		if a == m.R[b] {
			m.R[c] = 1
		} else {
			m.R[c] = 0
		}
	},
	"eqri": func(m *Machine, a, b, c int) {
		if m.R[a] == b {
			m.R[c] = 1
		} else {
			m.R[c] = 0
		}
	},
	"eqrr": func(m *Machine, a, b, c int) {
		if m.R[a] == m.R[b] {
			m.R[c] = 1
		} else {
			m.R[c] = 0
		}
	},
}

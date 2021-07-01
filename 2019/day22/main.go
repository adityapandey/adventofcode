package main

import (
	"bufio"
	"fmt"
	"math/big"
	"strings"

	"github.com/adityapandey/adventofcode/util"
)

const N = 10007

type deck [N]int

func (d deck) deal() deck {
	l := len(d)
	for i := 0; i < l/2; i++ {
		d[i], d[l-i-1] = d[l-i-1], d[i]
	}
	return d
}

func (d deck) cut(n int) deck {
	if n > 0 {
		c := make([]int, n)
		copy(c, d[:n])
		for i := 0; i < len(d)-n; i++ {
			d[i] = d[i+n]
		}
		for i := 0; i < n; i++ {
			d[i+len(d)-n] = c[i]
		}
	} else {
		n = util.Abs(n)
		c := make([]int, n)
		copy(c, d[len(d)-n:])
		for i := 0; i < len(d)-n; i++ {
			d[len(d)-i-1] = d[len(d)-n-i-1]
		}
		for i := 0; i < n; i++ {
			d[i] = c[i]
		}
	}
	return d
}

func (d deck) dealInc(n int) deck {
	var table deck
	var pos int
	for i := 0; i < len(d); i++ {
		table[pos] = d[i]
		pos += n
		pos %= len(d)
	}
	return table
}

func main() {
	var d deck
	for i := 0; i < len(d); i++ {
		d[i] = i
	}
	input := util.ReadAll()
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		switch {
		case strings.HasPrefix(line, "deal into"):
			d = d.deal()
		case strings.HasPrefix(line, "deal with increment"):
			var i int
			fmt.Sscanf(line, "deal with increment %d", &i)
			d = d.dealInc(i)
		default:
			var i int
			fmt.Sscanf(line, "cut %d", &i)
			d = d.cut(i)
		}
	}
	for i := 0; i < len(d); i++ {
		if d[i] == 2019 {
			fmt.Println(i)
			break
		}
	}

	/*
			deal into:
			x 0 1 2 3 4 5 6
			y 6 5 4 3 2 1 0
			  y = L-1-x

			cut N:
			x 0 1 2 3 4 5 6
			y 3 4 5 6 0 1 2
			  y = (x+N) % L

			deal with N:
			x 0 1 2 3 4 5 6
			y 0     1     2
			  x = (N*y) % L
			  y = (modinv(N, L)*x) % L

			y = (A*x + B) % L
			deal into: A = -1          , B = L-1
			      cut: A = 1           , B = N
		    deal with: A = modinv(N, L), B = 0

			f(x) = A1*x + B1
			g(x) = A2*x + B2
			f(g(x)) = A1*(A2*x + B2) + B1 = A1*A2*x + A1*B2 + B1
			  A = A1*A2, B = A1*B2 + B1

			f(f(x))    = A*(A*x + B) + B = A^2*x + A*B + B
			f(f(f(x))) = A*(A^2*x + A*B + B) + B = A^3*x + A^2*B + A*B + B
			f^n(x)     = A^n*x + B*(1 + A + ... + A^(n-1))
			  A = A^n
			  B = B*(1-A^n)/(1-A) = B * (1-A^n) * modinv(1-A, L)
	*/
	L := big.NewInt(119315717514047)
	n := big.NewInt(101741582076661)
	a, b := big.NewInt(1), big.NewInt(0)
	s = bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		var a2, b2 *big.Int
		switch {
		case strings.HasPrefix(line, "deal into"):
			a2, b2 = big.NewInt(-1), new(big.Int).Sub(L, big.NewInt(1))
		case strings.HasPrefix(line, "deal with increment"):
			var i int64
			fmt.Sscanf(line, "deal with increment %d", &i)
			a2, b2 = new(big.Int).ModInverse(big.NewInt(i), L), big.NewInt(0)
		default:
			var i int64
			fmt.Sscanf(line, "cut %d", &i)
			a2, b2 = big.NewInt(1), big.NewInt(i)
		}
		b2.Mul(b2, a)
		b.Add(b, b2)
		a.Mul(a, a2)
	}
	a_n := new(big.Int).Exp(a, n, L)
	modinv := new(big.Int).ModInverse(new(big.Int).Sub(big.NewInt(1), a), L)
	b.Mul(b, new(big.Int).Sub(big.NewInt(1), a_n))
	b.Mul(b, modinv)
	a.Exp(a, n, L)
	x := big.NewInt(2020)
	x.Mul(x, a)
	x.Add(x, b)
	x.Mod(x, L)
	fmt.Println(x)
}

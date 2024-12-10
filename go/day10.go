package main

import (
	"iter"
)

func day10Part1(lines iter.Seq[string]) int {
	mx := parseMatrix(lines)
	zeros := findPoints(mx, '0')
	sum := 0
	for _, p := range zeros {
		m := make(map[Point]bool)
		computeNines(mx, p, func(p Point) { m[p] = true })
		sum += len(m)
	}
	return sum
}

func day10Part2(lines iter.Seq[string]) int {
	mx := parseMatrix(lines)
	zeros := findPoints(mx, '0')
	sum := 0
	for _, p := range zeros {
		c := 0
		computeNines(mx, p, func(p Point) { c++ })
		sum += c
	}
	return sum

}

func computeNines(mx [][]byte, p Point, nineFound func(Point)) {
	n := len(mx)
	if !p.isInRage(n) {
		return
	}
	pval := mx[p.x][p.y]
	if pval == '9' {
		nineFound(p)
		return
	}

	if u := p.Up(); u.isInRage(n) && mx[u.x][u.y]-1 == pval {
		computeNines(mx, u, nineFound)
	}
	if d := p.Down(); d.isInRage(n) && mx[d.x][d.y]-1 == pval {
		computeNines(mx, d, nineFound)
	}
	if r := p.Right(); r.isInRage(n) && mx[r.x][r.y]-1 == pval {
		computeNines(mx, r, nineFound)
	}
	if l := p.Left(); l.isInRage(n) && mx[l.x][l.y]-1 == pval {
		computeNines(mx, l, nineFound)
	}
}

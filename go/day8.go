package main

import (
	"iter"
)

type day8 struct{}

func (d day8) Part1(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	antenas := findAllAntenas(mx)
	n := len(mx)
	points := make(map[Point]bool)

	for _, antenaGroup := range antenas {
		for i, a1 := range antenaGroup {
			for _, a2 := range antenaGroup[i+1:] {
				dx := abs(a1.x - a2.x)
				dy := abs(a1.y - a2.y)
				var p1 Point
				var p2 Point
				// TODO: extract this operations into a methdod?
				if a1.x < a2.x {
					p1.x = a1.x - dx
					p2.x = a2.x + dx
				} else if a1.x > a2.x {
					p1.x = a1.x + dx
					p2.x = a2.x - dx
				} else {
					p1.x = a1.x
					p2.x = a1.x
				}
				if a1.y < a2.y {
					p1.y = a1.y - dy
					p2.y = a2.y + dy
				} else if a1.y > a2.y {
					p1.y = a1.y + dy
					p2.y = a2.y - dy
				} else {
					p1.y = a1.y
					p2.y = a1.y
				}
				if p1.isInRage(n) {
					points[p1] = true
				}
				if p2.isInRage(n) {
					points[p2] = true
				}
			}
		}
	}

	return len(points)
}

func (d day8) Part2(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	antenas := findAllAntenas(mx)
	n := len(mx)
	points := make(map[Point]bool)

	for _, antenaGroup := range antenas {
		if len(antenaGroup) > 1 {
			for _, a := range antenaGroup {
				points[a] = true
			}
		}
		for i, a1 := range antenaGroup {
			for _, a2 := range antenaGroup[i+1:] {
				dx := abs(a1.x - a2.x)
				dy := abs(a1.y - a2.y)
				var p1 Point
				var p2 Point
				dxx := dx
				dyy := dy
				for {
					if a1.x < a2.x {
						p1.x = a1.x - dxx
						p2.x = a2.x + dxx
					} else if a1.x > a2.x {
						p1.x = a1.x + dxx
						p2.x = a2.x - dxx
					} else {
						p1.x = a1.x
						p2.x = a1.x
					}
					if a1.y < a2.y {
						p1.y = a1.y - dyy
						p2.y = a2.y + dyy
					} else if a1.y > a2.y {
						p1.y = a1.y + dyy
						p2.y = a2.y - dyy
					} else {
						p1.y = a1.y
						p2.y = a1.y
					}
					if p1.isInRage(n) {
						points[p1] = true
					}
					if p2.isInRage(n) {
						points[p2] = true
					}
					if !p1.isInRage(n) && !p2.isInRage(n) {
						break
					}
					dxx += dx
					dyy += dy
				}
			}
		}
	}

	return len(points)
}

func findAllAntenas(mx [][]byte) map[byte][]Point {
	antenas := make(map[byte][]Point)
	for i, row := range mx {
		for j, c := range row {
			if isDigit(c) || isAlpha(c) {
				antenas[c] = append(antenas[c], Point{i, j})
			}
		}
	}
	return antenas
}

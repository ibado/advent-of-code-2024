package main

import (
	"aoc2024/ds"
	"iter"
	"maps"
	"slices"
)

type day6 struct{}

func (d day6) Part1(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	start := findPoints(mx, '^')[0]
	return len(findSeen(mx, start))
}

func (d day6) Part2(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	start := findPoints(mx, '^')[0]

	visited := maps.Keys(findSeen(mx, start))

	loops := 0
	for v := range visited {
		if mx[v.x][v.y] == '#' || v.x == start.x && v.y == start.y {
			continue
		}
		mxc := make([][]byte, len(mx))
		for i := range mx {
			mxc[i] = make([]byte, len(mx[i]))
			copy(mxc[i], mx[i])
		}

		mxc[v.x][v.y] = '#'
		if isLoop(mxc, start) {
			loops++
		}
	}

	return loops
}

func findSeen(mx [][]byte, start Point) ds.Set[Point] {
	seen := ds.Set[Point]{}
	seen.Add(start)
	dir := Point{-1, 0}
	n := start
	np := start.Plus(dir)
	for np.isInRage(len(mx)) {
		if mx[np.x][np.y] == '#' {
			np = n
			dir = dir.RotateRight()
		} else {
			seen.Add(np)
			n = np
			np = np.Plus(dir)
		}
	}
	return seen
}

type pdir struct {
	p, dir Point
}

func isLoop(mx [][]byte, start Point) bool {
	// using a bool slice since Set[PointDir] aka map[PointDir]struct{} is too slow
	seen := make([]bool, len(mx)*len(mx)*4)
	dir := Point{-1, 0}
	n := start
	np := start.Plus(dir)
	for np.isInRage(len(mx)) {
		if mx[np.x][np.y] == '#' {
			np = n
			dir = dir.RotateRight()
		} else {
			idx := (np.x*len(mx)+np.y)*4 + slices.Index(Dirs, dir)
			if seen[idx] {
				return true
			}
			seen[idx] = true
			n = np
			np = np.Plus(dir)
		}
	}

	return false
}

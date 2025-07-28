package main

import (
	"iter"
	"slices"
)

func (p Point) move(dir byte) Point {
	var res Point
	switch dir {
	case '>':
		res = p.Right()
	case '<':
		res = p.Left()
	case '^':
		res = p.Up()
	case 'v':
		res = p.Down()
	}
	return res
}

type day15 struct{}

func (d day15) Part1(lines iter.Seq[string]) any {
	mx, dirs := parseDay15(lines)
	robotPoint := findPoints(mx, '@')[0]

	for _, dir := range dirs {
		nd := robotPoint.move(dir)
		switch mx[nd.x][nd.y] {
		case '.':
			mx[robotPoint.x][robotPoint.y] = '.'
			robotPoint = nd
			mx[robotPoint.x][robotPoint.y] = '@'
		case 'O':
			di := nd.move(dir)
			for mx[di.x][di.y] == 'O' {
				di = di.move(dir)
			}

			if mx[di.x][di.y] == '.' {
				mx[robotPoint.x][robotPoint.y] = '.'
				robotPoint = nd
				mx[robotPoint.x][robotPoint.y] = '@'
				mx[di.x][di.y] = 'O'
			}
		case '#':
			// do nth
		}
	}

	sum := 0
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[0]); j++ {
			if mx[i][j] == 'O' {
				sum += i*100 + j
			}
		}
	}

	return sum
}

func (d day15) Part2(lines iter.Seq[string]) any {
	mx, dirs := parseDay15(lines)
	mx = expandMxPart2(mx)
	robotPoint := findPoints(mx, '@')[0]

	for _, dir := range dirs {
		nextPos := robotPoint.move(dir)
		switch mx[nextPos.x][nextPos.y] {
		case '.':
			mx[robotPoint.x][robotPoint.y] = '.'
			robotPoint = nextPos
			mx[robotPoint.x][robotPoint.y] = '@'
		case '[', ']':
		// TODO: implement me
		case '#':
			// do nth
		}
	}

	sum := 0
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[0]); j++ {
			if mx[i][j] == '[' {
				sum += i*100 + j
			}
		}
	}

	return 0
}

func expandMxPart2(mx [][]byte) [][]byte {
	var dmx [][]byte
	for _, row := range mx {
		var r []byte
		for _, c := range row {
			switch c {
			case '#':
				r = append(r, '#')
				r = append(r, '#')
			case 'O':
				r = append(r, '[')
				r = append(r, ']')
			case '.':
				r = append(r, '.')
				r = append(r, '.')
			case '@':
				r = append(r, '@')
				r = append(r, '.')
			}
		}
		dmx = append(dmx, r)
	}
	return dmx
}

func parseDay15(lines iter.Seq[string]) ([][]byte, []byte) {
	s := slices.Collect(lines)
	idx := slices.IndexFunc(s, func(e string) bool { return e == "" })
	mx := parseMatrix(slices.Values(s[:idx]))
	var dirs []byte
	for _, l := range s[idx:] {
		for i, _ := range l {
			dirs = append(dirs, l[i])
		}
	}
	return mx, dirs
}

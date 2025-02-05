package main

import (
	"iter"
	"maps"
	"slices"
)

var dirs = []byte{'U', 'R', 'D', 'L'}

type PointDir struct {
	x, y int
	dir  byte
}

func rotateRight(dir byte) byte {
	idx := slices.Index(dirs, dir)
	assert(idx >= 0)
	return dirs[(idx+1)%len(dirs)]
}

type day6 struct{}

func (d day6) Part1(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	start := findPoints(mx, '^')[0]
	i := start.x
	j := start.y

	var dir byte = 'U'
	m := make(map[Point]int)
	for {
		found := true
		if dir == 'U' {
			for i = i - 1; i >= 0; i-- {
				if mx[i][j] == '#' {
					dir = rotateRight(dir)
					i++
					found = false
					break
				}
				m[Point{i, j}] = 1
			}
		} else if dir == 'R' {
			for j = j + 1; j < len(mx); j++ {
				if mx[i][j] == '#' {
					dir = rotateRight(dir)
					j--
					found = false
					break
				}
				m[Point{i, j}] = 1
			}
		} else if dir == 'D' {
			for i = i + 1; i < len(mx); i++ {
				if mx[i][j] == '#' {
					dir = rotateRight(dir)
					i--
					found = false
					break
				}
				m[Point{i, j}] = 1
			}
		} else if dir == 'L' {
			for j = j - 1; j >= 0; j-- {
				if mx[i][j] == '#' {
					dir = rotateRight(dir)
					j++
					found = false
					break
				}
				m[Point{i, j}] = 1
			}
		}

		if found {
			break
		}
	}

	return len(m) + 1
}

func (d day6) Part2(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	start := findPoints(mx, '^')[0]

	visited := maps.Keys(traverseMap(mx, start))

	mp := make(map[Point]int)
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
			mp[Point{v.x, v.y}] = 1
		}
	}

	return len(mp)
}

// TODO: clean up this mess, isLoop is almost the same as traverseMap
func isLoop(mx [][]byte, start Point) bool {
	i := start.x
	j := start.y
	var dir byte = 'U'
	m := make([]bool, len(mx)*len(mx)*4)
	for {
		found := true
		if dir == 'U' {
			for i = i - 1; i >= 0; i-- {
				if mx[i][j] == '#' {
					dir = 'R'
					i++
					found = false
					break
				}
				h := (i*len(mx)+j)*4 + 0
				if m[h] {
					return true
				}
				m[h] = true

			}
		} else if dir == 'R' {
			for j = j + 1; j < len(mx); j++ {
				if mx[i][j] == '#' {
					dir = 'D'
					j--
					found = false
					break
				}

				h := (i*len(mx)+j)*4 + 1
				if m[h] {
					return true
				}
				m[h] = true
			}
		} else if dir == 'D' {
			for i = i + 1; i < len(mx); i++ {
				if mx[i][j] == '#' {
					dir = 'L'
					i--
					found = false
					break
				}
				h := (i*len(mx)+j)*4 + 2
				if m[h] {
					return true
				}
				m[h] = true
			}
		} else if dir == 'L' {
			for j = j - 1; j >= 0; j-- {
				if mx[i][j] == '#' {
					dir = 'U'
					j++
					found = false
					break
				}
				h := (i*len(mx)+j)*4 + 3
				if m[h] {
					return true
				}
				m[h] = true
			}
		}

		if found {
			break
		}
	}

	return false
}

func traverseMap(mx [][]byte, start Point) map[PointDir]byte {
	i := start.x
	j := start.y
	var dir byte = 'U'
	m := make(map[PointDir]byte)
	for {
		found := true
		if dir == 'U' {
			for i = i - 1; i >= 0; i-- {
				if mx[i][j] == '#' {
					dir = 'R'
					i++
					found = false
					break
				}
				if m[PointDir{i, j, dir}] == 1 {
					return nil
				}
				m[PointDir{i, j, dir}] = 1

			}
		} else if dir == 'R' {
			for j = j + 1; j < len(mx); j++ {
				if mx[i][j] == '#' {
					dir = 'D'
					j--
					found = false
					break
				}
				if m[PointDir{i, j, dir}] == 1 {
					return nil
				}
				m[PointDir{i, j, dir}] = 1
			}
		} else if dir == 'D' {
			for i = i + 1; i < len(mx); i++ {
				if mx[i][j] == '#' {
					dir = 'L'
					i--
					found = false
					break
				}
				if m[PointDir{i, j, dir}] == 1 {
					return nil
				}
				m[PointDir{i, j, dir}] = 1
			}
		} else if dir == 'L' {
			for j = j - 1; j >= 0; j-- {
				if mx[i][j] == '#' {
					dir = 'U'
					j++
					found = false
					break
				}
				if m[PointDir{i, j, dir}] == 1 {
					return nil
				}
				m[PointDir{i, j, dir}] = 1
			}
		}

		if found {
			break
		}
	}

	return m
}

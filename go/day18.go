package main

import (
	"iter"
)

type day18 struct{}

const Limit = 71
const BytesToScan = 1024

func (d day18) Part1(lines iter.Seq[string]) any {
	obstacles := make(map[Point]bool)
	idx := 0
	for l := range lines {
		nums := parseNums([]byte(l))
		x := nums[0]
		y := nums[1]
		obstacles[Point{x, y}] = true
		idx++
		if idx == BytesToScan {
			break
		}
	}

	graph := genGridGraph(obstacles, Limit)
	_, level := bfs(graph)
	return level
}

func (d day18) Part2(lines iter.Seq[string]) any {
	var obstacles []Point
	for l := range lines {
		nums := parseNums([]byte(l))
		x := nums[0]
		y := nums[1]
		obstacles = append(obstacles, Point{x, y})
	}

	min := BytesToScan + 1
	max := len(obstacles)
	i := (max + min) / 2
	for i != min {
		graph := genGridGraph(asMap(obstacles[:i]), Limit)
		found, _ := bfs(graph)
		if found {
			min = i
			i = (i + max) / 2
		} else {
			max = i
			i = (i + min) / 2
		}
	}
	return obstacles[max-1]
}

func asMap(s []Point) map[Point]bool {
	m := make(map[Point]bool)
	for _, si := range s {
		m[si] = true
	}
	return m
}

// Starts at 0,0 and search for Limit,Limit
// returns (true, level) if found it
// returns (false,   -1) otherwise
func bfs(graph map[Point][]Point) (found bool, level int) {
	seen := make(map[Point]bool)
	var q Queue[Point]
	q.Push(Point{0, 0})
	lvl := 0
	for q.Len() != 0 {
		lvlSize := q.Len()
		for lvlSize > 0 {
			point := q.Pop()
			if !seen[point] {
				if point.x == Limit-1 && point.y == Limit-1 {
					return true, lvl
				} else {
					q.PushSlice(graph[point])
					seen[point] = true
				}
			}
			lvlSize--
		}
		lvl++
	}

	return false, -1
}

func genGridGraph(obs map[Point]bool, size int) map[Point][]Point {
	graph := make(map[Point][]Point)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			p := Point{j, i}
			if obs[p] {
				continue
			}
			if r := p.Right(); r.isInRage(size) && !obs[r] {
				graph[p] = append(graph[p], r)
			}
			if l := p.Left(); l.isInRage(size) && !obs[l] {
				graph[p] = append(graph[p], l)
			}
			if u := p.Up(); u.isInRage(size) && !obs[u] {
				graph[p] = append(graph[p], u)
			}
			if d := p.Down(); d.isInRage(size) && !obs[d] {
				graph[p] = append(graph[p], d)
			}
		}
	}
	return graph
}

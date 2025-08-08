package main

import (
	"aoc2024/ds"
	"iter"
)

type day12 struct{}

type node struct {
	p      Point
	letter byte
}

type graph map[node][]node

func (d day12) Part1(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	gridSize := len(mx)
	var g graph = genGraph(mx, gridSize)
	components := connectedComp(g)
	total := 0
	for _, component := range components {
		area := len(component)
		perimeter := perimeter(component)
		price := perimeter * area
		total += price
	}
	return total
}

func (d day12) Part2(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	gridSize := len(mx)
	var g graph = genGraph(mx, gridSize)
	components := connectedComp(g)
	total := 0
	for _, component := range components {
		area := len(component)
		sides := sides(component)
		price := sides * area
		total += price
	}
	return total

}

func genGraph(mx [][]byte, gridSize int) graph {
	g := make(map[node][]node)
	for i, row := range mx {
		for j, c := range row {
			n := node{Point{i, j}, c}
			for _, d := range Dirs {
				if other := n.p.Plus(d); other.isInRage(gridSize) {
					g[n] = append(g[n], node{other, mx[other.x][other.y]})
				}
			}
		}
	}

	return g
}

func connectedComp(g graph) (components [][]node) {
	points := make(map[Point]byte)
	for n := range g {
		if _, ok := points[n.p]; ok {
			continue
		}
		seen := make(map[node]bool)
		var comp []node
		var q ds.Queue[node]
		q.Push(n)
		for q.Len() != 0 {
			node := q.Pop()
			if !seen[node] {
				if node.letter == n.letter {
					q.PushSlice(g[node])
					seen[node] = true
					points[node.p] = node.letter
					comp = append(comp, node)
				}
			}
		}
		components = append(components, comp)
	}

	return components
}

func perimeter(nodes []node) int {
	m := make(map[Point]bool)
	count := 0
	for _, node := range nodes {
		m[node.p] = true
	}
	for point := range m {
		for _, d := range Dirs {
			if neighbor := point.Plus(d); !m[neighbor] {
				count++
			}
		}
	}
	return count
}

var diagonals = []Point{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func sides(nodes []node) int {
	points := make(map[Point]bool)
	for _, node := range nodes {
		points[node.p] = true
	}
	// in a rectilinear shape, the vertices are the same as sides
	vertices := 0
	for p := range points {
		for i, d := range Dirs {
			if !points[p.Plus(d)] && !points[p.Plus(Dirs[mod(i+1, 4)])] {
				vertices++
			}
		}
		for _, d := range diagonals {
			if pd := p.Plus(d); !points[pd] {
				count := 0
				for _, dd := range Dirs {
					// good luck future me understanding this heuristic
					if pdnext := pd.Plus(dd); points[pdnext] &&
						abs(pdnext.x-p.x) <= 1 &&
						abs(pdnext.y-p.y) <= 1 {
						count++
					}
				}
				if count >= 2 {
					vertices++
				}
			}
		}
	}

	return vertices
}

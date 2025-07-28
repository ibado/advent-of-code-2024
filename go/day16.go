package main

import (
	"aoc2024/ds"
	"iter"
	"math"
	"slices"
)

type day16 struct{}

type DirPoint struct {
	p, dir Point
}

func (d day16) Part1(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	g := generateGraph(mx)
	start := findPoints(mx, 'S')[0]
	end := findPoints(mx, 'E')[0]
	costs, _ := g.Dijkstra(DirPoint{p: start, dir: Point{0, 1}})
	res := math.MaxInt
	for _, d := range Dirs {
		res = min(res, costs[DirPoint{p: end, dir: d}])
	}
	return res
}

func (d day16) Part2(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	g := generateGraph(mx)
	start := findPoints(mx, 'S')[0]
	end := findPoints(mx, 'E')[0]
	startDirPoint := DirPoint{p: start, dir: Point{0, 1}}
	costs, parents := g.Dijkstra(startDirPoint)

	endDirPoint := slices.MinFunc([]DirPoint{
		{p: end, dir: Dirs[0]},
		{p: end, dir: Dirs[1]},
		{p: end, dir: Dirs[2]},
		{p: end, dir: Dirs[3]},
	}, func(a, b DirPoint) int { return costs[a] - costs[b] })

	return countTiles(parents, endDirPoint)
}

func countTiles(parents map[DirPoint][]DirPoint, start DirPoint) int {
	visited := make(map[DirPoint]bool)
	var recFunc func(start DirPoint)
	recFunc = func(start DirPoint) {
		visited[start] = true
		for _, parent := range parents[start] {
			if !visited[parent] {
				recFunc(parent)
			}
		}
	}
	recFunc(start)

	// only count the Points no matter the direction
	tiles := make(map[Point]bool)
	for dirPoint := range visited {
		tiles[dirPoint.p] = true
	}

	return len(tiles)
}

func generateGraph(mx [][]byte) ds.WGraph[DirPoint] {
	g := make(map[DirPoint][]ds.WGraphNode[DirPoint])

	newNode := func(p, d Point, weight int) ds.WGraphNode[DirPoint] {
		return ds.NewNode(DirPoint{p, d}, weight)
	}

	for i, row := range mx {
		for j, c := range row {
			if c != '#' {
				p := Point{i, j}
				for i, d := range Dirs {
					dr := DirPoint{p, d}
					g[dr] = append(g[dr], newNode(p, Dirs[mod(i+1, 4)], 1000))
					g[dr] = append(g[dr], newNode(p, Dirs[mod(i-1, 4)], 1000))
					if np := p.Plus(d); np.isInRage(len(mx)) && mx[np.x][np.y] != '#' {
						g[dr] = append(g[dr], newNode(np, d, 1))
					}
				}
			}
		}
	}

	return g
}

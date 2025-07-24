package main

import (
	"container/heap"
	"iter"
	"math"
)

type day16 struct{}

func (d day16) Part1(lines iter.Seq[string]) any {
	mx := parseMatrix(lines)
	g := generateGraph(mx)
	src := findPoints(mx, 'S')[0]
	dst := findPoints(mx, 'E')[0]
	costs := g.Dijkstra(DirPoint{p: src, dir: Point{0, 1}})
	m := math.MaxInt
	for _, d := range Dirs {
		m = min(m, costs[DirPoint{dir: d, p: dst}])
	}
	return m
}

func (d day16) Part2(lines iter.Seq[string]) any {
	return 0
}

type DirPoint struct {
	p, dir Point
}

var Dirs = [4]Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func generateGraph(mx [][]byte) WGraph[DirPoint] {
	g := make(map[DirPoint][]GraphNode[DirPoint])

	for i, row := range mx {
		for j, c := range row {
			if c != '#' {
				p := Point{i, j}
				for i, d := range Dirs {
					dr := DirPoint{p, d}
					g[dr] = append(g[dr], GraphNode[DirPoint]{DirPoint{p, Dirs[mod(i+1, 4)]}, 1000})
					g[dr] = append(g[dr], GraphNode[DirPoint]{DirPoint{p, Dirs[mod(i-1, 4)]}, 1000})
					if np := p.Plus(d); np.isInRage(len(mx)) && mx[np.x][np.y] != '#' {
						g[dr] = append(g[dr], GraphNode[DirPoint]{DirPoint{np, d}, 1})
					}
				}
			}
		}
	}

	return g
}

type GraphNode[T any] struct {
	id     T
	weight int
}

type NHeap[T any] []GraphNode[T]

func (nh *NHeap[T]) Push(x any) {
	*nh = append(*nh, x.(GraphNode[T]))
}

func (nh *NHeap[T]) Pop() any {
	assertMsg(len(*nh) > 0, "heap is empty!")
	lastIdx := len(*nh) - 1
	last := (*nh)[lastIdx]
	*nh = (*nh)[:lastIdx]
	return last
}

func (nh NHeap[T]) Less(i, j int) bool {
	return nh[i].weight < nh[j].weight
}

func (nh NHeap[T]) Swap(i, j int) {
	tmp := nh[i]
	nh[i] = nh[j]
	nh[j] = tmp
}
func (nh NHeap[T]) Len() int { return len(nh) }

type WGraph[T comparable] map[T][]GraphNode[T]

func (g WGraph[T]) Dijkstra(src T) map[T]int {
	costs := make(map[T]int)
	parents := make(map[T]T)
	var costHeap NHeap[T]
	heap.Init(&costHeap)

	heap.Push(&costHeap, GraphNode[T]{src, 0})
	costs[src] = 0

	for costHeap.Len() != 0 {
		minCostNode := heap.Pop(&costHeap).(GraphNode[T])

		for _, node := range g[minCostNode.id] {
			newCost := node.weight + minCostNode.weight
			if nodeCost, ok := costs[node.id]; ok && newCost < nodeCost || !ok {
				costs[node.id] = newCost
				heap.Push(&costHeap, GraphNode[T]{node.id, newCost})
				parents[minCostNode.id] = node.id
			}
		}
	}

	return costs
}

func assertMsg(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

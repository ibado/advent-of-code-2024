package main

import (
	"container/heap"
	"iter"
	"math"
	"slices"
)

type day16 struct{}

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

type DirPoint struct {
	p, dir Point
}

func generateGraph(mx [][]byte) WGraph[DirPoint] {
	g := make(map[DirPoint][]GraphNode[DirPoint])

	newNode := func(p, d Point, weight int) GraphNode[DirPoint] {
		return GraphNode[DirPoint]{DirPoint{p, d}, weight}
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

func (g WGraph[T]) Dijkstra(src T) (costs map[T]int, parents map[T][]T) {
	costs = make(map[T]int)
	parents = make(map[T][]T)
	var costHeap NHeap[T]
	heap.Init(&costHeap)

	heap.Push(&costHeap, GraphNode[T]{src, 0})
	costs[src] = 0

	for costHeap.Len() != 0 {
		minCostNode := heap.Pop(&costHeap).(GraphNode[T])

		for _, node := range g[minCostNode.id] {
			newCost := node.weight + minCostNode.weight
			nodeCost, ok := costs[node.id]
			if ok && newCost < nodeCost || !ok {
				costs[node.id] = newCost
				heap.Push(&costHeap, GraphNode[T]{node.id, newCost})
				parents[node.id] = []T{minCostNode.id}
			} else if newCost == nodeCost {
				parents[node.id] = append(parents[node.id], minCostNode.id)
			}
		}
	}

	return costs, parents
}

func assertMsg(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

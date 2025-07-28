package ds

import (
	"container/heap"
)

type WGraph[T comparable] map[T][]WGraphNode[T]

type WGraphNode[T any] struct {
	id     T
	weight int
}

func NewNode[T any](id T, weight int) WGraphNode[T] {
	return WGraphNode[T]{id, weight}
}

func (g WGraph[T]) Dijkstra(src T) (costs map[T]int, parents map[T][]T) {
	costs = make(map[T]int)
	parents = make(map[T][]T)
	var costHeap nodeHeap[T]
	heap.Init(&costHeap)

	heap.Push(&costHeap, WGraphNode[T]{src, 0})
	costs[src] = 0

	for costHeap.Len() != 0 {
		minCostNode := heap.Pop(&costHeap).(WGraphNode[T])

		for _, node := range g[minCostNode.id] {
			newCost := node.weight + minCostNode.weight
			nodeCost, ok := costs[node.id]
			if ok && newCost < nodeCost || !ok {
				costs[node.id] = newCost
				heap.Push(&costHeap, WGraphNode[T]{node.id, newCost})
				parents[node.id] = []T{minCostNode.id}
			} else if newCost == nodeCost {
				parents[node.id] = append(parents[node.id], minCostNode.id)
			}
		}
	}

	return costs, parents
}

// heap for dijkstra implementation
type nodeHeap[T any] []WGraphNode[T]

func (nh *nodeHeap[T]) Push(x any) {
	*nh = append(*nh, x.(WGraphNode[T]))
}

func (nh *nodeHeap[T]) Pop() any {
	if len(*nh) == 0 {
		panic("heap is empty!")
	}
	lastIdx := len(*nh) - 1
	last := (*nh)[lastIdx]
	*nh = (*nh)[:lastIdx]
	return last
}

func (nh nodeHeap[T]) Less(i, j int) bool {
	return nh[i].weight < nh[j].weight
}

func (nh nodeHeap[T]) Swap(i, j int) {
	tmp := nh[i]
	nh[i] = nh[j]
	nh[j] = tmp
}

func (nh nodeHeap[T]) Len() int { return len(nh) }

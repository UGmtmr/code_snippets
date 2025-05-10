package codesnippets

import (
	"container/heap"
)

func dijkstra(n int, edges [][]Edge, startPoint int) []int {
	vertices := make([]int, n)
	for i := range vertices {
		vertices[i] = 2147483647
	}
	vertices[startPoint] = 0
	pq := &IntHeap{}
	heap.Init(pq)
	heap.Push(pq, Node{startPoint, 0})

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		v, d := node.id, node.cost
		if d > vertices[v] {
			continue
		}
		for _, edge := range edges[v] {
			if vertices[edge.to] > d+edge.cost {
				vertices[edge.to] = d + edge.cost
				heap.Push(pq, Node{edge.to, vertices[edge.to]})
			}
		}
	}

	return vertices
}

type Edge struct {
	to   int
	cost int
}

type Node struct {
	id   int
	cost int
}

// 整数ヒープの実装(minヒープ)
type IntHeap []Node

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].cost < h[j].cost } // maxヒープにしたいときはh[i] > h[j]に修正して使う
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

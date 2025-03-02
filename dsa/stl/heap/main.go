package main

import (
	"container/heap"
	"fmt"
)

/*
container/heap manages a heap using a slice (array-based binary heap).

- priority queue min-heap and max-heap



*/

type PriorityQueue []int

// Implement heap.Interface
func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i] > h[j] } // Min-Heap: Smallest first
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1] // last element
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &PriorityQueue{}
	heap.Init(h)
	heap.Push(h, 10)
	heap.Push(h, 5)
	heap.Push(h, 121)
	heap.Push(h, 14)
	heap.Push(h, 1)
	heap.Push(h, 1000)
	fmt.Println("Heap after inserts:", *h)
	heap.Pop(h)
	fmt.Println("Heap after inserts:", *h)
	heap.Pop(h)
	fmt.Println("Heap after inserts:", *h)
	heap.Pop(h)
	fmt.Println("Heap after inserts:", *h)
	fmt.Println("Heap after inserts:", h.Len())

	value := heap.Pop(h).(int)
	fmt.Println(value)

}

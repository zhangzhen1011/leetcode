package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Key int
	Val int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val > pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	*pq = old[:n-1]
	return old[n-1]
}

func main() {

	pq := make(PriorityQueue, 0)

	for i := 0; i < 10; i++ {
		item := Item{Key: i, Val: i}
		pq.Push(&item)
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		i := heap.Pop(&pq)
		item := i.(*Item)
		fmt.Println(item.Val)
	}
}

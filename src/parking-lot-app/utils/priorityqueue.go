package utils

import "sync"

// PriorityQueue is min-heap of ints.
// This will help to assaign the nearest free slot.
type PriorityQueue []int

// To perform atomic operation on priority queue
var lock = &sync.Mutex{}

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	lock.Lock()
	defer lock.Unlock()

	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	lock.Lock()
	defer lock.Unlock()

	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]

	return x
}

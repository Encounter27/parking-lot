package utils

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue_Pop(t *testing.T) {
	pq := new(PriorityQueue)
	heap.Init(pq)

	heap.Push(pq, 4)
	heap.Push(pq, 1)
	heap.Push(pq, 2)
	heap.Push(pq, 3)

	i := 1
	j := 3
	t.Run("Priority queue", func(t *testing.T) {
		for pq.Len() > 2 {
			x := heap.Pop(pq)
			l := pq.Len()
			if x != i || l != j {
				t.Errorf("%d == %d - %d == %d\n", x, i, l, j)
			}

			i++
			j--
		}
	})

	i = 1
	j = 3
	heap.Push(pq, 2)
	heap.Push(pq, 1)

	t.Run("Priority queue", func(t *testing.T) {
		for pq.Len() > 0 {
			x := heap.Pop(pq)
			l := pq.Len()
			if x != i || l != j {
				t.Errorf("%d == %d - %d == %d\n", x, i, l, j)
			}

			i++
			j--
		}
	})
}

package priorityqueue

import "container/heap"

type PriorityQueue[T any] struct {
	items pqHeapInterface[T]
}

func New[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(&pq.items)
	return pq
}

func (pq *PriorityQueue[T]) Push(item T, priority int) {
	heap.Push(&pq.items, pqItem[T]{
		value:    item,
		priority: priority,
	})
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.items.Len()
}

func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(&pq.items).(pqItem[T]).value //nolint:forcetypeassert
}

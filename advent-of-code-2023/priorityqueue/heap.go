package priorityqueue

// from the example here: https://pkg.go.dev/container/heap

import (
	"container/heap"
)

// A pqItem is something we manage in a priority queue.
type pqItem[T any] struct {
	value    T   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A pqHeapInterface implements heap.Interface and holds Items.
type pqHeapInterface[T any] []*pqItem[T]

func (pq pqHeapInterface[T]) Len() int { return len(pq) }

func (pq pqHeapInterface[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq pqHeapInterface[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pqHeapInterface[T]) Push(x any) {
	n := len(*pq)
	item := x.(*pqItem[T]) //nolint:forcetypeassert
	item.index = n
	*pq = append(*pq, item)
}

func (pq *pqHeapInterface[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an item in the queue.
func (pq *pqHeapInterface[T]) update(item *pqItem[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

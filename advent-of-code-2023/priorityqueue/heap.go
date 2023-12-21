package priorityqueue

// from the example here: https://pkg.go.dev/container/heap

import (
	"container/heap"
)

// A pqItem is something we manage in a priority queue.
type pqItem struct {
	value    any // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A pqHeapInterface implements heap.Interface and holds Items.
type pqHeapInterface []*pqItem

func (pq pqHeapInterface) Len() int { return len(pq) }

func (pq pqHeapInterface) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq pqHeapInterface) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pqHeapInterface) Push(x any) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *pqHeapInterface) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an item in the queue.
func (pq *pqHeapInterface) update(item *pqItem, value any, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

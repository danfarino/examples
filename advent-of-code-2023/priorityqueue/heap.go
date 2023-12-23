package priorityqueue

// from the example here: https://pkg.go.dev/container/heap

// A pqItem is something we manage in a priority queue.
type pqItem[T any] struct {
	value    T   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
}

// A pqHeapInterface implements heap.Interface and holds Items.
type pqHeapInterface[T any] []pqItem[T]

func (pq pqHeapInterface[T]) Len() int { return len(pq) }

func (pq pqHeapInterface[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq pqHeapInterface[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pqHeapInterface[T]) Push(x any) {
	item := x.(pqItem[T]) //nolint:forcetypeassert
	*pq = append(*pq, item)
}

func (pq *pqHeapInterface[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	//old[n-1] = *new(pqItem[T]) // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

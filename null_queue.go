package queue

// just for benchmarking
type NullQueue[T any] struct {
	v T
}

// NewBoundedQueue create a BoundedQueue.
func NewNullQueue[T any](value T) *NullQueue[T] {
	return &NullQueue[T]{
		v: value,
	}
}

// Enqueue puts the given value v at the tail of the queue.
// If this queue if full, the caller will be blocked.
func (q *NullQueue[T]) Enqueue(v T) {
}

// Dequeue removes and returns the value at the head of the queue.
// It will be blocked if the queue is empty.
func (q *NullQueue[T]) Dequeue() T {
	return q.v
}

// Len returns length of this queue.
func (q *NullQueue[T]) Len() int {
	return 0
}

package queue

import (
	"math/rand"
	"sync/atomic"
	"testing"
)

func BenchmarkQueue(b *testing.B) {
	queues := map[string]Queue[int]{
		"lock-free queue":   NewLKQueue[int](),
		"two-lock queue":    NewCQueue[int](),
		"slice-based queue": NewSliceQueue[int](0),
	}

	length := 1 << 12
	inputs := make([]int, length)
	for i := 0; i < length; i++ {
		v := rand.Intn(100)
		inputs = append(inputs, v)
	}

	for name, q := range queues {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			var c int64
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					i := int(atomic.AddInt64(&c, 1)-1) % length
					v := inputs[i]
					if v >= 50 {
						q.Enqueue(v)
					} else {
						q.Dequeue()
					}
				}
			})
		})
	}
}

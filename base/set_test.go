package base

import (
	"sync"
	"testing"
)

func TestThreadSafeSet_Add(t *testing.T) {
	tsSet := MakeThreadSafeSet[int]()

	N := 10

	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			tsSet.Add(i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < N; i++ {
		CheckEqual(tsSet.UnsafeHas(i), true, t)
	}
}

func BenchmarkThreadSafeSet_(b *testing.B) {
	tsSet := MakeThreadSafeSet[int]()

	tsSet.Add(1)
	b.Run("Has", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = tsSet.Has(1)
			_ = tsSet.Has(2)
		}
	})

	b.Run("UnsafeHas", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = tsSet.UnsafeHas(1)
			_ = tsSet.UnsafeHas(2)
		}
	})
}

package singleton

import (
	"sync"
	"testing"
)

func BenchmarkGetInstanceDoOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(10000)
		for k := 0; k < 10000; k++ {
			go func() {
				defer wg.Done()
				_ = GetInstance_DoOnce()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkGetInstance_checked_atomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(10000)
		for k := 0; k < 10000; k++ {
			go func() {
				defer wg.Done()
				_ = GetInstance_checked_atomic()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkGetInstance_locking_always(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(10000)
		for k := 0; k < 10000; k++ {
			go func() {
				defer wg.Done()
				_ = GetInstance_locking_always()
			}()
		}
		wg.Wait()
	}
}

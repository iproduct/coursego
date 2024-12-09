package benchmarks

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	b.Run("BasicMutex-3-0", func(b *testing.B) {
		BasicMutex_Load(b, 3, 0)
	})
	b.Run("BasicMutex-6-1", func(b *testing.B) {
		BasicMutex_Load(b, 6, 1)
	})
	b.Run("BasicMutex-3-3", func(b *testing.B) {
		BasicMutex_Load(b, 3, 3)
	})
	b.Run("BasicMutex-1-6", func(b *testing.B) {
		BasicMutex_Load(b, 1, 6)
	})
	b.Run("BasicMutex-0-6", func(b *testing.B) {
		BasicMutex_Load(b, 0, 6)
	})
	b.Run("RWMutex-3-0", func(b *testing.B) {
		RWMutex_Load(b, 3, 0)
	})
	b.Run("RWMutex-6-1", func(b *testing.B) {
		RWMutex_Load(b, 6, 1)
	})
	b.Run("RWMutex-3-3", func(b *testing.B) {
		RWMutex_Load(b, 3, 3)
	})
	b.Run("RWMutex-1-6", func(b *testing.B) {
		RWMutex_Load(b, 1, 6)
	})
	b.Run("RWMutex-0-6", func(b *testing.B) {
		RWMutex_Load(b, 0, 6)
	})
}

// compare BenchmarkBasicMutex_Load with BenchmarkRWMutex_Load
func BasicMutex_Load(b *testing.B, nLoads int, nStores int) {
	mu := BasicMutex{}
	mu.Store(10)
	for i := 0; i < b.N; i++ {
		for j := 0; j < nLoads; j++ {
			go mu.Load()
		}
		for j := 0; j < nStores; j++ {
			go mu.Store(i)
		}
	}
}

func RWMutex_Load(b *testing.B, nLoads int, nStores int) {
	mu := RWMutex{}
	mu.Store(10)
	for i := 0; i < b.N; i++ {
		for j := 0; j < nLoads; j++ {
			go mu.Load()
		}
		for j := 0; j < nStores; j++ {
			go mu.Store(i)
		}
	}
}

//func BenchmarkBasicMutex_Store(b *testing.B) {
//	mu := BasicMutex{}
//	for i := 0; i < b.N; i++ {
//		go mu.Store(i)
//	}
//}
//
//func BenchmarkBasicMutex_Hybrid(b *testing.B) {
//	mu := BasicMutex{}
//	for i := 0; i < b.N; i++ {
//		go mu.Load()
//		go mu.Load()
//		go mu.Load()
//		go mu.Store(i)
//	}
//}
//
//// compare BenchmarkRWMutex_Load with BenchmarkBasicMutex_Load
//func BenchmarkRWMutex_Load(b *testing.B) {
//	mu := RWMutex{}
//	mu.Store(10)
//	for i := 0; i < b.N; i++ {
//		go mu.Load()
//		go mu.Load()
//		go mu.Load()
//	}
//}
//
//func BenchmarkRWMutex_Store(b *testing.B) {
//	mu := RWMutex{}
//	for i := 0; i < b.N; i++ {
//		go mu.Store(i)
//	}
//}
//
//func BenchmarkRWMutex_Hybrid(b *testing.B) {
//	mu := RWMutex{}
//	for i := 0; i < b.N; i++ {
//		go mu.Load()
//		go mu.Load()
//		go mu.Load()
//		go mu.Store(i)
//	}
//}

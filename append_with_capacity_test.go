// go test -run=none -bench=. -v -test.v -benchmem append_with_capacity_test.go
package benchmark

import "testing"

func Benchmark_Append_WithCapacity(b *testing.B) {
	slice := make([]int, 0 /* length */, b.N /* capacity */)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
}

func Benchmark_Append_WithoutCapacity(b *testing.B) {
	slice := make([]int, 0 /* length */, 0 /* capacity */)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
}

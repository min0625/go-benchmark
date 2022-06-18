// go test -run=none -bench=. -v -test.v -benchmem append_vs_index_assign_test.go
package benchmark

import "testing"

func Benchmark_Slice_Append(b *testing.B) {
	slice := make([]int, 0 /* length */, b.N /* capacity */)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
}

func Benchmark_Slice_IndexAssign(b *testing.B) {
	slice := make([]int, b.N /* length */)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		slice[i] = i
	}
}

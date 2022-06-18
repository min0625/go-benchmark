// go test -run=none -bench=. -v -test.v -benchmem map_size_hint_test.go
package benchmark

import "testing"

func Benchmark_Map_WithHint(b *testing.B) {
	m := make(map[int]int, b.N /* size hint */)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func Benchmark_Map_WithoutHint(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

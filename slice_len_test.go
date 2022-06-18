// go test -run=none -bench=. -v -test.v -benchmem slice_len_test.go
package benchmark

import "testing"

func Benchmark_Len_MultipleCall(b *testing.B) {
	slice := make([]int, 1000)
	var sum int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += len(slice)
	}
}

func Benchmark_Len_SingleCall(b *testing.B) {
	slice := make([]int, 1000)
	var sum int

	b.ResetTimer()
	length := len(slice)
	for i := 0; i < b.N; i++ {
		sum += length
	}
}

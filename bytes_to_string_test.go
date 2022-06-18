// go test -run=none -bench=. -v -test.v -benchmem bytes_to_string_test.go
package benchmark

import (
	"testing"
)

func Benchmark_BytesToString_MultipleCall(b *testing.B) {
	data := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = string(data)
	}
}

func Benchmark_BytesToString_SingleCall(b *testing.B) {
	data := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	b.ResetTimer()

	s := string(data)
	for i := 0; i < b.N; i++ {
		_ = s
	}
}

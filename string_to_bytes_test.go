// go test -run=none -bench=. -v -test.v -benchmem string_to_bytes_test.go
package benchmark

import (
	"testing"
)

func Benchmark_StringToBytes_MultipleCall(b *testing.B) {
	data := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = []byte(data)
	}
}

func Benchmark_StringToBytes_SingleCall(b *testing.B) {
	data := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	b.ResetTimer()

	bytesData := []byte(data)
	for i := 0; i < b.N; i++ {
		_ = bytesData
	}
}

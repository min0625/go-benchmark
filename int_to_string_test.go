// go test -run=none -bench=. -v -test.v -benchmem int_to_string_test.go
package benchmark

import (
	"fmt"
	"strconv"
	"testing"
)

func Benchmark_IntToString_itoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(i)
	}
}

func Benchmark_IntToString_Sprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(i)
	}
}

func Benchmark_IntToString_Sprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", i)
	}
}

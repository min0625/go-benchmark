// go test -run=none -bench=. -v -test.v -benchmem string_concatenation_test.go
package benchmark

import (
	"strings"
	"testing"
)

func Benchmark_StringConcatenation_WithBuilder(b *testing.B) {
	data := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	var builder strings.Builder

	for i := 0; i < b.N; i++ {
		_, err := builder.WriteString(data)
		if err != nil {
			b.Fatal(err)
		}
	}

	result := builder.String()
	_ = result
}

func Benchmark_StringConcatenation_WithPlusOperator(b *testing.B) {
	data := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	var result string

	for i := 0; i < b.N; i++ {
		result += data
	}

	_ = result
}

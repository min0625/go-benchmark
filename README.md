# go-benchmark

## map size hint
```sh
$ go test -run=none -bench=. -v -test.v -benchmem map_size_hint_test.go
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_Map_WithHint
Benchmark_Map_WithHint-8        15478039               100.9 ns/op            41 B/op          0 allocs/op
Benchmark_Map_WithoutHint
Benchmark_Map_WithoutHint-8      9017071               162.0 ns/op            78 B/op          0 allocs/op
PASS
ok      command-line-arguments  3.318s
```

## slice len
```sh
$ go test -run=none -bench=. -v -test.v -benchmem slice_len_test.go
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_Len_MultipleCall
Benchmark_Len_MultipleCall-8    1000000000               0.2445 ns/op          0 B/op          0 allocs/op
Benchmark_Len_SingleCall
Benchmark_Len_SingleCall-8      1000000000               0.2445 ns/op          0 B/op          0 allocs/op
PASS
ok      command-line-arguments  0.577s
```

## int to string
```sh
$ go test -run=none -bench=. -v -test.v -benchmem int_to_string_test.go
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_IntToString_itoa
Benchmark_IntToString_itoa-8            49804102                23.76 ns/op            7 B/op          0 allocs/op
Benchmark_IntToString_Sprint
Benchmark_IntToString_Sprint-8          16376864                70.86 ns/op           16 B/op          1 allocs/op
Benchmark_IntToString_Sprintf
Benchmark_IntToString_Sprintf-8         17159427                71.01 ns/op           16 B/op          1 allocs/op
PASS
ok      command-line-arguments  3.777s
```

## bytes to string
```sh
$ go test -run=none -bench=. -v -test.v -benchmem bytes_to_string_test.go
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_BytesToString_MultipleCall
Benchmark_BytesToString_MultipleCall-8          49451703                22.50 ns/op           64 B/op          1 allocs/op
Benchmark_BytesToString_SingleCall
Benchmark_BytesToString_SingleCall-8            1000000000               0.2582 ns/op          0 B/op          0 allocs/op
PASS
ok      command-line-arguments  1.461s
```

## string to bytes
```sh
$ go test -run=none -bench=. -v -test.v -benchmem string_to_bytes_test.go
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_StringToBytes_MultipleCall
Benchmark_StringToBytes_MultipleCall-8          44851430                30.63 ns/op           64 B/op          1 allocs/op
Benchmark_StringToBytes_SingleCall
Benchmark_StringToBytes_SingleCall-8            1000000000               0.2484 ns/op          0 B/op          0 allocs/op
PASS
ok      command-line-arguments  2.568s
```

## append with capacity
```sh
$ go test -run=none -bench=. -v -test.v -benchmem append_with_capacity_test.go
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_Append_WithCapacity
Benchmark_Append_WithCapacity-8         1000000000               1.200 ns/op           0 B/op          0 allocs/op
Benchmark_Append_WithoutCapacity
Benchmark_Append_WithoutCapacity-8      166767306                6.700 ns/op          42 B/op          0 allocs/op
PASS
ok      command-line-arguments  6.416s
```

## append vs index assign
```sh
$ go test -run=none -bench=. -v -test.v -benchmem append_vs_index_assign_test.go
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_Slice_Append
Benchmark_Slice_Append-8                1000000000               1.090 ns/op           0 B/op          0 allocs/op
Benchmark_Slice_IndexAssign
Benchmark_Slice_IndexAssign-8           1000000000               1.113 ns/op           0 B/op          0 allocs/op
PASS
ok      command-line-arguments  6.953s
```

## string concatenation
```sh
$ go test -run=none -bench=. -v -test.v -benchmem string_concatenation_test.go
goos: windows
goarch: amd64
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
Benchmark_StringConcatenation_WithBuilder
Benchmark_StringConcatenation_WithBuilder-8             13478571                83.73 ns/op          340 B/op          0 allocs/op
Benchmark_StringConcatenation_WithPlusOperator
Benchmark_StringConcatenation_WithPlusOperator-8           36211            190210 ns/op         1126625 B/op          1 allocs/op
PASS
ok      command-line-arguments  8.490s
```

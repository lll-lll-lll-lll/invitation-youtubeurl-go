package backend

import (
	"fmt"
	"testing"
	"unsafe"
)

var data = []byte("はろーhogehogeです。")

func Benchmark_Sprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", data)
	}
}

func Benchmark_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(data)
	}
}

func Benchmark_Unsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = *(*string)(unsafe.Pointer(&data))
	}
}

package fibo

import (
	"testing"
)

func BenchmarkFibo1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fibo1(1000000)
	}
}

func BenchmarkFibo2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fibo2(1000000)
	}
}
func BenchmarkFibo3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fibo3(1000000)
	}
}
func BenchmarkFibo4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fibo4(45)
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/mooncaker816/CodePieces/fibonaccis
// BenchmarkFibo1-4   	    1000	   2385281 ns/op
// BenchmarkFibo2-4   	     500	   2973291 ns/op
// BenchmarkFibo3-4   	      10	 125219384 ns/op
// BenchmarkFibo4-4   	       1	7664120762 ns/op
// PASS
// ok  	github.com/mooncaker816/CodePieces/fibonaccis	13.474s

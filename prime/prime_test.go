package prime

import (
	"fmt"
	"testing"
)

func TestPrime(t *testing.T) {
	fmt.Println(PrintPrimeUntil(10))
	fmt.Println(PrintPrimeUntil(100))
}

func TestPrime2(t *testing.T) {
	fmt.Println(PrintPrimeUntil2(10))
	fmt.Println(PrintPrimeUntil2(100))
}

func TestPrime3(t *testing.T) {
	fmt.Println(PrintPrimeUntil3(10))
	fmt.Println(PrintPrimeUntil3(100))
}

var funcs = []struct {
	name string
	f    func(n int) []int
}{
	{"goroutines", PrintPrimeUntil},
	{"iterator", PrintPrimeUntil2},
	{"goroutines2", PrintPrimeUntil3},
}

func BenchmarkPrime(b *testing.B) {
	for _, f := range funcs {
		for n := 1000; n <= 10000; n = n + 1000 {
			b.Run(fmt.Sprintf("%s/%d", f.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.f(n)
				}
			})
		}
	}
}

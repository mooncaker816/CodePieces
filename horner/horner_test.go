package horner

import (
	"fmt"
	"testing"
)

var funcs = []struct {
	name string
	f    func(float64, ...float64) float64
}{
	// {"horner", horner},
	// {"hornerChan", hornerChan},
	{"hornerChan2", hornerChan2},
}

func BenchmarkHorner(b *testing.B) {
	for _, f := range funcs {
		for n := 100; n <= 110; n++ {
			coeff := make([]float64, n)
			for i := 0; i < n; i++ {
				coeff = append(coeff, float64(i)+1.12345)
			}
			b.Run(fmt.Sprintf("%s-%d\n", f.name, n), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f.f(1.12345, coeff...)
				}
			})
		}
	}
}

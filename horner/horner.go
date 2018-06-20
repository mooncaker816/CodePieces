package horner

import (
	"sync"
)

func horner(x float64, c ...float64) float64 {
	i := len(c) - 1
	y := c[i]
	for i > 0 {
		i--
		y = y*x + c[i]
	}
	return y
}

func hornerChan(x float64, c ...float64) float64 {
	y := c[0]
	ch := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(len(c) - 1)
	for i := 1; i < len(c); i++ {
		go func(i int, v float64) {
			for j := 0; j < i; j++ {
				v *= x
			}
			ch <- v
			wg.Done()
		}(i, c[i])
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		y += v
	}
	return y
}

type ab struct {
	a float64
	b float64
}

func hornerChan2(x float64, c ...float64) float64 {
	rightmost := make(chan ab)
	left := rightmost
	right := rightmost
	for i := 0; i < len(c); i++ {
		left = make(chan ab)
		go func(coeff float64, left, right chan ab) {
			term := <-left
			term.a = term.a * x
			term.b = coeff*term.a + term.b
			right <- term
		}(c[i], left, right)
		right = left
	}
	go func(ch chan ab) { ch <- ab{1, 0} }(left)
	return (<-rightmost).b
	// fmt.Println(y.b)
	// return y.b
}

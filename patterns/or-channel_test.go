package patterns

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	st := time.Now()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	<-or(
		randSig(r),
		randSig(r),
		randSig(r),
		randSig(r),
		randSig(r),
	)
	fmt.Printf("closed after %v seconds!\n", time.Now().Sub(st))
}

func randSig(r *rand.Rand) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		sec := time.Duration(r.Int63n(10) + 3)
		fmt.Printf("closing after %d seconds!\n", sec)
		time.Sleep(sec * time.Second)
	}()
	return ch
}

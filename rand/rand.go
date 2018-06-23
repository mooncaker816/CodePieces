package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	"sync"
)

func main() {

	var n int32
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	fmt.Println(n)
	ch := make(chan uint64)
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			result, _ := rand.Int(rand.Reader, big.NewInt(2))
			ch <- result.Uint64()
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

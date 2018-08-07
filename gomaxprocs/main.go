package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// [Min] 只是限制操作系统的线程数为1，所有 goroutine 会起在该线程上，但是 goroutine 之前还是并发
	// [Min] 也就是说还是无序的
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(2000)
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println("A: ", i) // 全是1000
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("B: ", i) // 无序的0-999
			wg.Done()
		}(i)
	}
	// [Min] A，B 之间也无序
	wg.Wait()
}

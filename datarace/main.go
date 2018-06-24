package main

import "fmt"

func main() {
	n := 0
	for i := 0; i < 10; i++ {
		go func() {
			n++
		}()
	}
	fmt.Println(n)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(nonNamedReturnValue())    // 0
	fmt.Println(namedReturnValue())       // 2
	fmt.Println(*nonNamedReturnPointer()) // 2
	exitBeforeDefer()
}

// [Min] defer 无法访问未命名的返回值
func nonNamedReturnValue() int {
	fmt.Println("nonNamedReturnValue():")
	var i int
	defer func() {
		i++
		fmt.Println(i) // 2
	}()
	defer func() {
		i++
		fmt.Println(i) // 1
	}()
	// [Min] 1. 首先返回值会在 return 语句的时候定义，并且将 i 的值拷贝到定义的返回值中
	// [Min] 2. 然后再检查 defer 并执行，修改的只是 i 并不是返回值本身
	// [Min] 3. 最后正式带着返回值结束函数，返回
	return i
}

// [Min] defer 可以访问命名返回值 i
func namedReturnValue() (i int) {
	fmt.Println("namedReturnValue():")
	defer func() {
		i++
		fmt.Println(i) // 2
	}()
	defer func() {
		i++
		fmt.Println(i) // 1
	}()
	// [Min] 返回值已经事先声明为 i，所以 defer 修改的就是返回值本身
	return i
}

// [Min] defer 修改的是 i 指向的内存块数据，也就是返回值指向的内存块数据，相当于修改了返回值
func nonNamedReturnPointer() *int {
	fmt.Println("nonNamedReturnPointer():")
	var i int
	defer func() {
		i++
		fmt.Println(i) // 2
	}()
	defer func() {
		i++
		fmt.Println(i) // 1
	}()
	// [Min] 未命名返回值与 i 指向同一地址，修改 i 就等于修改返回值
	return &i
}

// [Min] defer 之前如果有os.Exit，那么 defer 不会执行
func exitBeforeDefer() {
	fmt.Println("exitBeforeDefer():")
	defer func() {
		fmt.Println("in defer")
	}()

	os.Exit(1)
}

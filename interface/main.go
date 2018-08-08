package main

import "fmt"

func main() {
	var a Integer = 1
	var b Integer = 2

	// [Min] 不通过接口包装，指针接收者和值接受者可以互相混合调用

	// 值变量调用值接受者方法
	fmt.Println(a.Add1(b))
	// 值变量调用指针接受者方法
	fmt.Println(a.Add2(b))
	// 指针变量调用值接受者方法
	fmt.Println((&a).Add1(b))
	// 指针变量调用指针接受者方法
	fmt.Println((&a).Add2(b))

	// [Min] 经过接口封装后，指针接收者方法只可以被指针变量调用，
	// [Min] 而值接收者方法既可被指针变量调用，也可被值变量调用
	var i interface{} = a
	var j interface{} = &a
	fmt.Println(j.(*Integer).Add1(b)) // 指针调用值方法
	fmt.Println(i.(Integer).Add1(b))  // 值调用值方法
	fmt.Println(j.(*Integer).Add2(b)) // 指针调用指针方法
	//fmt.Println(i.(Integer).Add2(b))  // 值调用指针方法 - 编译报错
}

type Integer int

func (a Integer) Add1(b Integer) Integer {
	return a + b
}

func (a *Integer) Add2(b Integer) Integer {
	return *a + b
}

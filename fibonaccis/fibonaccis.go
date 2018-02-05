package fibo

import (
	"math/big"

	"github.com/mooncaker816/CodePieces/util/matrix"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func generateFibonaccis() func() uint64 {
	a, b := uint64(0), uint64(1)
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func fibo1(n uint64) {
	fib := generateFibonaccis()
	var i uint64
	for i = 0; i < n; i++ {
		fib()
	}
}

var fibarry = [3]uint64{0, 1, 0}

func fibo2(n uint64) uint64 {
	var i uint64
	for i = 2; i <= n; i++ {
		fibarry[2] = fibarry[0] + fibarry[1]
		fibarry[0] = fibarry[1]
		fibarry[1] = fibarry[2]
	}
	return fibarry[2]
}

//求矩阵的n次幂
func MatPow(a matrix.Matrix, b int) matrix.Matrix {
	arr0 := [4]*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)}
	s := matrix.New(2, 2, arr0[0:4])
	for b > 0 {
		if b&1 == 1 {
			s = *matrix.Multiply(s, a)
			b = b >> 1
		} else {
			b = b >> 1
		}
		a = *matrix.Multiply(a, a)
	}
	return s
}

//矩阵的N次幂与fib(1)和Fib(0)组成的2行1列的矩阵相乘求fib(n+1)和Fib(n)组成的2行1列的矩阵
//从fib(n+1)和Fib(n)的2行1列的矩阵中取出fib(n)
func fibo3(n int) *big.Int {
	arr0 := [6]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0), big.NewInt(2), big.NewInt(1)}
	k := matrix.New(2, 2, arr0[0:4])
	s := MatPow(k, n)
	d := matrix.New(2, 1, arr0[0:2])
	s = *matrix.Multiply(s, d)
	return s.Get(2, 1)
}

func fibo4(n uint) uint {
	if n < 2 {
		return n
	}
	return fibo4(n-2) + fibo4(n-1)
}

package main

import (
	"os"

	"github.com/k0kubun/pp"
)

func fib(n int) int {
	if n < 0 {
		pp.Printf("%s is unsupported number\n", n)
		os.Exit(1)
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	// 0
	pp.Println(fib(0))
	// 1
	pp.Println(fib(1))
	// 1
	pp.Println(fib(2))
	// 2
	pp.Println(fib(3))
	// 3
	pp.Println(fib(4))
	// 5
	pp.Println(fib(5))
	// 55
	pp.Println(fib(10))
	// -10 is unsupported number
	pp.Println(fib(-10))
}

package main

import (
	"os"
	"strconv"

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
	// `os.Args[0]` にはコマンド名が入る
	if len(os.Args) < 2 {
		pp.Println("Missing params!")
		os.Exit(1)
	}

	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		pp.Printf("%s\n", err)
		os.Exit(1)
	}

	pp.Println(fib(i))
}

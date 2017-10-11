package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	// https://github.com/golang/go/wiki/TableDrivenTests
	type Case struct {
		n, out int
	}
	cases := []Case{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}
	for i, c := range cases {
		if got := fib(c.n); got != c.out {
			t.Errorf("#%d: fib(%d) want %d, got %d\n", i, c.n, c.out, got)
		}
	}
}

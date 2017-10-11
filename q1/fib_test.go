package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	if fib(0) != 0 {
		t.Fatalf("Error fib(0)")
	}
	if fib(1) != 1 {
		t.Fatalf("Error fib(1)")
	}
	if fib(2) != 1 {
		t.Fatalf("Error fib(2)")
	}
	if fib(3) != 2 {
		t.Fatalf("Error fib(3)")
	}
	if fib(4) != 3 {
		t.Fatalf("Error fib(4)")
	}
	if fib(5) != 5 {
		t.Fatalf("Error fib(5)")
	}
	if fib(6) != 8 {
		t.Fatalf("Error fib(6)")
	}
	if fib(7) != 13 {
		t.Fatalf("Error fib(7)")
	}
	if fib(8) != 21 {
		t.Fatalf("Error fib(8)")
	}
	if fib(9) != 34 {
		t.Fatalf("Error fib(9)")
	}
	if fib(10) != 55 {
		t.Fatalf("Error fib(10)")
	}
}

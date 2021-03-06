package main

import "testing"

func TestFib(t *testing.T) {
	s := &Stack{limit: 2}

	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")

	v := s.Pop()
	if v != "dataC" {
		t.Error("Error!")
	}

	v = s.Pop()
	if v != "dataB" {
		t.Error("Error!")
	}

	v = s.Pop()
	if v != "" {
		t.Error("Error!")
	}

	s.Push("dataD")

	v = s.Pop()
	if v != "dataD" {
		t.Error("Error!")
	}

	v = s.Pop()
	if v != "" {
		t.Error("Error!")
	}

	v = s.Pop()
	if v != "" {
		t.Error("Error!")
	}
}

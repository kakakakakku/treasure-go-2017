package main

import (
	"github.com/k0kubun/pp"
)

type Stack struct {
	store []string
	limit int
}

func (s *Stack) Pop() string {
	if len(s.store) == 0 {
		return ""
	}
	v := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return v
}

func (s *Stack) Push(str string) {
	if len(s.store) == s.limit {
		s.store = s.store[1:]
	}
	s.store = append(s.store, str)
}

func main() {
	s := &Stack{limit: 2}

	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")

	// "dataC"
	pp.Println(s.Pop())

	// "dataB"
	pp.Println(s.Pop())

	// ""
	pp.Println(s.Pop())

	s.Push("dataD")

	// "dataD"
	pp.Println(s.Pop())

	// ""
	pp.Println(s.Pop())

	// ""
	pp.Println(s.Pop())
}

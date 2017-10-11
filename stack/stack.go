package main

import (
	"github.com/k0kubun/pp"
)

type Stack struct {
	Store []string
}

func (s *Stack) Pop() string {
	if len(s.Store) == 0 {
		return ""
	}
	v := s.Store[len(s.Store)-1]
	s.Store = s.Store[:len(s.Store)-1]
	return v
}

func (s *Stack) Push(str string) {
	s.Store = append(s.Store, str)
}

func main() {
	s := &Stack{}

	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")

	// "dataC"
	pp.Println(s.Pop())

	// "dataB"
	pp.Println(s.Pop())

	s.Push("dataD")

	// "dataD"
	pp.Println(s.Pop())

	// "dataA"
	pp.Println(s.Pop())

	// ""
	pp.Println(s.Pop())

	// ""
	pp.Println(s.Pop())
}

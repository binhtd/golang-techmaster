package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"
)

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{}, 0)
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func demoSet() {
	/*
		s := NewSet()
		s.Add("Peter")
		s.Add("David")

		fmt.Println(s.Contains("Peter"))
		fmt.Println(s.Contains("George"))

		s.Remove("David")
		fmt.Println(s.Contains("David"))
	*/

	s := sets.String{}

	s.Insert("a", "b")

	if s.Has("a") {
		fmt.Println("s.Has(a)")
	}

	s1 := sets.NewString("a", "b", "c")

	fmt.Println(s1)
	if !s1.Has("a") || !s1.Has("b") || !s1.Has("c") {
		fmt.Println("s1 has a, b, c")
	}

}

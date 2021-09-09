package main

import (
	"fmt"
	"sort"
)

func sortIntSlice() {
	intSlice := []int{10, 5, 25, 351, 14, 9}
	fmt.Println("Slice of integer BEFORE sort:", intSlice)
	sort.Ints(intSlice)
	fmt.Println("Slice of integer AFTER sort:", intSlice)
}

func sortReverse() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}

func sortSliceWithFunc() {
	people := []struct {
		name string
		age  int
	}{
		{"binh", 39},
		{"hung", 36},
		{"quynh", 41},
		{"minh", 10},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})
	fmt.Println("sort by name:", people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println("sort by age:", people)
}

func sortSliceStable() {
	people := []struct {
		name string
		age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75}, //--
		{"Bob", 75},
		{"Alice", 75}, //--
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].name < people[j].name
	})
	fmt.Println("sort by name:", people)
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println("sort by age:", people)

}

func search() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	element := 3
	index := sort.SearchInts(a, element)
	fmt.Printf("element %d found at %d", element, index)
	element = 9
	index = sort.SearchInts(a, element)
	fmt.Printf("element %d found at %d", element, index)
}

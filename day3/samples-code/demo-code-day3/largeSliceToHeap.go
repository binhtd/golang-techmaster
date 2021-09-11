package main

func LargeSliceToHeap() {
	LargeArray := make([]Person, 10000)

	for _, person := range LargeArray {
		person.Age = 0
		person.Name = ""
	}
}

func SmallSliceToStack() {
	smallArray := make([]Person, 1000)
	for _, person := range smallArray {
		person.Name = ""
		person.Age = 0
	}
}

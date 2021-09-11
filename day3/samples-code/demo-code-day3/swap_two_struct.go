package main

import "fmt"

func SwapPerson(a *(*Person), b *(*Person)) {
	fmt.Println("a = %p\n", a)
	fmt.Println("b = %p\n", b)
	a, b = b, a
	fmt.Println("a = %p\n", a)
	fmt.Println("b = %p\n", b)
	fmt.Println("alice ", *a)
	fmt.Print("bob ", *b)
}

func DemoSwapTwoStructs() {
	alice := Person{Name: "Alice", Age: 18}
	bob := Person{Name: "Bob", Age: 19}

	alice, bob = bob, alice
	fmt.Println("alice ", alice)
	fmt.Println("bob ", bob)
}

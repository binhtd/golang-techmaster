package main

import "fmt"

func (p *Person) increaseAge() {
	p.Age += 1
}

func (p *Person) reverseName() (result string) {
	for _, s := range p.Name {
		result = string(s) + result
	}
	return
}

func makeAPerson(name string, age int) (person *Person) {
	person = new(Person)
	person.Name = name
	person.Age = age
	return
}

func WillReceiverEscapeToHeap() {
	alice := new(Person)
	alice.Name = "Alice"
	alice.Age = 21
	alice.increaseAge()

	temp := alice.reverseName()
	fmt.Println(temp)

	bob := makeAPerson("Bob", 22)
	bob.reverseName()
}

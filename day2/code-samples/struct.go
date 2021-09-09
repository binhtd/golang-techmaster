package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address1  Address
	Address2  *Address
}

type Address struct {
	Country string
	City    string
}

func (p *Person) FullName() string {
	fmt.Printf("pointer receiver %p\n", p)
	p.Age = 100
	return p.FirstName + " " + p.LastName
}

func (p Person) String() string {
	fmt.Printf("value receiver: %p\n", &p)
	p.Age = 200
	return fmt.Sprintf("%v is %v years old ", p.FirstName, p.Age)
}

func NewAPerson(firstName string, lastName string, age int) *Person {
	if age < 0 {
		return nil
	}
	p := new(Person)
	p.FirstName = firstName
	p.LastName = lastName
	p.Age = age
	return p
}

func BuildPerson() *Person {
	return new(Person)
}

func (p *Person) WithFirstName(firstName string) *Person {
	p.FirstName = firstName
	return p
}

func (p *Person) WithLastName(lastName string) *Person {
	p.LastName = lastName
	return p
}

func (p *Person) WithAge(age int) *Person {
	p.Age = age
	return p
}

func demoStruct() {
	person := Person{"Trinh", "Cuong", 45,
		Address{"Vietnam", "Hanoi"},
		&Address{"USA", "California"}}

	fmt.Printf("address pointer:%p\n", &person)
	fmt.Println(person.FullName())
	fmt.Println(person.String())

	tom := NewAPerson("Le", "lan anh", -2)
	if tom != nil {
		fmt.Println(tom.FullName())
	}
	jack := BuildPerson().WithFirstName("Jack").WithLastName("London").WithAge(12)
	fmt.Println(jack)
}

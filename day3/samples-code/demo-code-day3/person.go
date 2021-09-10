package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address1  Address
}

type Address struct {
	Country string
	City    string
}

//init function
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

func (p *Person) WithLastNem(lastName string) *Person {
	p.LastName = lastName
	return p
}

func (p *Person) WithAge(age int) *Person {
	p.Age = age
	return p
}

func demoStruct() {
	person := Person{"le", "nam", 32, Address{"vietnam", "hanoi"}}
	fmt.Println(person)

	p1 := BuildPerson().WithFirstName("da").WithLastNem("la").WithAge(52)
	fmt.Println(p1)
}

package main

import "fmt"

type Person1 struct {
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
func NewAPerson(firstName string, lastName string, age int) *Person1 {
	if age < 0 {
		return nil
	}
	p := new(Person1)
	p.FirstName = firstName
	p.LastName = lastName
	p.Age = age
	return p
}

func BuildPerson() *Person1 {
	return new(Person1)
}

func (p *Person1) WithFirstName(firstName string) *Person1 {
	p.FirstName = firstName
	return p
}

func (p *Person1) WithLastNem(lastName string) *Person1 {
	p.LastName = lastName
	return p
}

func (p *Person1) WithAge(age int) *Person1 {
	p.Age = age
	return p
}

func demoStruct() {
	person := Person1{"le", "nam", 32, Address{"vietnam", "hanoi"}}
	fmt.Println(person)

	p1 := BuildPerson().WithFirstName("da").WithLastNem("la").WithAge(52)
	fmt.Println(p1)
}

package main

import "fmt"

func (p *Person) string() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func (p *Person) IncreaseAge() {
	p.Age += 1
}

func (p Person) NotReallyIncreaseAge() {
	p.Age += 1
}

func EncryptPersonNamePointerArg(p *Person) {
	var result string
	for _, v := range p.Name {
		result = string(v) + result
	}
	p.Name = result
}

func EncryptPersonNameValueArg(p Person) {
	var result string
	for _, v := range p.Name {
		result = string(v) + result
	}
	p.Name = result
}

func DemoPointerStruct() {
	bob := &Person{"Bob", 46}
	bob.NotReallyIncreaseAge()
	fmt.Println(bob)

	tom := new(Person)
	tom.Name = "Tom"
	tom.Age = 27
	tom.IncreaseAge()
	tom.NotReallyIncreaseAge()
	fmt.Println(tom)

	alice := Person{Name: "alice", Age: 18}
	alice.IncreaseAge()
	fmt.Println(&alice)

	EncryptPersonNamePointerArg(tom)
	fmt.Println(tom)

	EncryptPersonNameValueArg(alice)
	fmt.Println(&alice)
}

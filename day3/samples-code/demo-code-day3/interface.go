package main

import (
	"fmt"
)

type Animal interface {
	Move()
	Sound()
}

type Cat struct {
	Name string
}

func (c *Cat) Move() {
	fmt.Println(c.Name, " move")
}

func (c *Cat) Sound() {
	fmt.Println(c.Name, " Sound")
}

type Bird struct {
	Name string
}

func (b *Bird) Move() {
	fmt.Println(b.Name, " Move")
}

func (b *Bird) Sound() {
	fmt.Println(b.Name, " Sound")
}

func demoInterface() {
	animal := make([]Animal, 2)
	animal[0] = &Bird{Name: "Bird"}
	animal[1] = &Cat{Name: "Cat"}

	for _, animal := range animal {
		animal.Move()
		animal.Sound()
	}
}

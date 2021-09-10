package main

import "fmt"

type Product struct {
	id    string
	name  string
	price int
}

func (product Product) increasePrice1(percentage int) {
	fmt.Printf("%p\n", &product)
	product.price = product.price * (100 + percentage) / 100
}
func (product *Product) increasePrice2(percentage int) {
	fmt.Printf("%p\n", product)
	product.price = product.price * (100 + percentage) / 100
}

func (product Product) incPrice1(percentage int) {
	product.price = product.price * (100 + percentage) / 100
}

func (product *Product) incPrice2(percentage int) {
	product.price = product.price * (100 + percentage) / 100
}

func demoProduct() {
	nike := Product{
		id:    "abc",
		name:  "name prod",
		price: 100,
	}
	fmt.Printf("%p\n", &nike)
	nike.increasePrice1(20)
	fmt.Println(nike.price)

	nike.increasePrice2(20)
	fmt.Println(nike.price)
}

package basic

import (
	"fmt"
)

type Opertor2operand = func(int, int) int
type AnyType = interface{}

type Product struct {
	Id    string
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("Id: %v - Name: %v - Price: %v", p.Id, p.Name, p.Price)
}

type Point struct{ X, Y float64 }
type Color struct{ R, G, B int }
type ColoredPoint struct {
	Point
	color Color
}

func (cp ColoredPoint) String() string {
	return fmt.Sprintf("Point at %.2f - %.2f color {%v, %v, %v}", cp.X, cp.Y, cp.color.R, cp.color.G, cp.color.B)
}

func DemoStringerInterface() {
	product := Product{
		Id:    "ax123",
		Name:  "Sony walkman",
		Price: 52343,
	}
	fmt.Println(product)
	arrayAnyType := []interface{}{
		3.14, "Hello", true, product, Add, []string{"John", "Anna", "Tom"},
	}
	fmt.Println(arrayAnyType)
	cp := ColoredPoint{
		Point: Point{X: 10.5, Y: 20.1},
		color: Color{100, 255, 255},
	}
	fmt.Println(cp)
}

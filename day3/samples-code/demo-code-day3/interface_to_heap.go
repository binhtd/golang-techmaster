package main

import (
	"fmt"
	"reflect"
)

func GetType(a interface{}) string {
	return reflect.TypeOf(a).String()
}

func InterfaceArgumentEscapeToHeap() {
	tom := Person{Name: "Tom", Age: 18}
	if GetType(tom) == "*main.Person" {
		fmt.Println("tom is person")
	}
}

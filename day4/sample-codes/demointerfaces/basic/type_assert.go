package basic

import "fmt"

func DemoInterfaceTypeAssert() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/* Ví dụ này sẽ gây lỗi
	f = i.(float64) // panic
	fmt.Println(f)
	*/

	if f, ok := i.(string); ok {
		fmt.Println(f)
	}

	person := Person{
		Id:   "OX-13",
		Name: "Donald Trump",
	}

	i = person

	if f, ok := i.(Person); ok {
		fmt.Println(f.Id)
	}
}

func DemoSwitchType() {
	person := Person{
		Id:   "123",
		Name: "Demo",
	}

	arrayAnyType := []interface{}{
		1, 3.14, "Hello", true, person, Add, []string{"a", "b", "c"},
	}
	for _, v := range arrayAnyType {
		switch v.(type) {
		case int:
			fmt.Printf("%v is int\n", v)
		case string:
			fmt.Printf("%v is string\n", v)
		case Person:
			fmt.Printf("%v is Person\n", v)
		case func(int, int) int:
			fmt.Printf("%v is func(int,int) int\n", v)
		}
	}
}

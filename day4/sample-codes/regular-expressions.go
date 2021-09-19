package main

import (
	"fmt"
	"regexp"
)

func demoRegularExpression() {
	match, _ := regexp.MatchString("p[a-z]ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p[a-z]+ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))

}

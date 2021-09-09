package main

import (
	"fmt"
)

func demoSlice() {
	a := []string{"a", "b", "c", "d"}
	a = append(a, "f")
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	fmt.Println(a[:2])
	fmt.Println(a[2:])
	fmt.Println(a[len(a)-2:])
	fmt.Println(a[1:3])
}

func removeItemSliceNotKeepOrder(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}

func removeItemSliceNotKeepOrder2(a []string, i int) []string {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func removeItemSliceKeepOrder(a []string, i int) []string {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}

func removeItemSliceKeepOrder2(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

func reserveNotKeepOriginSlice(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func reserveKeepOriginSlice(a []string) (reserved []string) {
	reserved = []string{}
	for i := 0; i < len(a); i++ {
		reserved = append(reserved, a[len(a)-i-1])
	}
	return reserved
}

func removeDuplicate(a []string) (result []string) {
	keys := map[string]bool{}
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return
}

func reserveSliceAnyType(a []interface{}) (reversed []interface{}) {
	for i := range a {
		n := a[len(a)-1-i]
		reversed = append(reversed, n)
	}
	return
}

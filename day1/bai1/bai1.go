package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const Delta_Less_Than_Zero_Mesage = "Delta is less than zero"

func main() {
	a, err := readNumberFromKeyboard("input a:")
	checkError(err)

	b, err := readNumberFromKeyboard("input b:")
	checkError(err)

	c, err := readNumberFromKeyboard("input a:")
	checkError(err)

	x1, x2 := QuadraticEquation(a, b, c)
	fmt.Printf("x1=%f, x2=%f", x1, x2)
}

/*
read number from keyboard
*/
func readNumberFromKeyboard(msg string) (result float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")

	if result, err = strconv.ParseFloat(str, 64); err != nil {
		return 0.0, err
	}

	return result, nil
}

/*
check if object error is different than nil return panic error
*/
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
get the x1, x2 from quadratic equation
*/
func QuadraticEquation(a, b, c float64) (x1, x2 float64) {
	if b*b-4*a*c < 0 {
		panic(errors.New(Delta_Less_Than_Zero_Mesage))
	}

	return (-b + math.Sqrt(b*b-4*a*c)) / 2 * a, (-b - math.Sqrt(b*b-4*a*c)) / 2 * a
}

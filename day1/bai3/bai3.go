package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	number, err := readNumberFromKeyboard("number N is ")
	checkError(err)

	if number > 100000 {
		fmt.Println("too large number")
		os.Exit(0)
	}

	for i := 2; i <= number; i++ {
		if IsPrime(i) {
			fmt.Print(i, " ")
		}
	}
}

/*
read number from keyboard
*/
func readNumberFromKeyboard(msg string) (result int, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")

	if result, err = strconv.Atoi(str); err != nil {
		return 0, err
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
IsPrime
*/
func IsPrime(number int) bool {
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	numberComputerThink := rand.Intn(101)

	for {
		numberYouThink, err := readNumberFromKeyboard("your number is ")
		checkError(err)

		if numberYouThink == numberComputerThink {
			fmt.Println("well done")
			break
		}

		if numberYouThink < numberComputerThink {
			fmt.Println("your number is less than computer's number")
		} else {
			fmt.Println("your number is greater than computer's number")
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

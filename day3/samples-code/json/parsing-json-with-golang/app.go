package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println(result["users"])
}

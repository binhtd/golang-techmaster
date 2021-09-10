package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
"name":"Dee Leng",
"email":"dleng0@cocolog-nifty.com",
"job":"developer",
"gender":"Female",
"city":"London",
"salary":9662,
"birthdate":"2007-09-30" */
type Person struct {
	Name     string `json:name`
	Email    string `json:email`
	Job      string `json:job`
	Gender   string `json:gender`
	City     string `json:city`
	Salary   int    `json:salary`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	jsonFile, err := os.ReadFile("person.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened user.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var people []Person
	json.Unmarshal(byteValue, &people)
}

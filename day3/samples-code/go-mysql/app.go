package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Demo some action with mysql in golang")
	//demoInsert()
	//demoFetch()

	demoSelectRow()
}

func demoInsert() {
	db, err := sql.Open("mysql", "root:D@ngha123456@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO demo(id, name) values (1, 'le')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func demoFetch() {
	db, err := sql.Open("mysql", "root:D@ngha123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	result, err := db.Query("select id, name from tag")

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var tag Tag
		err = result.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(tag.Name)
	}
}

func demoSelectRow() {
	var tag Tag

	db, err := sql.Open("mysql", "root:D@ngha123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	err = db.QueryRow("select id, name from tag where id= ?", 2).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error())
	}

	log.Println(tag.ID)
	log.Println(tag.Name)

}

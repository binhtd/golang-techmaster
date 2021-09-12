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

	//demoSelectRow()
	//demoConnect()
	//demoTechmasterInsert()

	demoTechmasterPrepareStatement()
}

//https://medium.com/hackernoon/today-i-learned-storing-emoji-to-mysql-with-golang-204a093454b7
func demoTechmasterPrepareStatement() {
	categories := map[string]([]string){
		"điện thoại, máy tính bảng": []string{"smart phone", "điện thoại phổ thông", "máy tính bảng", "máy đọc sách"},
		"điện tử, điện lạnh": []string{"tivi", "dàn âm thanh", "tủ lạnh - tủ mát", "máy điều hoà", "phụ kiện điện lạnh", "máy giặt", "hút bụi", "lọc không khí", "rủa bát"},
		"máy tính, laptop": []string{"desktop", "server", "laptop", "phụ kiện máy tính"},
		"camera, camcoder": []string{"máy ảnh", "máy ảnh kỹ thuật số", "máy quay phim", "camera giám sá", "camera hành trình", "balo", "ống kính"},
		"đồ bếp":           []string{"lò vi sóng", "máy say sinh tố", "máy đánh trứng", "bếp từ", "bếp hồng ngoại", "bếp ga"},
	}
	db, err := sql.Open("mysql", "root:D@ngha123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	statement, err := db.Prepare("INSERT INTO category (name, parent_id) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()

	for key, subcategories := range categories {
		result, err := statement.Exec(key, nil)
		if err != nil {
			panic(err.Error())
		}
		id, err := result.LastInsertId()

		for _, subcat := range subcategories {
			_, err = statement.Exec(subcat, id)
		}
		if err != nil {
			panic(err.Error())
		}
	}
}

func demoTechmasterInsert() {
	db, err := sql.Open("mysql", "root:D@ngha123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	countries := map[string]string{
		"VN": "Viet nam",
		"CN": "China",
		"US": "USA",
		"DE": "Germanny",
		"SG": "Singapore",
		"ES": "Spain",
		"KR": "South Korea",
		"JP": "Japan",
	}
	var insert *sql.Rows
	for country_code, country_name := range countries {
		insert, err = db.Query("INSERT INTO country (code, name) VALUES ('" + country_code + "' ,'" + country_name + "')")
	}
	defer insert.Close()
}

func demoTechmasterConnect() {
	db, err := sql.Open("mysql", "root:D@ngha123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
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

package router

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"path"
)

type student struct {
	Id   int    `json:"student_id"`
	Name string `json:"full_name"`
}

func returnPlainText(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Write([]byte("Post method"))
	} else if r.Method == http.MethodPut {
		w.Write([]byte("Put method"))
	}
}

func returnJson(w http.ResponseWriter, r *http.Request) {
	students := []student{
		{Id: 1, Name: "Nguyễn Văn A"},
		{Id: 2, Name: "Trịnh Văn B"},
		{Id: 3, Name: "Ngô Thị C"},
	}
	studentsJson, err := json.Marshal(students)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(studentsJson)
}

func returnXml(w http.ResponseWriter, r *http.Request) {
	students := []student{
		{Id: 1, Name: "Nguyễn Văn A"},
		{Id: 2, Name: "Trịnh Văn B"},
		{Id: 3, Name: "Ngô Thị C"},
	}
	studentsXml, err := xml.Marshal(students)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(studentsXml)
}

func returnHtml(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta http-equiv="X-UA-Compatible" content="IE=edge">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Books</title>
				<link rel="stylesheet" href="http://localhost:3000/static/css/style.css">
			</head>
			<body>
				<h1>Những cuốn sách được yêu thích</h1>
				<ul>
					<li>Những người khốn khổ</li>
					<li>Đắc nhân tâm</li>
				</ul>
				<img src="http://localhost:3000/static/image/nature.jpg">
				<script src="http://localhost:3000/static/js/main.js"></script>
			</body>
		</html>
	`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, html)
}

func returnFile(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("static/image", "nature.jpg")
	http.ServeFile(w, r, filePath)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("OK"))
}

func DemoRouter() {
	http.HandleFunc("/", returnPlainText)
	http.HandleFunc("/json", returnJson)
	http.HandleFunc("/xml", returnXml)
	http.HandleFunc("/file", returnFile)
	http.HandleFunc("/html", returnHtml)
	http.HandleFunc("/health", HealthCheckHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Fatal(http.ListenAndServe(":3000", nil))
}

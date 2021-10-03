package upload_file

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)

	fmt.Println("path", r.URL.Path)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Jenny!")
}

type Account struct {
	UserName string
	Password string
}

var accounts = []Account{
	{UserName: "jenny", Password: "111"},
	{UserName: "bob", Password: "222"},
	{UserName: "foo", Password: "333"},
	{UserName: "bar", Password: "444"},
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./view/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["passsword"])

		username := strings.Join(r.Form["username"], "")
		password := strings.Join(r.Form["password"], "")

		for _, acc := range accounts {
			if username == acc.UserName && password == acc.Password {
				fmt.Fprintln(w, "dang nhap thanh cong")
				return
			}
		}
		fmt.Fprintln(w, "dang nhap sai, thu lai")
	}
}

func DemoForm() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	fmt.Println(http.ListenAndServe(":3000", nil))
}

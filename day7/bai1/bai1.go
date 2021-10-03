package bai1

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Id       int
	FullName string
	Email    string
	Phone    string
	Age      int
	Sex      string
}

var users = []User{
	{Id: 1, FullName: "user a", Email: "user-a@gmail.com", Phone: "09234324", Age: 32, Sex: "Male"},
	{Id: 2, FullName: "user b", Email: "user-b@gmail.com", Phone: "09234325", Age: 22, Sex: "Female"},
	{Id: 3, FullName: "user c", Email: "user-c@gmail.com", Phone: "09234326", Age: 35, Sex: "Male"},
	{Id: 4, FullName: "user d", Email: "user-d@gmail.com", Phone: "09234327", Age: 30, Sex: "Female"},
	{Id: 5, FullName: "user e", Email: "user-e@gmail.com", Phone: "09234328", Age: 39, Sex: "Male"},
	{Id: 6, FullName: "user f", Email: "user-f@gmail.com", Phone: "09234329", Age: 50, Sex: "Female"},
}

func usersHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		id, ok := query["id"]
		if ok {
			filterById, _ := strconv.Atoi(strings.Join(id, ","))
			user := filterByUserId(users, filterById)

			usersJson, err := json.Marshal(user)
			if err != nil {
				panic(err.Error())
			}
			w.Write(usersJson)
			return
		}

		usersJson, err := json.Marshal(users)
		if err != nil {
			panic(err.Error())
		}
		w.Write(usersJson)

		return
	case "POST":
		w.WriteHeader(http.StatusCreated)
		var user User

		id, ok := query["id"]
		if ok {
			user.Id, _ = strconv.Atoi(strings.Join(id, ","))
		}

		fullname, ok := query["fullname"]
		if ok {
			user.FullName = strings.Join(fullname, "")
		}

		email, ok := query["email"]
		if ok {
			user.Email = strings.Join(email, "")
		}

		phone, ok := query["phone"]
		if ok {
			user.Phone = strings.Join(phone, "")
		}

		age, ok := query["age"]
		if ok {
			user.Age, _ = strconv.Atoi(strings.Join(age, ""))
		}

		sex, ok := query["sex"]
		if ok {
			user.Sex = strings.Join(sex, "")
		}

		usersJson, err := json.Marshal(user)
		if err != nil {
			panic(err.Error())
		}
		w.Write(usersJson)
		return
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		var user User

		id, ok := query["id"]
		if !ok {
			panic("id field must input")
		}
		filterById, _ := strconv.Atoi(strings.Join(id, ","))
		user = filterByUserId(users, filterById)
		fullname, ok := query["fullname"]
		if ok {
			user.FullName = strings.Join(fullname, "")
		}

		email, ok := query["email"]
		if ok {
			user.Email = strings.Join(email, "")
		}

		phone, ok := query["phone"]
		if ok {
			user.Phone = strings.Join(phone, "")
		}

		age, ok := query["age"]
		if ok {
			user.Age, _ = strconv.Atoi(strings.Join(age, ""))
		}

		sex, ok := query["sex"]
		if ok {
			user.Sex = strings.Join(sex, "")
		}

		usersJson, err := json.Marshal(user)
		if err != nil {
			panic(err.Error())
		}
		w.Write(usersJson)
		return
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		id, ok := query["id"]
		if !ok {
			panic("id field must input")
		}

		filterById, _ := strconv.Atoi(strings.Join(id, ","))
		users = deleteUserById(users, filterById)

		usersJson, err := json.Marshal(users)
		if err != nil {
			panic(err.Error())
		}
		w.Write(usersJson)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func filterByUserId(users []User, userId int) (result User) {
	for _, user := range users {
		if user.Id == userId {
			result = user
			break
		}
	}
	return result
}

func deleteUserById(users []User, userId int) []User {
	for index, user := range users {
		if user.Id == userId {
			users[index] = users[len(users)-1]
			users[len(users)-1] = User{}
			return users[:len(users)-1]
			break
		}
	}
	return users
}
func Bai1Demo() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandle)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

package server

import (
	"fmt"
	"net/http"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Home page"))
	} else if r.URL.Path == "/about" {
		w.Write([]byte("About page"))
	} else {
		w.Write([]byte("Page not found"))
	}
}

func DemmoServerDefault() {
	http.HandleFunc("/", homeHandle)
	fmt.Println("Server listening on port 3000 ...")
	fmt.Println(http.ListenAndServe(":3000", nil))
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}

func DemoServerServerMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", aboutPage)
	fmt.Println("Server listenning on port 3000 ...")
	fmt.Println(http.ListenAndServe(":3000", mux))
}

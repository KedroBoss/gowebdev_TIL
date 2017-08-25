package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Home Page")
	})
	http.HandleFunc("/dog/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Dog page")
	})
	http.HandleFunc("/me/", func(w http.ResponseWriter, req *http.Request) {
		name := "Jeff"
		io.WriteString(w, "My name is "+name)
	})

	http.ListenAndServe(":8080", nil)
}

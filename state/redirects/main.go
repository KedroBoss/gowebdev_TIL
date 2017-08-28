package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

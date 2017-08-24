package main

import (
	"fmt"
)

import (
	"net/http"
)

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Something-something")
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}

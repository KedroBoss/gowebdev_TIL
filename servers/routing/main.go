package main

import (
	"io"
	"net/http"
)

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/handle":
		io.WriteString(w, "This is a handle. Handle of what? Stuff.")
	case "/hand":
		io.WriteString(w, "Hand of a person.")
	}
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}

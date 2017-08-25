package main

import (
	"io"
	"net/http"
)

type handHandler int

func (h handHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Handy hand")
}

type handleHandler int

func (h handleHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Handy handle is handled")
}
func main() {
	var hand handHandler
	var handle handleHandler

	mux := http.NewServeMux()
	mux.Handle("/hand/", hand)
	mux.Handle("/handle", handle)
	http.ListenAndServe(":8080", mux)
}

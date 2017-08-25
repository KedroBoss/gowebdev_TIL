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

	// /hand/moreURL will be handled by the Handle
	mux.Handle("/hand/", hand)

	// /handle only be handled
	mux.Handle("/handle", handle)
	http.ListenAndServe(":8080", mux)
}

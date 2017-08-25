package main

import (
	"io"
	"net/http"
)

func hand(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Handy hand")
}

func handle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Handy handle is handled")
}
func main() {
	// /hand/moreURL will be handled by the Handle
	http.HandleFunc("/hand/", hand)

	// /handle only be handled
	http.HandleFunc("/handle", handle)

	http.HandleFunc("/func/in/func", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "This is totaly legit way of doing it.")
	})

	// nil = use DefaultServeMux
	http.ListenAndServe(":8080", nil)
}

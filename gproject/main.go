package main

import (
	"net/http"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("favicon.ico", http.NotFoundHandler())
}

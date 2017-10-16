package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		// Before exeuting handler
		log.Printf("Started %s %s", req.Method, req.URL.Path)
		next.ServeHTTP(w, req)
		// After executing handler
		log.Printf("Finished %s in %v", req.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, req *http.Request) {
	log.Println("Executing Index")
	fmt.Fprintln(w, "Hello!")
}
func main() {
	// Convert index into HandlerFunc
	http.Handle("/", loggingHandler(http.HandlerFunc(index)))
	log.Println("Listening")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"io"
	"net/http"
	"strconv"
)

// MaxAge sets the expiration time for a cookie
// To delete is: use MaxAge: -1

func increment(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:   "counter",
			Value:  "0",
			MaxAge: 15,
		}
	}
	tmp, _ := strconv.Atoi(c.Value)
	tmp++
	c.Value = strconv.Itoa(tmp)

	http.SetCookie(w, c)

	io.WriteString(w, c.Value)
}

func main() {
	http.HandleFunc("/", increment)
	http.ListenAndServe(":8080", nil)
}

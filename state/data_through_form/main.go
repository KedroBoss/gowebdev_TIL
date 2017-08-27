package main

import (
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=urf-8")
	io.WriteString(w,
		`			<form method="post">
						<input type="text" name="q">
						<input type="submit">
					</form>
					<br>`+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

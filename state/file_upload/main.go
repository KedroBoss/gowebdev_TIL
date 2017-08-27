package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		// open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nFile:", f, "\nHeader:", h, "\nErr:", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		dst, err := os.Create(filepath.Join(".", h.Filename) + "-copy")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
	w.Header().Set("Content-Type", "text/html; charset=urf-8")
	io.WriteString(w,
		`			<form method="post" enctype="multipart/form-data">
						<input type="file" name="q">
						<input type="submit">
					</form>
					<br>`+s)

}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

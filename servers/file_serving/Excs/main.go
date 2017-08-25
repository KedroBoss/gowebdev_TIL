package main

import (
	"html/template"
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dwarf(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("../assets/dwarf.html"))
	tpl.ExecuteTemplate(w, "dwarf.html", nil)
}

func dwarfImg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../assets/dwarf.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dwarf/", dwarf)
	http.HandleFunc("/dwarf.jpg", dwarfImg)
	http.ListenAndServe(":8080", nil)
}

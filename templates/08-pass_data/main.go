package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("tmp.html"))
}

func main() {
	err := tmp.Execute(os.Stdout, 42) // passing 42 to the template
	if err != nil {
		log.Fatal(err)
	}
}

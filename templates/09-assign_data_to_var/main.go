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
	// passing string to the template
	err := tmp.Execute(os.Stdout, `This is different type`)
	if err != nil {
		log.Fatal(err)
	}
}

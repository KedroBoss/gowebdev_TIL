package main

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

// To parse template once use init method
func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))
}

// Simple example of using custom variables in html
func main() {
	// tpm is "container" holding all the parsed templates
	// if Execute multiple templates, it will execute the first one
	err := tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmp.ExecuteTemplate(os.Stdout, "tmp1.html", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

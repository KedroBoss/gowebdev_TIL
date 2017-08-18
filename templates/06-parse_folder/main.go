package main

import (
	"log"
	"os"
	"text/template"
)

// Simple example of using custom variables in html
func main() {
	// tpm is "container" holding all the parsed templates
	// if Execute multiple templates, it will execute the first one
	tmp, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmp.ExecuteTemplate(os.Stdout, "tmp.html", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

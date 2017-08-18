package main

import (
	"log"
	"os"
	"text/template"
)

// Simple example of using custom variables in html
func main() {
	// tpm is "container" holding all the parsed templates
	tmp, err := template.ParseFiles("tmp.html")
	if err != nil {
		log.Fatalln("Error parsing template", err)
	}
	err = tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

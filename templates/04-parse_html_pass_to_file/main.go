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
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer nf.Close()

	err = tmp.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

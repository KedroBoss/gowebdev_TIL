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
	tmp, err := template.ParseFiles("tmp.html")
	if err != nil {
		log.Fatalln("Error parsing template", err)
	}
	err = tmp.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
	}

	tmp, err = tmp.ParseFiles("tmp.html", "tmp1.html")
	err = tmp.ExecuteTemplate(os.Stdout, "tmp1.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

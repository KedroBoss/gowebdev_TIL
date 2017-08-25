package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

// TemplateFolder is location for storing templates
var TemplateFolder = "templates/"

// HandleTemplate handles a template
// Name of a template, writer, context, *template.Template
func HandleTemplate(templateName string, w io.Writer, context interface{}, tpl *template.Template) {
	// parse base and nested templates
	tpl = template.Must(template.ParseFiles(TemplateFolder+"base.html", TemplateFolder+templateName))

	// execute base template
	err := tpl.ExecuteTemplate(w, "base.html", context)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/base.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// io.WriteString(w, "Home Page")
		HandleTemplate("index.html", io.Writer(w), nil, tpl)
	})
	http.HandleFunc("/dog/", func(w http.ResponseWriter, req *http.Request) {
		// io.WriteString(w, "Dog page")
		HandleTemplate("dog.html", io.Writer(w), nil, tpl)
	})
	http.HandleFunc("/me/", func(w http.ResponseWriter, req *http.Request) {
		name := "Jeff"
		// io.WriteString(w, "My name is "+name)
		HandleTemplate("me.html", w, name, tpl)
	})

	http.ListenAndServe(":8080", nil)
}

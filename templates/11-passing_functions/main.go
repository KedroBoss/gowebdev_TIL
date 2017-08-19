package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name string
}

var tmp *template.Template

// uc - pass predefined function to template
// ft - pass custom function to template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	// ParseFiles returns: *Template, Error
	// Funcs returns: *Template
	// New returns: *Template
	// Must takes: *Template, Error
	// Must(*Template.*Template.*Template, Error)
	tmp = template.Must(template.New("").Funcs(fm).ParseFiles("tmp.html"))
}

func main() {
	p := Person{"Boss"}
	err := tmp.ExecuteTemplate(os.Stdout, "tmp.html", p)
	if err != nil {
		log.Fatal(err)
	}
}

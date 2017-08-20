package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmp *template.Template

var fm = template.FuncMap{
	"mdy": monthDayYear,
}

func monthDayYear(t time.Time) string {
	// The numbers matter
	// 02 - day, 01 - date, 06 - year
	return t.Format("01 - 02 - 2006")
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
	err := tmp.ExecuteTemplate(os.Stdout, "tmp.html", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}

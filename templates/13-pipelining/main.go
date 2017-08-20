package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tmp *template.Template

var fm = template.FuncMap{
	"double": double,
	"square": square,
	"sqRoot": sqRoot,
}

func double(x float64) float64 {
	return float64(x + x)
}

func square(x float64) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(float64(x))
}

func init() {
	// ParseFiles returns: *Template, Error
	// Funcs returns: *Template
	// New returns: *Template
	// Must takes: *Template, Error
	// Must(*Template.*Template.*Template, Error)
	tmp = template.Must(template.New("").Funcs(fm).ParseFiles("tmp.html"))
}

// Pipelining - passing value of one function to another
func main() {
	err := tmp.ExecuteTemplate(os.Stdout, "tmp.html", 5.0)
	if err != nil {
		log.Fatal(err)
	}
}

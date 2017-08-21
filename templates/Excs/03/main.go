package main

import (
	"log"
	"os"
	"text/template"
)

type Menu struct {
	BrMenu Breakfast
	LnMenu Lunch
	DnMenu Dinner
}
type Breakfast struct {
	Meals []meal
}
type Lunch struct {
	Meals []meal
}
type Dinner struct {
	Meals []meal
}

type meal struct {
	Name  string
	Price float64
}

var tpl *template.Template
var meals = []meal{
	meal{"Meal1", 69.99},
	meal{"Meal2", 99.69},
	meal{"Meal3", 13.37},
	meal{"Meal4", 05.99},
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

// Create a data structure to pass to a template which
// contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
func main() {
	br := Breakfast{meals}
	ln := Lunch{meals}
	dn := Dinner{meals}

	menu := Menu{br, ln, dn}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", menu)
	if err != nil {
		log.Fatal(err)
	}
}

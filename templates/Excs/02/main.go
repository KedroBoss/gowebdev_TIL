package main

import (
	"log"
	"os"
	"text/template"
)

// Regions to have
var Regions = []string{"Southern", "Central", "Northen"}

type hotel struct {
	Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

// Create a data structure to pass to a template which
// contains information about California hotels including Name, Address, City, Zip, Region
// region can be: Southern, Central, Northern
// can hold an unlimited number of hotels
func main() {
	hotels := []hotel{
		hotel{
			Name:    "Boss Hotel",
			Address: "Boss Street",
			City:    "England",
			Zip:     "691337",
			Region:  Regions[0],
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", hotels)
	if err != nil {
		log.Fatal(err)
	}
}

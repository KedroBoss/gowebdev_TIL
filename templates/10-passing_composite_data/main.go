package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("range.html", "struct.html"))
}

func main() {

	{ // Passing slice
		listOfPeople := []string{"Boss1", "Boss2", "Boss3"}

		err := tmp.ExecuteTemplate(os.Stdout, "range.html", listOfPeople)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("**************")
	{ // Passing map
		mapOfPeople := map[string]int{
			"Boss1": 1,
			"Boss2": 2,
			"Boss3": 3,
		}
		err := tmp.ExecuteTemplate(os.Stdout, "range.html", mapOfPeople)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("**************")
	{ // Passing struct, must be exported
		boss1 := person{
			"Boss1",
			42,
		}
		err := tmp.ExecuteTemplate(os.Stdout, "struct.html", boss1)
		if err != nil {
			log.Fatal(err)
		}

	}

}

type person struct {
	Name string
	Age  int
}

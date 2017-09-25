package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

var db *sql.DB
var err error

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test_schema?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		name := req.FormValue("cName")
		stmt, err := db.Prepare(`INSERT INTO customer SET cName=?`)
		if err != nil {
			http.Redirect(w, req, "/", http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}

		r, err := stmt.Exec(name)
		if err != nil {
			http.Redirect(w, req, "/", http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		n, err := r.RowsAffected()
		if err != nil {
			http.Redirect(w, req, "/", http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		fmt.Println(n)
	}

	rows, err := db.Query(`SELECT cName FROM customer`)
	check(err)

	var (
		name  string
		names []string
	)

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		names = append(names, name)
	}
	tpl.ExecuteTemplate(w, "list.html", names)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

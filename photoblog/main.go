package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var tpl *template.Template

var db *sql.DB
var err error

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

}
func main() {
	if db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/users?charset=utf8"); err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/images", images)
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./assets/images"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.html", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		pingDB(w)
		n := req.FormValue("username")
		p := req.FormValue("password")

		// if _, ok := dbUsers[n]; ok {
		// 	http.Error(w, "Username is already taken", http.StatusForbidden)
		// 	return
		// }

		id := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}

		http.SetCookie(w, c)
		bp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Redirect(w, req, "/", http.StatusInternalServerError)
			return
		}

		// try inserting user into database
		// !IMPORTANT
		if _, err := db.Exec(
			"INSERT INTO users (username, password) VALUES (?, ?)",
			n,
			bp,
		); err != nil {
			// if there are any errors
			// convert error into MySQL error
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				if mysqlErr.Number == 1062 {
					// if the error is 1062: duplicate entry
					// show error
					http.Error(w, "Such user exists", http.StatusForbidden)
					return
				}
				// if not this error
				// fatal error
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Fatal(err)
				log.Fatal(mysqlErr)
				return
			}
		}
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	if req.Method == http.MethodPost {
		pingDB(w)

		n := req.FormValue("username")
		p := req.FormValue("password")

		var pass []byte

		// insert password into pass
		// if no user - error
		// if other error - fatal
		// if user exist: if password didn't match - error
		err := db.QueryRow(
			"SELECT password FROM users WHERE username=?",
			n,
		).Scan(&pass)
		switch {
		case err == sql.ErrNoRows:
			http.Error(w, "No Such User", http.StatusForbidden)
			return
		case err != nil:
			log.Fatal(err)
			return
		default:
			if bcrypt.CompareHashAndPassword(pass, []byte(p)) != nil {
				http.Error(w, "Password didn't match", http.StatusForbidden)
				return
			}
		}

		// if !ok {
		// 	http.Error(w, "No such user", http.StatusForbidden)
		// 	return
		// }

		id := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)
		createSession(c.Value, user{n})
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, err := req.Cookie("session")
	if err != nil {
		http.Error(w, "Error with cookies", http.StatusInternalServerError)
		return
	}
	deleteSession(c.Value)

	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	})

	http.Redirect(w, req, "/", http.StatusSeeOther)

}

func images(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		// read file from form
		file, header, err := req.FormFile("uploadfile")
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// convert file into byte slice
		bs, err := ioutil.ReadAll(file)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}

		// create a new file on the server under
		// the path: assets/images
		t := path.Join("assets", "images", time.Now().String()[:25]+header.Filename)
		err = ioutil.WriteFile(t, bs, os.ModePerm)
		if err != nil {
			log.Fatal(err)
			return
		}

	}

	// slice of images to use as a context
	var images []string

	// read all images
	fs, err := ioutil.ReadDir("./assets/images")
	if err != nil {
		fmt.Println(fs)
		log.Fatal(err)
	}
	for _, f := range fs {
		images = append(images, f.Name())
	}

	tpl.ExecuteTemplate(w, "images.html", images)
}

package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var dbSessions = map[string]session{}
var dbUsers = map[string]user{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	fmt.Println(dbSessions)
	fmt.Println(dbUsers)
	tpl.ExecuteTemplate(w, "index.html", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		n := req.FormValue("username")
		p := req.FormValue("password")

		if _, ok := dbUsers[n]; ok {
			http.Error(w, "Username is already taken", http.StatusForbidden)
			return
		}

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
		u := user{n, bp}

		dbUsers[n] = u

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
		n := req.FormValue("username")
		p := req.FormValue("password")

		u, ok := dbUsers[n]
		if !ok {
			http.Error(w, "No such user", http.StatusForbidden)
			return
		}
		if bcrypt.CompareHashAndPassword(u.Password, []byte(p)) != nil {
			http.Error(w, "Password didn't match", http.StatusForbidden)
			return
		}

		id := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{n}
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
		// TODO
		return
	}
	delete(dbSessions, c.Value)

	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	})

	http.Redirect(w, req, "/", http.StatusSeeOther)

}

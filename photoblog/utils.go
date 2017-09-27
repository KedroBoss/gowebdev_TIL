package main

import (
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

// take uuid from cookie
// if the uuid is in sessions return true
func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	u, ok := dbSessions[c.Value]
	if ok {
		if _, err = db.Query(
			"SELECT username FROM users WHERE username=?",
			u.un,
		); err != nil {
			log.Fatal(err)
			return false
		}
	}
	return ok

}

// take cookie
// if doesn't exist - create one and set it
// look for a session with the uuid
// if exists - take the user and return it
func getUser(w http.ResponseWriter, req *http.Request) user {
	c, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
	}
	http.SetCookie(w, c)

	var un string

	s, ok := dbSessions[c.Value]
	if ok {
		db.QueryRow(
			"SELECT username FROM users WHERE username=?",
			s.un,
		).Scan(&un)
	}
	return user{un}
}

func pingDB(w http.ResponseWriter) {
	if err := db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

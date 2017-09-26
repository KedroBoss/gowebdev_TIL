package main

import (
	"github.com/satori/go.uuid"
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
	_, ok = dbUsers[u.un]
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

	var u user

	s, ok := dbSessions[c.Value]
	if ok {
		u = dbUsers[s.un]
	}
	return u
}

package main

import (
	"github.com/satori/go.uuid"
	"net/http"
)

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	u, ok := dbSessions[c.Value]
	_, ok = dbUsers[u.un]
	return ok

}

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

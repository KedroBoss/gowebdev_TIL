package main

import (
	"github.com/go-sql-driver/mysql"
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
	u, ok := getUserFromSession(c.Value)
	if ok {
		if _, err = db.Query(
			"SELECT username FROM users WHERE username=?",
			u.Username,
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

	u, ok := getUserFromSession(c.Value)
	if ok {
		db.QueryRow(
			"SELECT username FROM users WHERE username=?",
			u.Username,
		).Scan(&un)
	}
	return user{un}
}

func getUserFromSession(uuid string) (u user, ok bool) {
	var n string
	if err := db.QueryRow(
		"SELECT user FROM sessions WHERE uuid=?",
		uuid,
	).Scan(&n); err != nil {
		return user{}, false
	}
	return user{n}, true
}

func getSessionFromUser(u user) string {
	var uuid string
	if err := db.QueryRow(
		"SELECT uuid FROM sessions WHERE user=?",
		u.Username,
	).Scan(&uuid); err != nil {
		log.Fatal(err)
	}
	return uuid
}

func createSession(uuid string, u user) {
	if _, err := db.Exec(
		"INSERT INTO sessions (uuid, user) VALUES (?, ?)",
		uuid,
		u.Username,
	); err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				if _, err := db.Exec(
					"UPDATE sessions SET uuid=? WHERE user=?",
					uuid,
					u.Username,
				); err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
		}
	}
}

func deleteSession(uuid string) {
	if _, err := db.Exec(
		"DELETE FROM sessions WHERE uuid=?",
		uuid,
	); err != nil {
		log.Fatal(err)
	}
}

func pingDB(w http.ResponseWriter) {
	if err := db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

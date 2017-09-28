package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"time"
)

const users int = 10

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789_")

var db *sql.DB
var err error

func randString(n int, symbols []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/users?charset=utf8"); err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	for i := 0; i < users; i++ {
		db.Exec(
			"INSERT INTO users (username, password) VALUES (?, ?)",
			randString(10, letterRunes),
			randString(10, numberRunes),
		)
	}
	time.Sleep(time.Second * 20)
	db.Exec("DELETE FROM users.users ORDER BY idusers DESC LIMIT 10;")

}

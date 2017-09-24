package main

import (
	"time"
)

type user struct {
	Username  string
	Password  []byte
	FirstName string
	LastName  string
	Role      string
}

type session struct {
	un           string
	lastActivity time.Time
}

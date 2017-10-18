package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Category struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
}

func main() {
	// Connect to MongoDB server
	mongodDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Timeout:  60 * time.Second,
		Database: "temp",
		Username: "boss",
		Password: "bossthebest",
	}
	// session, err := mgo.Dial("localhost")
	session, err := mgo.DialWithInfo(mongodDialInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Three type of consistency rules:
	// Eventual, Monotonic, Strong
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("temp").C("categories")

	doc := Category{
		bson.NewObjectId(),
		"Category 1",
		"Description for Category 1",
	}
	if err = c.Insert(&doc); err != nil {
		log.Fatal(err)
	}

	var count int
	count, err = c.Count()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v - %v", c, count)
}

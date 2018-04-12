package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// DB database
var DB *mgo.Database

// Profiles collection
var Profiles *mgo.Collection

func init() {
	// Your mongodb connection
	s, err := mgo.Dial("mongodb://localhost/gokitmgo")
	if err != nil {
		panic(err)
	}
	if err = s.Ping(); err != nil {
		panic(err)
	}
	// Your database name
	DB = s.DB("gokitmgo")
	fmt.Println("You connected to your mongo database.")
	// Profiles collection
	Profiles = DB.C("profiles")
}

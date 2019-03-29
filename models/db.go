package models

import (
	"github.com/globalsign/mgo"
	"log"
)

const (
	DbName = "goewanshop"
)

func OpenDB() {
	s, err := mgo.Dial(dbhost)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	globalS = s
}

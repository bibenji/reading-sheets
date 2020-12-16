package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

type poll struct {
	Options []string
}

var db *mgo.Session

func main() {}

func dialdb() error {
	var err error
	log.Println("dialing mongodb: localhost:32773")
	db, err = mgo.Dial("localhot:32773")
	return err
}

func closedb() {
	db.Close()
	log.Println("closed database connection")
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}

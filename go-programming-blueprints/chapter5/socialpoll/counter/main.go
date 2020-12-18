package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bitly/go-nsq"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const updateDuration = 1 * time.Second

var fatalErr error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

func main() {
	// this part of code is on top because all the defer that will follow will be execute before due to the LIFO (last in first out) system of defer
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()

	log.Println("Connecting to database...")
	db, err := mgo.Dial("localhost:32773")
	if err != nil {
		fatal(err)
		return
	}
	defer func() {
		log.Println("Closing database connection...")
		db.Close()
	}()
	pollData := db.DB("ballots").C("polls") // reference of the ballots.polls data collection

	// map and mutex = common combination in go, for multiple go routines access the map not at same time
	var counts map[string]int
	var countsLock sync.Mutex

	log.Println("Connecting to nsq...")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())
	if err != nil {
		fatal(err)
		return
	}

	q.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLock.Lock() // so others has to wait here before it unlocks in order to use it themselve
		defer countsLock.Unlock()
		if counts == nil {
			counts = make(map[string]int)
		}
		vote := string(m.Body)
		counts[vote]++
		return nil
	}))

	if err := q.ConnectToNSQD("localhost:32770"); err != nil {
		fatal(err)
		return
	}

	// connect to nsq service
	if err := q.ConnectToNSQLookupd("localhost:32769"); err != nil {
		fatal(err)
		return
	}

	ticker := time.NewTicker(updateDuration)
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		select {
		case <-ticker.C: // give us a channel on which current time is sent at the specified interval
			doCount(&countsLock, &counts, pollData)
		case <-termChan: // we go there many time until q effectively stopped
			ticker.Stop()
			q.Stop()
		case <-q.StopChan: // and when q effectively stopped we're done :-)
			// finished
			return
		}
	}

}

func doCount(countsLock *sync.Mutex, counts *map[string]int, pollData *mgo.Collection) {
	countsLock.Lock()
	defer countsLock.Unlock()

	if len(*counts) == 0 {
		log.Println("No new votes, skipping database update")
		return
	}

	log.Println("Updating database...")
	log.Println(*counts)

	ok := true
	for option, count := range *counts {
		// bson = Binary JSON
		// { "options": { "$in": ["banane"] } }
		// => select polls where "banane" is one of the items in the options array
		sel := bson.M{"options": bson.M{"$in": []string{option}}}
		// { "$inc": { "results.banane": 3 } }
		// => update operation
		up := bson.M{"$inc": bson.M{"results." + option: count}}

		if _, err := pollData.UpdateAll(sel, up); err != nil {
			log.Println("failed to update:", err)
			ok = false
		}
	}

	if ok {
		log.Println("Finished updating database...")
		*counts = nil // reset counts
	}
}

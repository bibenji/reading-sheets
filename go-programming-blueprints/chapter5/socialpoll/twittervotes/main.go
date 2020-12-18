package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bitly/go-nsq"

	"gopkg.in/mgo.v2"
)

type poll struct {
	Options []string
}

var db *mgo.Session

func main() {
	if err := dialdb(); err != nil {
		log.Fatalln("failed to dial MongoDB:", err)
	}
	defer closedb()

	var stoplock sync.Mutex // protects stop to access it from many go routines
	stop := false
	stopChan := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan // block until we get a signal on that channel
		stoplock.Lock()
		stop = true
		stoplock.Unlock()
		log.Println("Stopping...")
		stopChan <- struct{}{} // send signal through stopChan for twitterStreaming
		closeConn()
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM) // go send the signal down signalChan when someone try to kill with SIGINT or SIGTERM

	votes := make(chan string)                                // votes = channel of strings
	publisherStoppedChan := publishVotes(votes)               // stop signal channel
	twitterStoppedChan := startTwitterStream(stopChan, votes) // idem, stop signal channel as return
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			closeConn() // close the connnection to restart it and stream twitter api
			stoplock.Lock()
			if stop {
				stoplock.Unlock()
				return // we exit the go routine
			}
			stoplock.Unlock()
		}
	}()
	<-twitterStoppedChan   // script is block waiting for a signal
	close(votes)           // when there is a signal we close the votes channel
	<-publisherStoppedChan // we wait for it before exit
}

func dialdb() error {
	var err error
	log.Println("dialing mongodb: localhost:32773")
	db, err = mgo.Dial("localhost:32773")
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

func publishVotes(votes <-chan string) <-chan struct{} {
	stopchan := make(chan struct{}, 1)
	pub, _ := nsq.NewProducer("localhost:32770", nsq.NewConfig())
	go func() {
		for vote := range votes {
			pub.Publish("votes", []byte(vote))
		}
		log.Println("Publisher: Stopping")
		pub.Stop()
		log.Println("Published: Stopped")
		stopchan <- struct{}{}
	}()
	return stopchan
}

package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type tweet struct {
	Text string
}

var conn net.Conn
var reader io.ReadCloser
var (
	buildHTTPClientOnce sync.Once
	httpClient          *http.Client
)

func dial(netw, addr string) (net.Conn, error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}

	netc, err := net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil {
		return nil, err
	}

	conn = netc
	return netc, nil
}

// not used ???
func closeConn() {
	if conn != nil {
		conn.Close()
	}
	if reader != nil {
		reader.Close()
	}
}

func makeRequest(req *http.Request, params url.Values) (*http.Response, error) {
	buildHTTPClientOnce.Do(func() {
		httpClient = &http.Client{
			Transport: &http.Transport{
				Dial: dial,
			},
		}
	})

	formEnc := params.Encode()
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Content-Length", strconv.Itoa(len(formEnc)))
	// req.Header.Set("Authorization", authClient.AuthorizationHeader(creds, "POST", req.URL, params))

	return httpClient.Do(req)
}

func readFromTwitter(votes chan<- string) {
	// we read options from db
	options, err := loadOptions()
	if err != nil {
		log.Println("failed to load options:", err)
		return
	}

	// we prepare query params
	u, err := url.Parse("http://localhost:8090")
	if err != nil {
		log.Println("creating request failed:", err)
		return
	}
	query := make(url.Values)
	// query.Set("track", strings.Join(options, ","))

	// we make the query
	req, err := http.NewRequest("GET", u.String(), strings.NewReader(query.Encode()))
	if err != nil {
		log.Println("Creating filter request failed:", err)
		return
	}
	resp, err := makeRequest(req, query)
	if err != nil {
		log.Println("making request failed:", err)
		return
	}

	// get a reader and a decoder
	reader := resp.Body
	decoder := json.NewDecoder(reader)

	// infinite loop where we read tweets
	for {
		var t tweet
		if err := decoder.Decode(&t); err != nil {
			break
		}
		for _, option := range options {
			if strings.Contains(
				strings.ToLower(t.Text),
				strings.ToLower(option),
			) {
				log.Println("vote:", option)
				votes <- option
			}
		}
	}
}

func startTwitterStream(stopchan <-chan struct{}, votes chan<- string) <-chan struct{} {
	stoppedchan := make(chan struct{}, 1)
	go func() {
		defer func() {
			stoppedchan <- struct{}{}
		}()

		for {
			select {
			case <-stopchan:
				log.Println("stopping Twitter...")
				return
			default:
				log.Println("Querying Twitter...")
				readFromTwitter(votes)
				log.Println(" (waiting)")
				time.Sleep(10 * time.Second) // wait before reconnecting
			}
		}
	}()
	return stoppedchan
}

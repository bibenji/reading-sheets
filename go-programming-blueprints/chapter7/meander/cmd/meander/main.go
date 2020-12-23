package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"../../../meander"
)

func main() {
	apiKey, present := os.LookupEnv("GOOGLE_PLACES_API_KEY")
	if !present {
		log.Fatalln("GOOGLE_PLACES_API_KEY must be set as env variable")
	}
	if apiKey == "" {
		log.Fatalln("GOOGLE_PLACES_API_KEY has been wrongly setted")
	}
	meander.APIKey = apiKey

	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8040", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

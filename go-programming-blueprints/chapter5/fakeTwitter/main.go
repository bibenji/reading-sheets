package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type randomJSONHandler struct {
	content string
}

// Tweet one fake tweet
type Tweet struct {
	Text string `json:"text"`
}

var options = []string{"pomme", "poire", "fraise", "framboise", "banane", "kiwi"}

func randOption() string {
	return options[rand.Intn(len(options))]
}

func (h *randomJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var tweets = []Tweet{
		Tweet{
			Text: randOption(),
		},
		Tweet{
			Text: randOption(),
		},
		Tweet{
			Text: options[0],
		},
	}

	JSONContent, err := json.Marshal(tweets)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(JSONContent))
}

func main() {
	http.Handle("/stream", &randomJSONHandler{})

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

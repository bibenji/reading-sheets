package main

import (
	"flag"
	"log"
	"net/http"

	"../trace"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	r := newRoom()

	// set tracer
	// r.tracer = trace.New(os.Stdout)
	// silent tracer
	r.tracer = trace.Off()

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("/assets/"))))

	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

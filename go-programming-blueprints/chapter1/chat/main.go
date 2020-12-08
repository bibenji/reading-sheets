package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"../trace"
)

// var avatars Avatar = UserFileSystemAvatar

var avatars Avatar = TryAvatars{
	UserFileSystemAvatar,
	UserAuthAvatar,
	UserGravatar}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	providerIndex := newProviderIndex()

	// r := newRoom(UserAuthAvatar)
	// r := newRoom(UserGravatar)
	// r := newRoom(UserFileSystemAvatar)
	r := newRoom()

	// set tracer
	r.tracer = trace.New(os.Stdout)
	// silent tracer
	// r.tracer = trace.Off()

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))

	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars"))))

	http.Handle("/login", &templateHandler{
		filename: "login.html",
		data:     map[string]interface{}{"providerIndex": providerIndex}})

	http.HandleFunc("/auth/", loginHandler)

	http.Handle("/user", MustAuth(&templateHandler{filename: "user.html"}))

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))

	http.Handle("/room", MustAuth(r))

	http.Handle("/upload", MustAuth(&templateHandler{filename: "upload.html"}))

	http.Handle("/uploader", &uploaderHandler{})

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

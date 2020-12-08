package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"../trace"
)

type room struct {
	// channel that holds incoming messages
	// forward chan []byte
	forward chan *message

	// channel for clients wishing to join the room
	join chan *client

	// channel for clients wishing to leave the room
	leave chan *client

	// holds all current clients in this room
	clients map[*client]bool

	// tracer will receive trace information of activity in the room
	tracer trace.Tracer

	// data to pass to the handler
	data map[string]interface{}
}

func newRoom() *room {
	return &room{
		// forward: make(chan []byte),
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		// tracer: trace.Off(),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("New client joined: ", client.UserData["FirstName"])
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left: ", client.UserData["FirstName"])
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", msg.Message)
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- sent to client")
			}
		}
	}
}

const socketBufferSize = 1024
const messageBufferSize = 256

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.tracer.Trace("room.go: ServeHTTP")

	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		// send:   make(chan []byte, messageBufferSize),
		send:   make(chan *message, messageBufferSize),
		room:   r,
		tracer: r.tracer,
	}

	client.UserData = r.data

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}

func (r *room) SetData(data map[string]interface{}) {
	r.tracer.Trace("room.go: SetData", data)
	r.data = data
}

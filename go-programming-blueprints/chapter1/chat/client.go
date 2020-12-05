package main

import (
	"log"
	"time"

	"../trace"
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn

	// send     chan []byte
	send chan *message

	room *room

	// UserData holds informations about the user
	UserData map[string]interface{}

	// tracer will receive trace information of client's activity
	tracer trace.Tracer
}

func (c *client) read() {
	c.tracer.Trace("client.go: read")
	defer c.socket.Close()
	for {
		// _, msg, err := c.socket.ReadMessage()
		var msg *message
		if err := c.socket.ReadJSON(&msg); err != nil {
			c.tracer.Trace("client.go: read inside for err: ", err)
			return
		}

		msg.When = time.Now()
		msg.Name = c.UserData["FirstName"].(string)
		// if avatarURL, ok := c.UserData["AvatarURL"]; ok {
		// 	msg.AvatarURL = avatarURL.(string)
		// }
		msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)

		log.Println(c.UserData)

		c.room.forward <- msg
	}
}

func (c *client) write() {
	c.tracer.Trace("client.go: write")
	defer c.socket.Close()
	for msg := range c.send {
		// err := c.socket.WriteMessage(websocket.TextMessage, msg)
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}

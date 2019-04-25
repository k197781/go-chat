package main

import (
	"github.com/gorilla/websocket"
)

// client is struct of a chatting user
type client struct {
	socket *websocket.Conn
	// recieved message is added to send field and wait for being sent to clients.
	send chan []byte
	// chatroom that user is showing is added to room field.
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

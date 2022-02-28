package app

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

var clients []*Client

type Client struct {
	id       string
	username string
	ip       string
	ws       *websocket.Conn
}

func (c *Client) Listen() {

	buffer := make([]byte, 1024)

	for {

		n, err := c.ws.Read(buffer[0:])

		if err != nil {
			CloseConnection(c)

			exitMessage := &Message{
				SenderID: "System",
				Username: "System",
				Text:     fmt.Sprintf("%s has left the chat", c.username),
			}

			exitMessage.Broadcast()
			break
		} else {
			// Broadcast the message to all clients
			HandleInputMessage(c, buffer[0:n])
		}
	}

}

func HandleNewConnection(c *websocket.Conn) {
	log.Println("New connection")

	client := &Client{
		id:       uuid.New().String(),
		username: "",
		ip:       c.Request().RemoteAddr,
		ws:       c,
	}

	clients = append(clients, client)

	client.Listen()
}

func CloseConnection(c *Client) {
	log.Printf("Closing connection for %s, %s", c.username, c.ip)
	c.ws.Close()

	index := -1
	for i, client := range clients {
		if client == c {
			index = i
			break
		}
	}

	if index >= 0 {
		clients = append(clients[:index], clients[index+1:]...)
	}
}

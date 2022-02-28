package app

import (
	"encoding/json"
)

var messages []*Message

type Message struct {
	SenderID string `json:"senderId"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

func (m *Message) Post() {
	m.Broadcast()
	messages = append(messages, m)
}

func (m *Message) Broadcast() {
	for _, client := range clients {
		m.BroadcastTo(client)
	}
}

func (m *Message) BroadcastTo(to *Client) {
	byteMessage, _ := json.Marshal(m)
	to.ws.Write(byteMessage)
}

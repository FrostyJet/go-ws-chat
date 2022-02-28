package app

import (
	"encoding/json"
	"fmt"
)

func HandleInputMessage(from *Client, data []byte) {
	var input map[string]string
	json.Unmarshal(data, &input)

	switch input["action"] {
	case "post_message":
		message := &Message{
			SenderID: from.id,
			Username: from.username,
			Text:     input["text"],
		}

		message.Post()
	case "initial_connection":
		from.username = input["username"]
		if from.username == "" {
			from.username = "Guest"
		}

		response := &Message{
			SenderID: "System",
			Username: "System",
			Text:     fmt.Sprintf("YOURID:%s", from.id),
		}

		response.BroadcastTo(from)

		// Send all previous messages to the client
		history, _ := json.Marshal(&messages)
		from.ws.Write(history)

		message := &Message{
			SenderID: "System",
			Username: "System",
			Text:     fmt.Sprintf("%s has joined the chat", from.username),
		}

		message.Broadcast()
	}
}

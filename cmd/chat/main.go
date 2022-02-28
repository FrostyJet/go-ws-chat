package main

import (
	"log"
	"net/http"

	"github.com/frostyjest/go-ws-chat/internal/app"
	"golang.org/x/net/websocket"
)

func init() {
	http.Handle("/ws", websocket.Handler(app.HandleNewConnection))
}

func main() {
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

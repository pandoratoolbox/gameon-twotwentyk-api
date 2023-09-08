package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WSConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		// Read message from client
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			// Handle error or client disconnect
			return
		}

		// Handle the message or send notifications to other clients
		// You can broadcast this message to all connected clients here
		_ = messageType
		_ = p
	}
}

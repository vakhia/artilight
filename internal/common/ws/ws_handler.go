package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Clients   = make(map[*websocket.Conn]bool)
	Broadcast = make(chan Message)
)

type Message struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func HandleWebSocket(c *gin.Context) {
	ws, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer ws.Close()

	Clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			delete(Clients, ws)
			break
		}
		Broadcast <- msg
	}
}

func BroadcastToClients() {
	for {
		msg := <-Broadcast
		for client := range Clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Printf("WebSocket write error: %v", err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}

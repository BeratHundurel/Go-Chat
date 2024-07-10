package hub

import (
	"encoding/json"
	"go-chat/app/services"
	"go-chat/app/types"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var chatHub = NewHub()

func ReturnHub() *Hub {
	return chatHub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				client.Conn.Close()
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				client.Conn.WriteMessage(websocket.TextMessage, message)
			}
		}
	}
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	client := &Client{Conn: conn}
	chatHub.Register <- client

	defer func() {
		chatHub.Unregister <- client
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var initMsg types.Message
		if err := json.Unmarshal(msg, &initMsg); err != nil {
			log.Println("Failed to marshal message:", err)
			return
		}

		services.SaveMessage(initMsg)

		chatHub.Broadcast <- msg
	}
}

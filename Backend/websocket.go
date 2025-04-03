package backend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Lista de clientes conectados y Canal de mensajes
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error al conectar WebSocket:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true
	for {
		var msg string
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}

}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)

			}
		}
	}
}

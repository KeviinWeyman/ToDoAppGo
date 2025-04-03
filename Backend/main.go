package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handleConnections) // Ruta WebSocket
	go handleMessages()                       // Goroutine para manejar mensajes en segundo plano

	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error al iniciar servidor:", err)
	}
}

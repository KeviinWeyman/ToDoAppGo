package backend

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/w", handleConnections) // Ruta para WebSocket
	go handleMessages()

	fmt.Println("Servidor escuchando en LocalHost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)

	}
}

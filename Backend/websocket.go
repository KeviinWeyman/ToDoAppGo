package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Task struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
	Action    string `json:"action"`
}

var tasks = make(map[string]Task)
var tasksMutex sync.Mutex

var taskChannel = make(chan Task)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error al conectar WebSocket:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true
	fmt.Println("Nuevo cliente conectado:", ws.RemoteAddr())

	sendCurrentTasks(ws)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error leyendo mensaje de cliente:", err)
			delete(clients, ws)
			break
		}

		var task Task
		if err := json.Unmarshal(msg, &task); err != nil {
			fmt.Println("Error al deserializar tarea:", err)
			continue
		}

		taskChannel <- task
	}
}

func sendCurrentTasks(ws *websocket.Conn) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	for _, task := range tasks {
		taskJSON, _ := json.Marshal(task)
		ws.WriteMessage(websocket.TextMessage, taskJSON)
	}
}

func handleMessages() {
	for {
		task := <-taskChannel
		switch task.Action {
		case "add":
			addTask(task)
		case "edit":
			editTask(task)
		case "delete":
			deleteTask(task.ID)
		}
	}
}

func addTask(task Task) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	tasks[task.ID] = task
	broadcastTaskUpdate(task)
}

func editTask(task Task) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	if _, exists := tasks[task.ID]; exists {
		tasks[task.ID] = task
		broadcastTaskUpdate(task)
	}
}

func deleteTask(taskID string) {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	if _, exists := tasks[taskID]; exists {
		delete(tasks, taskID)
		broadcastTaskDelete(taskID)
	}
}

func broadcastTaskUpdate(task Task) {
	for client := range clients {
		taskJSON, _ := json.Marshal(task)
		err := client.WriteMessage(websocket.TextMessage, taskJSON)
		if err != nil {
			log.Println("Error enviando tarea:", err)
		}
	}
}

func broadcastTaskDelete(taskID string) {
	deleteMessage := map[string]string{
		"id":     taskID,
		"action": "delete",
	}

	deleteJSON, _ := json.Marshal(deleteMessage)

	for client := range clients {
		client.WriteMessage(websocket.TextMessage, deleteJSON)
	}
}

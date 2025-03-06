package internal

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Разрешить все соединения
}

var clients = make(map[*websocket.Conn]string) // Клиенты и их ID
var mutex sync.Mutex                           // Защита карты клиентов

type MessageChat struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ошибка при обновлении до WebSocket:", err)
		return
	}
	defer ws.Close()

	userID := r.URL.Query().Get("id") // Получаем ID пользователя
	if userID == "" {
		userID = "пользователь_неизвестный"
	} else if userID == "admin" {
		userID = "администратор"
	} else {
		userID = "пользователь_" + userID
	}

	mutex.Lock()
	clients[ws] = userID
	mutex.Unlock()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			break
		}

		message := MessageChat{Sender: userID, Message: string(msg)}
		broadcastMessage(message)
	}
}

func broadcastMessage(msg MessageChat) {
	messageJSON, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, messageJSON); err != nil {
			fmt.Println("Ошибка отправки:", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func main() {
	http.HandleFunc("/ws", HandleConnections)
	fmt.Println("WebSocket сервер запущен на :10003")
	err := http.ListenAndServe(":10003", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

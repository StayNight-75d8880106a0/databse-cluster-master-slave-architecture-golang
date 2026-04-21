package helper

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WSManager struct {
	Clients map[*websocket.Conn]bool
	Mu      sync.Mutex
}

var Manager = &WSManager{
	Clients: make(map[*websocket.Conn]bool),
}

func (m *WSManager) Broadcast(message interface{}) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	for client := range m.Clients {
		err := client.WriteJSON(message)
		if err != nil {
			client.Close()
			delete(m.Clients, client)
		}
	}
}

func (m *WSManager) SendRaw(conn *websocket.Conn, message interface{}) error {

	m.Mu.Lock()

	defer m.Mu.Unlock()

	err := conn.WriteJSON(message)

	if err != nil {
		return err
	}

	return nil
}

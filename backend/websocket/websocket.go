package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"product_manage/model"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Cấu hình CORS nếu cần thiết
	},
}

type WebSocketManager struct {
	clients   map[*websocket.Conn]bool
	mu        sync.Mutex       // Đảm bảo an toàn trong môi trường đồng thời
	broadcast chan interface{} // Unified broadcast channel
}

// NewWebSocketManager khởi tạo WebSocket manager
func NewWebSocketManager() *WebSocketManager {
	manager := &WebSocketManager{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan interface{}),
	}
	go manager.handleBroadcast()
	return manager
}

// HandleConnections quản lý kết nối WebSocket
func (manager *WebSocketManager) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	manager.AddClient(conn)
	log.Println("New WebSocket client connected")

	defer func() {
		manager.RemoveClient(conn)
		conn.Close()
		log.Println("WebSocket client disconnected")
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		manager.broadcast <- msg // Gửi tin nhắn vào channel phát
	}
}

// AddClient thêm client vào manager
func (manager *WebSocketManager) AddClient(client *websocket.Conn) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	manager.clients[client] = true
}

// RemoveClient xóa client khỏi manager
func (manager *WebSocketManager) RemoveClient(client *websocket.Conn) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	delete(manager.clients, client)
}

// Broadcast gửi tin nhắn tới tất cả client
func (manager *WebSocketManager) Broadcast(message interface{}) {
	select {
	case manager.broadcast <- message:
		// Gửi thành công vào channel
	case <-time.After(1 * time.Second):
		log.Println("Broadcast timeout: message could not be sent")
	}
}

func (manager *WebSocketManager) sendMessage(client *websocket.Conn, message interface{}) error {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	messageData, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error encoding message data: %v", err)
		return err
	}
	if err := client.WriteMessage(websocket.TextMessage, messageData); err != nil {
		log.Printf("Error sending message to client: %v", err)
		client.Close()
		manager.RemoveClient(client)
		return err
	}
	return nil
}

// handleBroadcast lắng nghe channel phát và gửi tới từng client
func (manager *WebSocketManager) handleBroadcast() {
	for data := range manager.broadcast {
		manager.mu.Lock()
		for client := range manager.clients {
			go manager.sendMessage(client, data)
		}
		manager.mu.Unlock()
	}
}

// Gửi thông báo khi có sự thay đổi sản phẩm
func (manager *WebSocketManager) NotifyProductChange(productID string, action string) {
	messageData := map[string]interface{}{
		"action":    action,
		"productID": productID,
	}
	// messageData, err := json.Marshal(message)
	// if err != nil {
	// 	log.Printf("Error encoding message data: %v", err)
	// 	return
	// }
	manager.Broadcast(messageData)
	log.Printf("Product change message: %s\n", messageData)
}

// Gửi sản phẩm khi có sự thay đổi sản phẩm
func (manager *WebSocketManager) SendProductChange(data interface{}, action string) {
	var message map[string]interface{}
	log.Printf("Received data type: %T\n", data)
	switch v := data.(type) {
	case *model.Product:
		message = map[string]interface{}{
			"action":  action,
			"product": v,
		}
	case string:
		message = map[string]interface{}{
			"action":    action,
			"productID": v,
		}
	default:
		log.Println("Invalid data type")
		return
	}
	log.Printf("Sent product change message: %s\n", message)
	manager.Broadcast(message)
}

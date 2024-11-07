package controller

import (
	"net/http"
	"product_manage/websocket"
)

type WebSocketController struct {
	manager *websocket.WebSocketManager
}

func NewWebSocketController(manager *websocket.WebSocketManager) *WebSocketController {
	return &WebSocketController{manager: manager}
}

func (c *WebSocketController) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	c.manager.HandleConnections(w, r)
}

func (c *WebSocketController) NotifyProductChange(productID string, action string) {
	c.manager.NotifyProductChange(productID, action)
}
func (c *WebSocketController) SendProductChange(data interface{}, action string) {
	c.manager.SendProductChange(data, action)
}

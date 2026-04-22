package ws_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/ai/ai_controller"
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"fmt"

	"github.com/gin-gonic/gin"
)

type WebSocket_Controller struct {
	AI_Controller *ai_controller.AI_Controller
}

func NewWebSocketControllerRegistry(ai_controller *ai_controller.AI_Controller) *WebSocket_Controller {
	return &WebSocket_Controller{
		AI_Controller: ai_controller,
	}
}

// @Summary Koneksi WebSocket AI Chat
// @Description Membuka koneksi WebSocket untuk chat AI. Endpoint ini menghubungkan client ke server WebSocket, menyimpan koneksi, dan meneruskan pesan ke AI Controller untuk diproses secara real-time. Cocok untuk fitur chat interaktif berbasis AI.
// @Tags WebSocket
// @Produce json
// @Success 101 {string} string "Switching Protocols: WebSocket handshake berhasil, koneksi terbuka"
// @Failure 500 {object} map[string]interface{} "Gagal membuka koneksi WebSocket atau error internal"
// @Router /api/ws/connect [get]
func (c *WebSocket_Controller) Connect(ctx *gin.Context) {

	conn, err := helper.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		fmt.Println("An error has occurred with the WebSocket!!")
		return
	}

	helper.Manager.Mu.Lock()
	helper.Manager.Clients[conn] = true
	helper.Manager.Mu.Unlock()

	defer func() {
		helper.Manager.Mu.Lock()
		delete(helper.Manager.Clients, conn)
		helper.Manager.Mu.Unlock()
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		go c.AI_Controller.HandleAiChat(conn, string(msg))
	}

}

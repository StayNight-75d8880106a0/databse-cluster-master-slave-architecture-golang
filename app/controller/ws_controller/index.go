package ws_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"fmt"

	"github.com/gin-gonic/gin"
)

type WebSocket_Controller struct{}

func NewWebSocketControllerRegistry() *WebSocket_Controller {
	return &WebSocket_Controller{}
}

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
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

}

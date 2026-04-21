package ai_controller

import (
	"context"
	"databse-cluster-master-slave-architecture-golang/app/ai/chains"
	"databse-cluster-master-slave-architecture-golang/app/helper"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type AI_Controller struct {
	Chain *chains.AI_Chains
}

func NewAiControllerRegistry(chain *chains.AI_Chains) *AI_Controller {
	return &AI_Controller{
		Chain: chain,
	}
}

func (c *AI_Controller) HandleAiChat(conn *websocket.Conn, userMSG string) {
	ctx := context.Background()

	_, err := c.Chain.GenerateResponse(ctx, userMSG, func(ctx context.Context, chunk []byte) error {
		helper.Manager.SendRaw(conn, gin.H{
			"event": "AI_CHAT_STREAM",
			"data":  string(chunk),
		})
		return nil
	})

	if err != nil {
		helper.Manager.SendRaw(conn, gin.H{"event": "AI_CHAT_ERROR", "data": err.Error()})
		return
	}

	helper.Manager.SendRaw(conn, gin.H{"event": "AI_CHAT_DONE", "data": ""})

}

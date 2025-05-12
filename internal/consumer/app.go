package consumer

import "github.com/gin-gonic/gin"

type Consumer struct {
	Engine *gin.Engine
}

func New() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Run() {

}

func (c *Consumer) Stop() {

}

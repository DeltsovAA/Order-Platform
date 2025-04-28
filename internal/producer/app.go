package producer

import (
	"github.com/gin-gonic/gin"
)

type Producer struct {
	Engine *gin.Engine
}

func New() *Producer {
	return &Producer{
		Engine: gin.Default(),
	}
}

func (p *Producer) Run() error {
	if err := p.Engine.Run(":8080"); err != nil {
		return err
	}
	return nil
}

func (p *Producer) Stop() {

}

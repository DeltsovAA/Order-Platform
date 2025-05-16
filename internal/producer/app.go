package producer

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func New() *App {
	return &App{
		router: NewRouter(),
	}
}

func (p *App) Run() error {
	if err := p.router.Run(":8080"); err != nil {
		return err
	}
	return nil
}

func (p *App) Stop() {

}

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lapeko/udemy__github-actions-the-complete-guide/internal/api/router"
	"os"
)

type Api interface {
	Init()
	Start() error
}

type apiInstance struct {
	port   string
	engine *gin.Engine
}

func New() Api {
	return &apiInstance{}
}

func (a *apiInstance) Init() {
	a.port = os.Getenv("PORT")
	if a.port == "" {
		a.port = "8000"
	}
	a.engine = gin.Default()
	router.ApplyRoutes(a.engine)
}

func (a *apiInstance) Start() error {
	return a.engine.Run(fmt.Sprintf(":%s", a.port))
}

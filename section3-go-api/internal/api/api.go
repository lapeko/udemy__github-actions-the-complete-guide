package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Api interface {
	Init()
	Start()
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
}

func (a *apiInstance) Start() {
	a.engine.Run(fmt.Sprintf(":%s", a.port))
}

package router

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	SetupHealth(r.Group("/healthz"))
	SetupUsers(r.Group("/users"))
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lapeko/udemy__github-actions-the-complete-guide/internal/api/controller"
)

func SetupHealth(r *gin.RouterGroup) {
	r.GET("/", controller.GetHealthController)
}
